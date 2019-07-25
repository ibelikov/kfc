package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type WafregionalWebACL struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalWebACLSpec   `json:"spec,omitempty"`
	Status            WafregionalWebACLStatus `json:"status,omitempty"`
}

type WafregionalWebACLSpecDefaultAction struct {
	Type string `json:"type" tf:"type"`
}

type WafregionalWebACLSpecLoggingConfigurationRedactedFieldsFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty" tf:"data,omitempty"`
	Type string `json:"type" tf:"type"`
}

type WafregionalWebACLSpecLoggingConfigurationRedactedFields struct {
	// +kubebuilder:validation:UniqueItems=true
	FieldToMatch []WafregionalWebACLSpecLoggingConfigurationRedactedFieldsFieldToMatch `json:"fieldToMatch" tf:"field_to_match"`
}

type WafregionalWebACLSpecLoggingConfiguration struct {
	LogDestination string `json:"logDestination" tf:"log_destination"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RedactedFields []WafregionalWebACLSpecLoggingConfigurationRedactedFields `json:"redactedFields,omitempty" tf:"redacted_fields,omitempty"`
}

type WafregionalWebACLSpecRuleAction struct {
	Type string `json:"type" tf:"type"`
}

type WafregionalWebACLSpecRuleOverrideAction struct {
	Type string `json:"type" tf:"type"`
}

type WafregionalWebACLSpecRule struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Action []WafregionalWebACLSpecRuleAction `json:"action,omitempty" tf:"action,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OverrideAction []WafregionalWebACLSpecRuleOverrideAction `json:"overrideAction,omitempty" tf:"override_action,omitempty"`
	Priority       int                                       `json:"priority" tf:"priority"`
	RuleID         string                                    `json:"ruleID" tf:"rule_id"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type WafregionalWebACLSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +kubebuilder:validation:MaxItems=1
	DefaultAction []WafregionalWebACLSpecDefaultAction `json:"defaultAction" tf:"default_action"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingConfiguration []WafregionalWebACLSpecLoggingConfiguration `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`
	MetricName           string                                      `json:"metricName" tf:"metric_name"`
	Name                 string                                      `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Rule []WafregionalWebACLSpecRule `json:"rule,omitempty" tf:"rule,omitempty"`
}

type WafregionalWebACLStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafregionalWebACLList is a list of WafregionalWebACLs
type WafregionalWebACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafregionalWebACL CRD objects
	Items []WafregionalWebACL `json:"items,omitempty"`
}