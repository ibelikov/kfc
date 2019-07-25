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

type StreamAnalyticsOutputMssql struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsOutputMssqlSpec   `json:"spec,omitempty"`
	Status            StreamAnalyticsOutputMssqlStatus `json:"status,omitempty"`
}

type StreamAnalyticsOutputMssqlSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	Database               string `json:"database" tf:"database"`
	Name                   string `json:"name" tf:"name"`
	ResourceGroupName      string `json:"resourceGroupName" tf:"resource_group_name"`
	Server                 string `json:"server" tf:"server"`
	StreamAnalyticsJobName string `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name"`
	Table                  string `json:"table" tf:"table"`
	User                   string `json:"user" tf:"user"`
}

type StreamAnalyticsOutputMssqlStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StreamAnalyticsOutputMssqlList is a list of StreamAnalyticsOutputMssqls
type StreamAnalyticsOutputMssqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StreamAnalyticsOutputMssql CRD objects
	Items []StreamAnalyticsOutputMssql `json:"items,omitempty"`
}