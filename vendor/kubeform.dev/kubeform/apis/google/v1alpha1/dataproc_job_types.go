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

// Code generated by Kubeform. DO NOT EDIT.

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

type DataprocJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataprocJobSpec   `json:"spec,omitempty"`
	Status            DataprocJobStatus `json:"status,omitempty"`
}

type DataprocJobSpecHadoopConfigLoggingConfig struct {
	// Optional. The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.
	// +optional
	DriverLogLevels map[string]string `json:"driverLogLevels,omitempty" tf:"driver_log_levels,omitempty"`
}

type DataprocJobSpecHadoopConfig struct {
	// +optional
	ArchiveUris []string `json:"archiveUris,omitempty" tf:"archive_uris,omitempty"`
	// +optional
	Args []string `json:"args,omitempty" tf:"args,omitempty"`
	// +optional
	FileUris []string `json:"fileUris,omitempty" tf:"file_uris,omitempty"`
	// +optional
	JarFileUris []string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`
	// The runtime logging config of the job
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingConfig []DataprocJobSpecHadoopConfigLoggingConfig `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`
	// +optional
	MainClass string `json:"mainClass,omitempty" tf:"main_class,omitempty"`
	// +optional
	MainJarFileURI string `json:"mainJarFileURI,omitempty" tf:"main_jar_file_uri,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
}

type DataprocJobSpecHiveConfig struct {
	// +optional
	ContinueOnFailure bool `json:"continueOnFailure,omitempty" tf:"continue_on_failure,omitempty"`
	// +optional
	JarFileUris []string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
	// +optional
	QueryFileURI string `json:"queryFileURI,omitempty" tf:"query_file_uri,omitempty"`
	// +optional
	QueryList []string `json:"queryList,omitempty" tf:"query_list,omitempty"`
	// +optional
	ScriptVariables map[string]string `json:"scriptVariables,omitempty" tf:"script_variables,omitempty"`
}

type DataprocJobSpecPigConfigLoggingConfig struct {
	// Optional. The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.
	// +optional
	DriverLogLevels map[string]string `json:"driverLogLevels,omitempty" tf:"driver_log_levels,omitempty"`
}

type DataprocJobSpecPigConfig struct {
	// +optional
	ContinueOnFailure bool `json:"continueOnFailure,omitempty" tf:"continue_on_failure,omitempty"`
	// +optional
	JarFileUris []string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`
	// The runtime logging config of the job
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingConfig []DataprocJobSpecPigConfigLoggingConfig `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
	// +optional
	QueryFileURI string `json:"queryFileURI,omitempty" tf:"query_file_uri,omitempty"`
	// +optional
	QueryList []string `json:"queryList,omitempty" tf:"query_list,omitempty"`
	// +optional
	ScriptVariables map[string]string `json:"scriptVariables,omitempty" tf:"script_variables,omitempty"`
}

type DataprocJobSpecPlacement struct {
	// The name of the cluster where the job will be submitted
	ClusterName string `json:"clusterName" tf:"cluster_name"`
	// Output-only. A cluster UUID generated by the Cloud Dataproc service when the job is submitted
	// +optional
	ClusterUUID string `json:"clusterUUID,omitempty" tf:"cluster_uuid,omitempty"`
}

type DataprocJobSpecPysparkConfigLoggingConfig struct {
	// Optional. The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.
	// +optional
	DriverLogLevels map[string]string `json:"driverLogLevels,omitempty" tf:"driver_log_levels,omitempty"`
}

type DataprocJobSpecPysparkConfig struct {
	// Optional. HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip
	// +optional
	ArchiveUris []string `json:"archiveUris,omitempty" tf:"archive_uris,omitempty"`
	// Optional. The arguments to pass to the driver. Do not include arguments, such as --conf, that can be set as job properties, since a collision may occur that causes an incorrect job submission
	// +optional
	Args []string `json:"args,omitempty" tf:"args,omitempty"`
	// Optional. HCFS URIs of files to be copied to the working directory of Python drivers and distributed tasks. Useful for naively parallel tasks
	// +optional
	FileUris []string `json:"fileUris,omitempty" tf:"file_uris,omitempty"`
	// Optional. HCFS URIs of jar files to add to the CLASSPATHs of the Python driver and tasks
	// +optional
	JarFileUris []string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`
	// The runtime logging config of the job
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingConfig []DataprocJobSpecPysparkConfigLoggingConfig `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`
	// Required. The HCFS URI of the main Python file to use as the driver. Must be a .py file
	MainPythonFileURI string `json:"mainPythonFileURI" tf:"main_python_file_uri"`
	// Optional. A mapping of property names to values, used to configure PySpark. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
	// Optional. HCFS file URIs of Python files to pass to the PySpark framework. Supported file types: .py, .egg, and .zip
	// +optional
	PythonFileUris []string `json:"pythonFileUris,omitempty" tf:"python_file_uris,omitempty"`
}

type DataprocJobSpecReference struct {
	// The job ID, which must be unique within the project. The job ID is generated by the server upon job submission or provided by the user as a means to perform retries without creating duplicate jobs
	// +optional
	JobID string `json:"jobID,omitempty" tf:"job_id,omitempty"`
}

type DataprocJobSpecScheduling struct {
	// Maximum number of times per hour a driver may be restarted as a result of driver terminating with non-zero code before job is reported failed.
	// +optional
	MaxFailuresPerHour int64 `json:"maxFailuresPerHour,omitempty" tf:"max_failures_per_hour,omitempty"`
}

type DataprocJobSpecSparkConfigLoggingConfig struct {
	// Optional. The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.
	// +optional
	DriverLogLevels map[string]string `json:"driverLogLevels,omitempty" tf:"driver_log_levels,omitempty"`
}

type DataprocJobSpecSparkConfig struct {
	// +optional
	ArchiveUris []string `json:"archiveUris,omitempty" tf:"archive_uris,omitempty"`
	// +optional
	Args []string `json:"args,omitempty" tf:"args,omitempty"`
	// +optional
	FileUris []string `json:"fileUris,omitempty" tf:"file_uris,omitempty"`
	// +optional
	JarFileUris []string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`
	// The runtime logging config of the job
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingConfig []DataprocJobSpecSparkConfigLoggingConfig `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`
	// +optional
	MainClass string `json:"mainClass,omitempty" tf:"main_class,omitempty"`
	// +optional
	MainJarFileURI string `json:"mainJarFileURI,omitempty" tf:"main_jar_file_uri,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
}

type DataprocJobSpecSparksqlConfigLoggingConfig struct {
	// Optional. The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.
	// +optional
	DriverLogLevels map[string]string `json:"driverLogLevels,omitempty" tf:"driver_log_levels,omitempty"`
}

type DataprocJobSpecSparksqlConfig struct {
	// +optional
	JarFileUris []string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`
	// The runtime logging config of the job
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingConfig []DataprocJobSpecSparksqlConfigLoggingConfig `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
	// +optional
	QueryFileURI string `json:"queryFileURI,omitempty" tf:"query_file_uri,omitempty"`
	// +optional
	QueryList []string `json:"queryList,omitempty" tf:"query_list,omitempty"`
	// +optional
	ScriptVariables map[string]string `json:"scriptVariables,omitempty" tf:"script_variables,omitempty"`
}

type DataprocJobSpecStatus struct {
	// Output-only. Optional job state details, such as an error description if the state is ERROR
	// +optional
	Details string `json:"details,omitempty" tf:"details,omitempty"`
	// Output-only. A state message specifying the overall job state
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
	// Output-only. The time when this state was entered
	// +optional
	StateStartTime string `json:"stateStartTime,omitempty" tf:"state_start_time,omitempty"`
	// Output-only. Additional state information, which includes status reported by the agent
	// +optional
	Substate string `json:"substate,omitempty" tf:"substate,omitempty"`
}

type DataprocJobSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Output-only. If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
	// +optional
	DriverControlsFilesURI string `json:"driverControlsFilesURI,omitempty" tf:"driver_controls_files_uri,omitempty"`
	// Output-only. A URI pointing to the location of the stdout of the job's driver program
	// +optional
	DriverOutputResourceURI string `json:"driverOutputResourceURI,omitempty" tf:"driver_output_resource_uri,omitempty"`
	// +optional
	ForceDelete bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HadoopConfig []DataprocJobSpecHadoopConfig `json:"hadoopConfig,omitempty" tf:"hadoop_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HiveConfig []DataprocJobSpecHiveConfig `json:"hiveConfig,omitempty" tf:"hive_config,omitempty"`
	// Optional. The labels to associate with this job.
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PigConfig []DataprocJobSpecPigConfig `json:"pigConfig,omitempty" tf:"pig_config,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Placement []DataprocJobSpecPlacement `json:"placement" tf:"placement"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PysparkConfig []DataprocJobSpecPysparkConfig `json:"pysparkConfig,omitempty" tf:"pyspark_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Reference []DataprocJobSpecReference `json:"reference,omitempty" tf:"reference,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// Optional. Job scheduling configuration.
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Scheduling []DataprocJobSpecScheduling `json:"scheduling,omitempty" tf:"scheduling,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SparkConfig []DataprocJobSpecSparkConfig `json:"sparkConfig,omitempty" tf:"spark_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SparksqlConfig []DataprocJobSpecSparksqlConfig `json:"sparksqlConfig,omitempty" tf:"sparksql_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Status []DataprocJobSpecStatus `json:"status,omitempty" tf:"status,omitempty"`
}

type DataprocJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DataprocJobSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataprocJobList is a list of DataprocJobs
type DataprocJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataprocJob CRD objects
	Items []DataprocJob `json:"items,omitempty"`
}
