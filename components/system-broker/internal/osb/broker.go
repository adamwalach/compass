/*
 * Copyright 2020 The Compass Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package osb

import "github.com/kyma-incubator/compass/components/system-broker/pkg/types"

type GqlClientForBroker interface {
	types.ApplicationsLister
	types.BundleCredentialsFetcher
	types.BundleCredentialsFetcherForInstance
	types.BundleCredentialsCreateRequester
	types.BundleCredentialsDeleteRequester
}

func NewSystemBroker(client GqlClientForBroker, ORDServiceURL string) *SystemBroker {
	return &SystemBroker{
		CatalogEndpoint:               NewCatalogEndpoint(client, &CatalogConverter{ORDServiceURL: ORDServiceURL}),
		ProvisionEndpoint:             NewProvisionEndpoint(),
		DeprovisionEndpoint:           NewDeprovisionEndpoint(),
		UpdateInstanceEndpoint:        NewUpdateInstanceEndpoint(),
		GetInstanceEndpoint:           NewGetInstanceEndpoint(),
		InstanceLastOperationEndpoint: NewInstanceLastOperationEndpoint(),
		BindEndpoint:                  NewBindEndpoint(client, client),
		UnbindEndpoint:                NewUnbindEndpoint(client, client),
		GetBindingEndpoint:            NewGetBindingEndpoint(client),
		BindLastOperationEndpoint:     NewBindLastOperationEndpoint(client),
	}
}

type SystemBroker struct {
	*CatalogEndpoint
	*ProvisionEndpoint
	*DeprovisionEndpoint
	*UpdateInstanceEndpoint
	*GetInstanceEndpoint
	*InstanceLastOperationEndpoint
	*BindEndpoint
	*UnbindEndpoint
	*GetBindingEndpoint
	*BindLastOperationEndpoint
}
