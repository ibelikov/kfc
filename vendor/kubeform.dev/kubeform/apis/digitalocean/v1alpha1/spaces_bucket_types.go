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

type SpacesBucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpacesBucketSpec   `json:"spec,omitempty"`
	Status            SpacesBucketStatus `json:"status,omitempty"`
}

type SpacesBucketSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Canned ACL applied on bucket creation
	// +optional
	Acl string `json:"acl,omitempty" tf:"acl,omitempty"`
	// The FQDN of the bucket
	// +optional
	BucketDomainName string `json:"bucketDomainName,omitempty" tf:"bucket_domain_name,omitempty"`
	// Unless true, the bucket will only be destroyed if empty
	// +optional
	ForceDestroy bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`
	// Bucket name
	Name string `json:"name" tf:"name"`
	// Bucket region
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// the uniform resource name for the bucket
	// +optional
	Urn string `json:"urn,omitempty" tf:"urn,omitempty"`
}

type SpacesBucketStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SpacesBucketSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SpacesBucketList is a list of SpacesBuckets
type SpacesBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpacesBucket CRD objects
	Items []SpacesBucket `json:"items,omitempty"`
}
