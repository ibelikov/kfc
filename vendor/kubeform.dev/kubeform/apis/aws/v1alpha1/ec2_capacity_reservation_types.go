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

type Ec2CapacityReservation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2CapacityReservationSpec   `json:"spec,omitempty"`
	Status            Ec2CapacityReservationStatus `json:"status,omitempty"`
}

type Ec2CapacityReservationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AvailabilityZone string `json:"availabilityZone" tf:"availability_zone"`
	// +optional
	EbsOptimized bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`
	// +optional
	EndDate string `json:"endDate,omitempty" tf:"end_date,omitempty"`
	// +optional
	EndDateType string `json:"endDateType,omitempty" tf:"end_date_type,omitempty"`
	// +optional
	EphemeralStorage bool  `json:"ephemeralStorage,omitempty" tf:"ephemeral_storage,omitempty"`
	InstanceCount    int64 `json:"instanceCount" tf:"instance_count"`
	// +optional
	InstanceMatchCriteria string `json:"instanceMatchCriteria,omitempty" tf:"instance_match_criteria,omitempty"`
	InstancePlatform      string `json:"instancePlatform" tf:"instance_platform"`
	InstanceType          string `json:"instanceType" tf:"instance_type"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Tenancy string `json:"tenancy,omitempty" tf:"tenancy,omitempty"`
}

type Ec2CapacityReservationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *Ec2CapacityReservationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Ec2CapacityReservationList is a list of Ec2CapacityReservations
type Ec2CapacityReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ec2CapacityReservation CRD objects
	Items []Ec2CapacityReservation `json:"items,omitempty"`
}
