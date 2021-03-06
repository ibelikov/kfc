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

type VpcPeeringConnectionOptions struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcPeeringConnectionOptionsSpec   `json:"spec,omitempty"`
	Status            VpcPeeringConnectionOptionsStatus `json:"status,omitempty"`
}

type VpcPeeringConnectionOptionsSpecAccepter struct {
	// +optional
	AllowClassicLinkToRemoteVpc bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`
	// +optional
	AllowRemoteVpcDNSResolution bool `json:"allowRemoteVpcDNSResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`
	// +optional
	AllowVpcToRemoteClassicLink bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

type VpcPeeringConnectionOptionsSpecRequester struct {
	// +optional
	AllowClassicLinkToRemoteVpc bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc,omitempty"`
	// +optional
	AllowRemoteVpcDNSResolution bool `json:"allowRemoteVpcDNSResolution,omitempty" tf:"allow_remote_vpc_dns_resolution,omitempty"`
	// +optional
	AllowVpcToRemoteClassicLink bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link,omitempty"`
}

type VpcPeeringConnectionOptionsSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	Accepter []VpcPeeringConnectionOptionsSpecAccepter `json:"accepter,omitempty" tf:"accepter,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Requester              []VpcPeeringConnectionOptionsSpecRequester `json:"requester,omitempty" tf:"requester,omitempty"`
	VpcPeeringConnectionID string                                     `json:"vpcPeeringConnectionID" tf:"vpc_peering_connection_id"`
}

type VpcPeeringConnectionOptionsStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VpcPeeringConnectionOptionsSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcPeeringConnectionOptionsList is a list of VpcPeeringConnectionOptionss
type VpcPeeringConnectionOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcPeeringConnectionOptions CRD objects
	Items []VpcPeeringConnectionOptions `json:"items,omitempty"`
}
