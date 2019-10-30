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

type PostgresqlServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlServerSpec   `json:"spec,omitempty"`
	Status            PostgresqlServerStatus `json:"status,omitempty"`
}

type PostgresqlServerSpecSku struct {
	Capacity int    `json:"capacity" tf:"capacity"`
	Family   string `json:"family" tf:"family"`
	Name     string `json:"name" tf:"name"`
	Tier     string `json:"tier" tf:"tier"`
}

type PostgresqlServerSpecStorageProfile struct {
	// +optional
	BackupRetentionDays int `json:"backupRetentionDays,omitempty" tf:"backup_retention_days,omitempty"`
	// +optional
	GeoRedundantBackup string `json:"geoRedundantBackup,omitempty" tf:"geo_redundant_backup,omitempty"`
	StorageMb          int    `json:"storageMb" tf:"storage_mb"`
}

type PostgresqlServerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	AdministratorLogin         string `json:"administratorLogin" tf:"administrator_login"`
	AdministratorLoginPassword string `json:"-" sensitive:"true" tf:"administrator_login_password"`
	// +optional
	Fqdn              string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
	Location          string `json:"location" tf:"location"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku            []PostgresqlServerSpecSku `json:"sku" tf:"sku"`
	SslEnforcement string                    `json:"sslEnforcement" tf:"ssl_enforcement"`
	// +kubebuilder:validation:MaxItems=1
	StorageProfile []PostgresqlServerSpecStorageProfile `json:"storageProfile" tf:"storage_profile"`
	// +optional
	Tags    map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Version string            `json:"version" tf:"version"`
}

type PostgresqlServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *PostgresqlServerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PostgresqlServerList is a list of PostgresqlServers
type PostgresqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PostgresqlServer CRD objects
	Items []PostgresqlServer `json:"items,omitempty"`
}
