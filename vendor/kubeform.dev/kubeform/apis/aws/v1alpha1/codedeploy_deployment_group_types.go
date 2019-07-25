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

type CodedeployDeploymentGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodedeployDeploymentGroupSpec   `json:"spec,omitempty"`
	Status            CodedeployDeploymentGroupStatus `json:"status,omitempty"`
}

type CodedeployDeploymentGroupSpecAlarmConfiguration struct {
	// +optional
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:UniqueItems=true
	Alarms []string `json:"alarms,omitempty" tf:"alarms,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	IgnorePollAlarmFailure bool `json:"ignorePollAlarmFailure,omitempty" tf:"ignore_poll_alarm_failure,omitempty"`
}

type CodedeployDeploymentGroupSpecAutoRollbackConfiguration struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Events []string `json:"events,omitempty" tf:"events,omitempty"`
}

type CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption struct {
	// +optional
	ActionOnTimeout string `json:"actionOnTimeout,omitempty" tf:"action_on_timeout,omitempty"`
	// +optional
	WaitTimeInMinutes int `json:"waitTimeInMinutes,omitempty" tf:"wait_time_in_minutes,omitempty"`
}

type CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption struct {
	// +optional
	Action string `json:"action,omitempty" tf:"action,omitempty"`
}

type CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess struct {
	// +optional
	Action string `json:"action,omitempty" tf:"action,omitempty"`
	// +optional
	TerminationWaitTimeInMinutes int `json:"terminationWaitTimeInMinutes,omitempty" tf:"termination_wait_time_in_minutes,omitempty"`
}

type CodedeployDeploymentGroupSpecBlueGreenDeploymentConfig struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DeploymentReadyOption []CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption `json:"deploymentReadyOption,omitempty" tf:"deployment_ready_option,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	GreenFleetProvisioningOption []CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption `json:"greenFleetProvisioningOption,omitempty" tf:"green_fleet_provisioning_option,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TerminateBlueInstancesOnDeploymentSuccess []CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess `json:"terminateBlueInstancesOnDeploymentSuccess,omitempty" tf:"terminate_blue_instances_on_deployment_success,omitempty"`
}

type CodedeployDeploymentGroupSpecDeploymentStyle struct {
	// +optional
	DeploymentOption string `json:"deploymentOption,omitempty" tf:"deployment_option,omitempty"`
	// +optional
	DeploymentType string `json:"deploymentType,omitempty" tf:"deployment_type,omitempty"`
}

type CodedeployDeploymentGroupSpecEc2TagFilter struct {
	// +optional
	Key string `json:"key,omitempty" tf:"key,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	Value string `json:"value,omitempty" tf:"value,omitempty"`
}

type CodedeployDeploymentGroupSpecEc2TagSetEc2TagFilter struct {
	// +optional
	Key string `json:"key,omitempty" tf:"key,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	Value string `json:"value,omitempty" tf:"value,omitempty"`
}

type CodedeployDeploymentGroupSpecEc2TagSet struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Ec2TagFilter []CodedeployDeploymentGroupSpecEc2TagSetEc2TagFilter `json:"ec2TagFilter,omitempty" tf:"ec2_tag_filter,omitempty"`
}

type CodedeployDeploymentGroupSpecEcsService struct {
	ClusterName string `json:"clusterName" tf:"cluster_name"`
	ServiceName string `json:"serviceName" tf:"service_name"`
}

type CodedeployDeploymentGroupSpecLoadBalancerInfoElbInfo struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
}

type CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupInfo struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
}

type CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute struct {
	// +kubebuilder:validation:UniqueItems=true
	ListenerArns []string `json:"listenerArns" tf:"listener_arns"`
}

type CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTargetGroup struct {
	Name string `json:"name" tf:"name"`
}

type CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute struct {
	// +kubebuilder:validation:UniqueItems=true
	ListenerArns []string `json:"listenerArns" tf:"listener_arns"`
}

type CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	ProdTrafficRoute []CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute `json:"prodTrafficRoute" tf:"prod_traffic_route"`
	// +kubebuilder:validation:MaxItems=2
	// +kubebuilder:validation:MinItems=1
	TargetGroup []CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTargetGroup `json:"targetGroup" tf:"target_group"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TestTrafficRoute []CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute `json:"testTrafficRoute,omitempty" tf:"test_traffic_route,omitempty"`
}

type CodedeployDeploymentGroupSpecLoadBalancerInfo struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ElbInfo []CodedeployDeploymentGroupSpecLoadBalancerInfoElbInfo `json:"elbInfo,omitempty" tf:"elb_info,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	TargetGroupInfo []CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupInfo `json:"targetGroupInfo,omitempty" tf:"target_group_info,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TargetGroupPairInfo []CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo `json:"targetGroupPairInfo,omitempty" tf:"target_group_pair_info,omitempty"`
}

type CodedeployDeploymentGroupSpecOnPremisesInstanceTagFilter struct {
	// +optional
	Key string `json:"key,omitempty" tf:"key,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	Value string `json:"value,omitempty" tf:"value,omitempty"`
}

type CodedeployDeploymentGroupSpecTriggerConfiguration struct {
	// +kubebuilder:validation:UniqueItems=true
	TriggerEvents    []string `json:"triggerEvents" tf:"trigger_events"`
	TriggerName      string   `json:"triggerName" tf:"trigger_name"`
	TriggerTargetArn string   `json:"triggerTargetArn" tf:"trigger_target_arn"`
}

type CodedeployDeploymentGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	AlarmConfiguration []CodedeployDeploymentGroupSpecAlarmConfiguration `json:"alarmConfiguration,omitempty" tf:"alarm_configuration,omitempty"`
	AppName            string                                            `json:"appName" tf:"app_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AutoRollbackConfiguration []CodedeployDeploymentGroupSpecAutoRollbackConfiguration `json:"autoRollbackConfiguration,omitempty" tf:"auto_rollback_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AutoscalingGroups []string `json:"autoscalingGroups,omitempty" tf:"autoscaling_groups,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BlueGreenDeploymentConfig []CodedeployDeploymentGroupSpecBlueGreenDeploymentConfig `json:"blueGreenDeploymentConfig,omitempty" tf:"blue_green_deployment_config,omitempty"`
	// +optional
	DeploymentConfigName string `json:"deploymentConfigName,omitempty" tf:"deployment_config_name,omitempty"`
	DeploymentGroupName  string `json:"deploymentGroupName" tf:"deployment_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DeploymentStyle []CodedeployDeploymentGroupSpecDeploymentStyle `json:"deploymentStyle,omitempty" tf:"deployment_style,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Ec2TagFilter []CodedeployDeploymentGroupSpecEc2TagFilter `json:"ec2TagFilter,omitempty" tf:"ec2_tag_filter,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Ec2TagSet []CodedeployDeploymentGroupSpecEc2TagSet `json:"ec2TagSet,omitempty" tf:"ec2_tag_set,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EcsService []CodedeployDeploymentGroupSpecEcsService `json:"ecsService,omitempty" tf:"ecs_service,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoadBalancerInfo []CodedeployDeploymentGroupSpecLoadBalancerInfo `json:"loadBalancerInfo,omitempty" tf:"load_balancer_info,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	OnPremisesInstanceTagFilter []CodedeployDeploymentGroupSpecOnPremisesInstanceTagFilter `json:"onPremisesInstanceTagFilter,omitempty" tf:"on_premises_instance_tag_filter,omitempty"`
	ServiceRoleArn              string                                                     `json:"serviceRoleArn" tf:"service_role_arn"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	TriggerConfiguration []CodedeployDeploymentGroupSpecTriggerConfiguration `json:"triggerConfiguration,omitempty" tf:"trigger_configuration,omitempty"`
}

type CodedeployDeploymentGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CodedeployDeploymentGroupList is a list of CodedeployDeploymentGroups
type CodedeployDeploymentGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CodedeployDeploymentGroup CRD objects
	Items []CodedeployDeploymentGroup `json:"items,omitempty"`
}