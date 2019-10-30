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

type Codepipeline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodepipelineSpec   `json:"spec,omitempty"`
	Status            CodepipelineStatus `json:"status,omitempty"`
}

type CodepipelineSpecArtifactStoreEncryptionKey struct {
	ID   string `json:"ID" tf:"id"`
	Type string `json:"type" tf:"type"`
}

type CodepipelineSpecArtifactStore struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EncryptionKey []CodepipelineSpecArtifactStoreEncryptionKey `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`
	Location      string                                       `json:"location" tf:"location"`
	Type          string                                       `json:"type" tf:"type"`
}

type CodepipelineSpecStageAction struct {
	Category string `json:"category" tf:"category"`
	// +optional
	Configuration map[string]string `json:"configuration,omitempty" tf:"configuration,omitempty"`
	// +optional
	InputArtifacts []string `json:"inputArtifacts,omitempty" tf:"input_artifacts,omitempty"`
	Name           string   `json:"name" tf:"name"`
	// +optional
	OutputArtifacts []string `json:"outputArtifacts,omitempty" tf:"output_artifacts,omitempty"`
	Owner           string   `json:"owner" tf:"owner"`
	Provider        string   `json:"provider" tf:"provider"`
	// +optional
	RoleArn string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
	// +optional
	RunOrder int    `json:"runOrder,omitempty" tf:"run_order,omitempty"`
	Version  string `json:"version" tf:"version"`
}

type CodepipelineSpecStage struct {
	Action []CodepipelineSpecStageAction `json:"action" tf:"action"`
	Name   string                        `json:"name" tf:"name"`
}

type CodepipelineSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	ArtifactStore []CodepipelineSpecArtifactStore `json:"artifactStore" tf:"artifact_store"`
	Name          string                          `json:"name" tf:"name"`
	RoleArn       string                          `json:"roleArn" tf:"role_arn"`
	// +kubebuilder:validation:MinItems=2
	Stage []CodepipelineSpecStage `json:"stage" tf:"stage"`
}

type CodepipelineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CodepipelineSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CodepipelineList is a list of Codepipelines
type CodepipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Codepipeline CRD objects
	Items []Codepipeline `json:"items,omitempty"`
}
