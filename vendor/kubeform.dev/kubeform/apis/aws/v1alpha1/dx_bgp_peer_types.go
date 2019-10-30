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

type DxBGPPeer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxBGPPeerSpec   `json:"spec,omitempty"`
	Status            DxBGPPeerStatus `json:"status,omitempty"`
}

type DxBGPPeerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AddressFamily string `json:"addressFamily" tf:"address_family"`
	// +optional
	AmazonAddress string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`
	// +optional
	AwsDevice string `json:"awsDevice,omitempty" tf:"aws_device,omitempty"`
	BgpAsn    int    `json:"bgpAsn" tf:"bgp_asn"`
	// +optional
	BgpAuthKey string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`
	// +optional
	BgpPeerID string `json:"bgpPeerID,omitempty" tf:"bgp_peer_id,omitempty"`
	// +optional
	BgpStatus string `json:"bgpStatus,omitempty" tf:"bgp_status,omitempty"`
	// +optional
	CustomerAddress    string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`
	VirtualInterfaceID string `json:"virtualInterfaceID" tf:"virtual_interface_id"`
}

type DxBGPPeerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DxBGPPeerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxBGPPeerList is a list of DxBGPPeers
type DxBGPPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxBGPPeer CRD objects
	Items []DxBGPPeer `json:"items,omitempty"`
}
