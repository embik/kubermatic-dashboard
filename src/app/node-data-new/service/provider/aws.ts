import {Observable} from 'rxjs';
import {filter, switchMap} from 'rxjs/operators';
import {DatacenterService, PresetsService} from '../../../core/services';
import {AWSSize, AWSSubnet} from '../../../shared/entity/provider/aws/AWS';
import {NodeProvider} from '../../../shared/model/NodeProviderConstants';
import {ClusterService} from '../../../wizard-new/service/cluster';
import {NodeDataMode} from '../../config';
import {NodeDataService} from '../service';

export class NodeDataAWSProvider {
  constructor(
      private readonly _nodeDataService: NodeDataService, private readonly _clusterService: ClusterService,
      private readonly _presetService: PresetsService, private readonly _datacenterService: DatacenterService) {}

  set tags(tags: object) {
    delete this._nodeDataService.nodeData.spec.cloud.aws.tags;
    this._nodeDataService.nodeData.spec.cloud.aws.tags = tags;
  }

  flavors(): Observable<AWSSize[]> {
    // TODO: support dialog mode
    switch (this._nodeDataService.mode) {
      case NodeDataMode.Wizard:
        return this._clusterService.datacenterChanges.pipe(switchMap(dc => this._datacenterService.getDataCenter(dc)))
            .pipe(switchMap(dc => this._presetService.provider(NodeProvider.AWS).region(dc.spec.aws.region).flavors()));
    }
  }

  subnets(): Observable<AWSSubnet[]> {
    // TODO: support dialog mode
    switch (this._nodeDataService.mode) {
      case NodeDataMode.Wizard:
        return this._clusterService.clusterChanges.pipe(filter(_ => this._clusterService.provider === NodeProvider.AWS))
            .pipe(switchMap(
                cluster => this._presetService.provider(NodeProvider.AWS)
                               .accessKeyID(cluster.spec.cloud.aws.accessKeyId)
                               .secretAccessKey(cluster.spec.cloud.aws.secretAccessKey)
                               .vpc(cluster.spec.cloud.aws.vpcId)
                               .credential(this._presetService.preset)
                               .subnets(cluster.spec.cloud.dc)));
    }
  }
}
