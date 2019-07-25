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

type GuarddutyThreatintelset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GuarddutyThreatintelsetSpec   `json:"spec,omitempty"`
	Status            GuarddutyThreatintelsetStatus `json:"status,omitempty"`
}

type GuarddutyThreatintelsetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Activate   bool   `json:"activate" tf:"activate"`
	DetectorID string `json:"detectorID" tf:"detector_id"`
	Format     string `json:"format" tf:"format"`
	Location   string `json:"location" tf:"location"`
	Name       string `json:"name" tf:"name"`
}

type GuarddutyThreatintelsetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GuarddutyThreatintelsetList is a list of GuarddutyThreatintelsets
type GuarddutyThreatintelsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GuarddutyThreatintelset CRD objects
	Items []GuarddutyThreatintelset `json:"items,omitempty"`
}