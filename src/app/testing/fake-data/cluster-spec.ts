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

import {MasterVersion} from '@shared/entity/cluster';

export function masterVersionsFake(): MasterVersion[] {
  return [
    {version: '1.8.0'},
    {version: '1.8.1'},
    {version: '1.8.2'},
    {version: '1.8.3'},
    {version: '1.8.4'},
    {version: '1.8.5'},
    {version: '1.8.6'},
    {version: '1.8.7'},
    {version: '1.8.8'},
    {version: '1.8.9'},
    {version: '1.8.10'},
    {version: '1.8.11'},
    {version: '1.8.12'},
    {version: '1.8.13'},
    {version: '1.9.0'},
  ];
}
