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

type ConfigConfigRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigConfigRuleSpec   `json:"spec,omitempty"`
	Status            ConfigConfigRuleStatus `json:"status,omitempty"`
}

type ConfigConfigRuleSpecScope struct {
	// +optional
	ComplianceResourceID string `json:"complianceResourceID,omitempty" tf:"compliance_resource_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	ComplianceResourceTypes []string `json:"complianceResourceTypes,omitempty" tf:"compliance_resource_types,omitempty"`
	// +optional
	TagKey string `json:"tagKey,omitempty" tf:"tag_key,omitempty"`
	// +optional
	TagValue string `json:"tagValue,omitempty" tf:"tag_value,omitempty"`
}

type ConfigConfigRuleSpecSourceSourceDetail struct {
	// +optional
	EventSource string `json:"eventSource,omitempty" tf:"event_source,omitempty"`
	// +optional
	MaximumExecutionFrequency string `json:"maximumExecutionFrequency,omitempty" tf:"maximum_execution_frequency,omitempty"`
	// +optional
	MessageType string `json:"messageType,omitempty" tf:"message_type,omitempty"`
}

type ConfigConfigRuleSpecSource struct {
	Owner string `json:"owner" tf:"owner"`
	// +optional
	// +kubebuilder:validation:MaxItems=25
	SourceDetail     []ConfigConfigRuleSpecSourceSourceDetail `json:"sourceDetail,omitempty" tf:"source_detail,omitempty"`
	SourceIdentifier string                                   `json:"sourceIdentifier" tf:"source_identifier"`
}

type ConfigConfigRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	InputParameters string `json:"inputParameters,omitempty" tf:"input_parameters,omitempty"`
	// +optional
	MaximumExecutionFrequency string `json:"maximumExecutionFrequency,omitempty" tf:"maximum_execution_frequency,omitempty"`
	Name                      string `json:"name" tf:"name"`
	// +optional
	RuleID string `json:"ruleID,omitempty" tf:"rule_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Scope []ConfigConfigRuleSpecScope `json:"scope,omitempty" tf:"scope,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Source []ConfigConfigRuleSpecSource `json:"source" tf:"source"`
}

type ConfigConfigRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ConfigConfigRuleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ConfigConfigRuleList is a list of ConfigConfigRules
type ConfigConfigRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConfigConfigRule CRD objects
	Items []ConfigConfigRule `json:"items,omitempty"`
}
