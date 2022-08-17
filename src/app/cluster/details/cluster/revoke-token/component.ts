// Copyright 2020 The Kubermatic Kubernetes Platform contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import {Component, Input, OnInit} from '@angular/core';
import {MatDialogRef} from '@angular/material/dialog';
import {NotificationService} from '@core/services/notification';
import {UserService} from '@core/services/user';
import {Cluster, Token} from '@shared/entity/cluster';
import {View} from '@shared/entity/common';
import {Member} from '@shared/entity/member';
import {GroupConfig} from '@shared/model/Config';
import {MemberUtils, Permission} from '@shared/utils/member';
import {take} from 'rxjs/operators';
import {ClusterService} from '@core/services/cluster';
import {Observable} from 'rxjs';

@Component({
  selector: 'km-revoke-token',
  templateUrl: './template.html',
  styleUrls: ['./style.scss'],
})
export class RevokeTokenComponent implements OnInit {
  @Input() cluster: Cluster;
  @Input() projectID: string;
  revokeAdminToken = false;
  revokeViewerToken = false;
  private _user: Member;
  private _currentGroupConfig: GroupConfig;

  constructor(
    private readonly _matDialogRef: MatDialogRef<RevokeTokenComponent>,
    private readonly _notificationService: NotificationService,
    private readonly _userService: UserService,
    private readonly _clusterService: ClusterService
  ) {}

  ngOnInit(): void {
    this._userService.currentUser.pipe(take(1)).subscribe(user => (this._user = user));

    this._userService.getCurrentUserGroup(this.projectID).subscribe(userGroup => {
      this._currentGroupConfig = this._userService.getCurrentUserGroupConfig(userGroup);
    });
  }

  isRevokeTokenEnabled(): boolean {
    return MemberUtils.hasPermission(this._user, this._currentGroupConfig, View.Clusters, Permission.Edit);
  }

  getObservable(): Observable<Token> | void {
    if (this.revokeAdminToken) {
      return this._clusterService.editToken(this.cluster, this.projectID, {token: ''});
    }

    if (this.revokeViewerToken) {
      return this._clusterService.editViewerToken(this.cluster, this.projectID, {token: ''});
    }
  }

  onNext(token: Token): void {
    if (this.revokeAdminToken) {
      this._notificationService.success(`Revoked the admin token from the ${this.cluster.name} cluster`);
      this._matDialogRef.close(token);
    }
    if (this.revokeViewerToken) {
      this._notificationService.success(`Revoked the viewer token from the ${this.cluster.name} cluster`);
      this._matDialogRef.close(token);
    }
  }
}