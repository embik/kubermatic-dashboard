import {ChangeDetectionStrategy, Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {FormBuilder, FormGroup} from '@angular/forms';
import {MatStepper} from '@angular/material/stepper';
import {Router} from '@angular/router';
import {forkJoin, of, Subject} from 'rxjs';
import {switchMap, takeUntil, tap} from 'rxjs/operators';

import {ClusterService, DatacenterService, NotificationService, ProjectService} from '../core/services';
import {GoogleAnalyticsService} from '../google-analytics.service';
import {NodeDataService} from '../node-data-new/service/service';
import {ClusterEntity} from '../shared/entity/ClusterEntity';
import {DataCenterEntity} from '../shared/entity/DatacenterEntity';
import {ProjectEntity} from '../shared/entity/ProjectEntity';
import {SSHKeyEntity} from '../shared/entity/SSHKeyEntity';
import {CreateClusterModel} from '../shared/model/CreateClusterModel';
import {NodeData} from '../shared/model/NodeSpecChange';

import {StepRegistry, steps, WizardStep} from './config';
import {ClusterService as ClusterModelService} from './service/cluster';
import {WizardService} from './service/wizard';

@Component({
  selector: 'kubermatic-wizard',
  templateUrl: './template.html',
  styleUrls: ['./style.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class WizardComponent implements OnInit, OnDestroy {
  form: FormGroup;
  project = {} as ProjectEntity;
  creating = false;
  readonly stepRegistry = StepRegistry;

  @ViewChild('stepper', {static: true}) private readonly _stepper: MatStepper;

  private _unsubscribe: Subject<void> = new Subject<void>();

  constructor(
      private readonly _formBuilder: FormBuilder, private readonly _projectService: ProjectService,
      private readonly _wizard: WizardService, private readonly _notificationService: NotificationService,
      private readonly _clusterModelService: ClusterModelService, private readonly _clusterService: ClusterService,
      private readonly _nodeDataService: NodeDataService, private readonly _router: Router,
      private readonly _datacenterService: DatacenterService,
      private readonly _googleAnalyticsService: GoogleAnalyticsService) {}

  get steps(): WizardStep[] {
    return this._wizard.steps.filter(step => step.enabled);
  }

  get active(): WizardStep {
    return this.steps[this._stepper.selectedIndex];
  }

  get first(): boolean {
    return this._stepper.selectedIndex === 0;
  }

  get last(): boolean {
    return this._stepper.selectedIndex === (this.steps.length - 1);
  }

  get invalid(): boolean {
    return this.form.get(this.active.name).invalid;
  }

  ngOnInit(): void {
    // Init steps for wizard
    this._wizard.steps = steps;
    this._wizard.stepper = this._stepper;

    this._initForm(this.steps);
    this._wizard.stepsChanges.pipe(takeUntil(this._unsubscribe)).subscribe(_ => this._initForm(this.steps));

    this._projectService.selectedProject.pipe(takeUntil(this._unsubscribe))
        .subscribe(project => this.project = project);
  }

  showNext(step: WizardStep): boolean {
    const restrictedList = [StepRegistry.Datacenter, StepRegistry.Provider, StepRegistry.Summary];
    return !restrictedList.includes(step.name as StepRegistry);
  }

  ngOnDestroy(): void {
    this._unsubscribe.next();
    this._unsubscribe.complete();
    this._wizard.reset();
  }

  create(): void {
    this.creating = true;
    let createdCluster: ClusterEntity;
    let datacenter: DataCenterEntity;
    const createCluster =
        this._getCreateClusterModel(this._clusterModelService.cluster, this._nodeDataService.nodeData);

    this._datacenterService.getDataCenter(this._wizard.datacenter)
        .pipe(tap(dc => datacenter = dc))
        .pipe(switchMap(_ => this._clusterService.create(this.project.id, datacenter.spec.seed, createCluster)))
        .pipe(tap(cluster => {
          this._notificationService.success(`Cluster ${createCluster.cluster.name} successfully created`);
          this._googleAnalyticsService.emitEvent('clusterCreation', 'clusterCreated');
          createdCluster = cluster;
        }))
        .pipe(switchMap(_ => this._clusterService.cluster(this.project.id, createdCluster.id, datacenter.spec.seed)))
        .pipe(switchMap(() => {
          this.creating = false;

          if (this._clusterModelService.sshKeys.length > 0) {
            return forkJoin(this._clusterModelService.sshKeys.map(
                key => this._clusterService.createSSHKey(
                    this.project.id, createdCluster.id, datacenter.spec.seed, key.id)));
          }

          return of([]);
        }))
        .pipe(takeUntil(this._unsubscribe))
        .subscribe(
            (keys: SSHKeyEntity[]) => {
              this._router.navigate(
                  [`/projects/${this.project.id}/dc/${datacenter.spec.seed}/clusters/${createdCluster.id}`]);
              keys.forEach(
                  key => this._notificationService.success(
                      `SSH key ${key.name} was added successfully to cluster ${createCluster.cluster.name}`));
            },
            () => {
              this._notificationService.error(`Could not create cluster ${createCluster.cluster.name}`);
              this._googleAnalyticsService.emitEvent('clusterCreation', 'clusterCreationFailed');
              this.creating = false;
            });
  }

  private _getCreateClusterModel(cluster: ClusterEntity, nodeData: NodeData): CreateClusterModel {
    const keyNames: string[] = [];
    for (const key of this._clusterModelService.sshKeys) {
      keyNames.push(key.name);
    }

    return {
      cluster: {
        name: cluster.name,
        labels: cluster.labels,
        spec: cluster.spec,
        type: cluster.type,
        sshKeys: keyNames,
        credential: cluster.credential,
      },
      nodeDeployment: {
        name: nodeData.name,
        spec: {
          template: nodeData.spec,
          replicas: nodeData.count,
          dynamicConfig: nodeData.dynamicConfig,
        },
      }
    };
  }

  private _initForm(steps: WizardStep[]): void {
    const controls = {};
    steps.forEach(step => controls[step.name] = this._formBuilder.control(''));
    this.form = this._formBuilder.group(controls);
  }
}
