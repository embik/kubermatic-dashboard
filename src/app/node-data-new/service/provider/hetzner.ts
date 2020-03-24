import {Observable} from 'rxjs';
import {filter, switchMap} from 'rxjs/operators';

import {PresetsService} from '../../../core/services';
import {HetznerTypes} from '../../../shared/entity/provider/hetzner/TypeEntity';
import {NodeProvider} from '../../../shared/model/NodeProviderConstants';
import {ClusterService} from '../../../wizard-new/service/cluster';
import {NodeDataMode} from '../../config';
import {NodeDataService} from '../service';

export class NodeDataHetznerProvider {
  constructor(
      private readonly _nodeDataService: NodeDataService, private readonly _clusterService: ClusterService,
      private readonly _presetService: PresetsService) {}

  flavors(): Observable<HetznerTypes> {
    // TODO: support dialog mode
    switch (this._nodeDataService.mode) {
      case NodeDataMode.Wizard:
        return this._clusterService.clusterChanges
            .pipe(filter(_ => this._clusterService.provider === NodeProvider.HETZNER))
            .pipe(switchMap(
                cluster => this._presetService.provider(NodeProvider.HETZNER)
                               .token(cluster.spec.cloud.hetzner.token)
                               .credential(this._presetService.preset)
                               .flavors()));
    }
  }
}
