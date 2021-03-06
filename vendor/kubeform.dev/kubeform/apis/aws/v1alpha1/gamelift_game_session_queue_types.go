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

type GameliftGameSessionQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GameliftGameSessionQueueSpec   `json:"spec,omitempty"`
	Status            GameliftGameSessionQueueStatus `json:"status,omitempty"`
}

type GameliftGameSessionQueueSpecPlayerLatencyPolicy struct {
	MaximumIndividualPlayerLatencyMilliseconds int64 `json:"maximumIndividualPlayerLatencyMilliseconds" tf:"maximum_individual_player_latency_milliseconds"`
	// +optional
	PolicyDurationSeconds int64 `json:"policyDurationSeconds,omitempty" tf:"policy_duration_seconds,omitempty"`
}

type GameliftGameSessionQueueSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	Destinations []string `json:"destinations,omitempty" tf:"destinations,omitempty"`
	Name         string   `json:"name" tf:"name"`
	// +optional
	PlayerLatencyPolicy []GameliftGameSessionQueueSpecPlayerLatencyPolicy `json:"playerLatencyPolicy,omitempty" tf:"player_latency_policy,omitempty"`
	// +optional
	TimeoutInSeconds int64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds,omitempty"`
}

type GameliftGameSessionQueueStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GameliftGameSessionQueueSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GameliftGameSessionQueueList is a list of GameliftGameSessionQueues
type GameliftGameSessionQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GameliftGameSessionQueue CRD objects
	Items []GameliftGameSessionQueue `json:"items,omitempty"`
}
