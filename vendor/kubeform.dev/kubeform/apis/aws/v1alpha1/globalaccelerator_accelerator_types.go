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

type GlobalacceleratorAccelerator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalacceleratorAcceleratorSpec   `json:"spec,omitempty"`
	Status            GlobalacceleratorAcceleratorStatus `json:"status,omitempty"`
}

type GlobalacceleratorAcceleratorSpecAttributes struct {
	// +optional
	FlowLogsEnabled bool `json:"flowLogsEnabled,omitempty" tf:"flow_logs_enabled,omitempty"`
	// +optional
	FlowLogsS3Bucket string `json:"flowLogsS3Bucket,omitempty" tf:"flow_logs_s3_bucket,omitempty"`
	// +optional
	FlowLogsS3Prefix string `json:"flowLogsS3Prefix,omitempty" tf:"flow_logs_s3_prefix,omitempty"`
}

type GlobalacceleratorAcceleratorSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	Attributes []GlobalacceleratorAcceleratorSpecAttributes `json:"attributes,omitempty" tf:"attributes,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	IpAddressType string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`
	Name          string `json:"name" tf:"name"`
}

type GlobalacceleratorAcceleratorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlobalacceleratorAcceleratorList is a list of GlobalacceleratorAccelerators
type GlobalacceleratorAcceleratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlobalacceleratorAccelerator CRD objects
	Items []GlobalacceleratorAccelerator `json:"items,omitempty"`
}