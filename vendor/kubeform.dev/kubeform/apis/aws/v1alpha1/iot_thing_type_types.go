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

type IotThingType struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IotThingTypeSpec   `json:"spec,omitempty"`
	Status            IotThingTypeStatus `json:"status,omitempty"`
}

type IotThingTypeSpecProperties struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=3
	// +kubebuilder:validation:UniqueItems=true
	SearchableAttributes []string `json:"searchableAttributes,omitempty" tf:"searchable_attributes,omitempty"`
}

type IotThingTypeSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Deprecated bool   `json:"deprecated,omitempty" tf:"deprecated,omitempty"`
	Name       string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Properties []IotThingTypeSpecProperties `json:"properties,omitempty" tf:"properties,omitempty"`
}

type IotThingTypeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IotThingTypeList is a list of IotThingTypes
type IotThingTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IotThingType CRD objects
	Items []IotThingType `json:"items,omitempty"`
}