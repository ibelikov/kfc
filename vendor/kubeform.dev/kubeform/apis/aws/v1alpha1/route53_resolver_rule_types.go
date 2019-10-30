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

type Route53ResolverRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ResolverRuleSpec   `json:"spec,omitempty"`
	Status            Route53ResolverRuleStatus `json:"status,omitempty"`
}

type Route53ResolverRuleSpecTargetIP struct {
	Ip string `json:"ip" tf:"ip"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
}

type Route53ResolverRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn        string `json:"arn,omitempty" tf:"arn,omitempty"`
	DomainName string `json:"domainName" tf:"domain_name"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	OwnerID string `json:"ownerID,omitempty" tf:"owner_id,omitempty"`
	// +optional
	ResolverEndpointID string `json:"resolverEndpointID,omitempty" tf:"resolver_endpoint_id,omitempty"`
	RuleType           string `json:"ruleType" tf:"rule_type"`
	// +optional
	ShareStatus string `json:"shareStatus,omitempty" tf:"share_status,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TargetIP []Route53ResolverRuleSpecTargetIP `json:"targetIP,omitempty" tf:"target_ip,omitempty"`
}

type Route53ResolverRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *Route53ResolverRuleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Route53ResolverRuleList is a list of Route53ResolverRules
type Route53ResolverRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Route53ResolverRule CRD objects
	Items []Route53ResolverRule `json:"items,omitempty"`
}
