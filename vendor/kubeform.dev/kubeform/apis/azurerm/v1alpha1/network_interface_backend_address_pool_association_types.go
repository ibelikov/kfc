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

type NetworkInterfaceBackendAddressPoolAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceBackendAddressPoolAssociationSpec   `json:"spec,omitempty"`
	Status            NetworkInterfaceBackendAddressPoolAssociationStatus `json:"status,omitempty"`
}

type NetworkInterfaceBackendAddressPoolAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendAddressPoolID string `json:"backendAddressPoolID" tf:"backend_address_pool_id"`
	IpConfigurationName  string `json:"ipConfigurationName" tf:"ip_configuration_name"`
	NetworkInterfaceID   string `json:"networkInterfaceID" tf:"network_interface_id"`
}

type NetworkInterfaceBackendAddressPoolAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkInterfaceBackendAddressPoolAssociationList is a list of NetworkInterfaceBackendAddressPoolAssociations
type NetworkInterfaceBackendAddressPoolAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkInterfaceBackendAddressPoolAssociation CRD objects
	Items []NetworkInterfaceBackendAddressPoolAssociation `json:"items,omitempty"`
}