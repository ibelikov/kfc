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

type CloudformationStack struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudformationStackSpec   `json:"spec,omitempty"`
	Status            CloudformationStackStatus `json:"status,omitempty"`
}

type CloudformationStackSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Capabilities []string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`
	// +optional
	DisableRollback bool `json:"disableRollback,omitempty" tf:"disable_rollback,omitempty"`
	// +optional
	IamRoleArn string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`
	Name       string `json:"name" tf:"name"`
	// +optional
	NotificationArns []string `json:"notificationArns,omitempty" tf:"notification_arns,omitempty"`
	// +optional
	OnFailure string `json:"onFailure,omitempty" tf:"on_failure,omitempty"`
	// +optional
	Outputs map[string]string `json:"outputs,omitempty" tf:"outputs,omitempty"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`
	// +optional
	PolicyBody string `json:"policyBody,omitempty" tf:"policy_body,omitempty"`
	// +optional
	PolicyURL string `json:"policyURL,omitempty" tf:"policy_url,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TemplateBody string `json:"templateBody,omitempty" tf:"template_body,omitempty"`
	// +optional
	TemplateURL string `json:"templateURL,omitempty" tf:"template_url,omitempty"`
	// +optional
	TimeoutInMinutes int `json:"timeoutInMinutes,omitempty" tf:"timeout_in_minutes,omitempty"`
}

type CloudformationStackStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CloudformationStackSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudformationStackList is a list of CloudformationStacks
type CloudformationStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudformationStack CRD objects
	Items []CloudformationStack `json:"items,omitempty"`
}
