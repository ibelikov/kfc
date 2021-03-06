/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"encoding/json"

	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type AzureAppService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzureAppServiceSpec   `json:"spec,omitempty"`
	Status            AzureAppServiceStatus `json:"status,omitempty"`
}

type AzureAppServiceSpec struct {
	// +optional
	SecretRef   *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	// +optional
	Source string `json:"source" tf:"source"`

	// +optional
	// Whether the App Service should always be on (Basic or above required).
	AlwaysOn json.RawMessage `json:"alwaysOn,omitempty" tf:"always_on,omitempty"`
	// +optional
	// Name of app service plan.
	AppServicePlanName json.RawMessage `json:"appServicePlanName,omitempty" tf:"app_service_plan_name,omitempty"`
	// +optional
	// App Service's configuration values.
	AppSettings json.RawMessage `json:"appSettings,omitempty" tf:"app_settings,omitempty"`
	// +optional
	// Whether the App Service only allows HTTPS connections.
	HttpsOnly json.RawMessage `json:"httpsOnly,omitempty" tf:"https_only,omitempty"`
	// +optional
	// Name of app service.
	Name json.RawMessage `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	// Resource Group location.
	RgLocation json.RawMessage `json:"rgLocation,omitempty" tf:"rg_location,omitempty"`
	// +optional
	// Resource Group name.
	RgName json.RawMessage `json:"rgName,omitempty" tf:"rg_name,omitempty"`
	// +optional
	// App Service Plan size.
	Size json.RawMessage `json:"size,omitempty" tf:"size,omitempty"`
	// +optional
	// App Service Plan tier.
	Tier json.RawMessage `json:"tier,omitempty" tf:"tier,omitempty"`
	// +optional
	// Whether the App Service should use the 32-bit worker process (needed for free plans, will be overwritten if Free tier selected).
	Use32BitWorkerProcess json.RawMessage `json:"use32BitWorkerProcess,omitempty" tf:"use_32_bit_worker_process,omitempty"`
}

type AzureAppServiceOutput struct {
	//
	// +optional
	PossibleOutboundIPAddresses string `json:"possibleOutboundIPAddresses" tf:"possible_outbound_ip_addresses"`
}

type AzureAppServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AzureAppServiceOutput `json:"output,omitempty"`
	// +optional
	State string `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzureAppServiceList is a list of AzureAppServices
type AzureAppServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzureAppService CRD objects
	Items []AzureAppService `json:"items,omitempty"`
}
