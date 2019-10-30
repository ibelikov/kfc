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

type EfsMountTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EfsMountTargetSpec   `json:"spec,omitempty"`
	Status            EfsMountTargetStatus `json:"status,omitempty"`
}

type EfsMountTargetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DnsName string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`
	// +optional
	FileSystemArn string `json:"fileSystemArn,omitempty" tf:"file_system_arn,omitempty"`
	FileSystemID  string `json:"fileSystemID" tf:"file_system_id"`
	// +optional
	IpAddress string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`
	// +optional
	NetworkInterfaceID string `json:"networkInterfaceID,omitempty" tf:"network_interface_id,omitempty"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	SubnetID       string   `json:"subnetID" tf:"subnet_id"`
}

type EfsMountTargetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EfsMountTargetSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EfsMountTargetList is a list of EfsMountTargets
type EfsMountTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EfsMountTarget CRD objects
	Items []EfsMountTarget `json:"items,omitempty"`
}
