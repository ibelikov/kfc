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
	"kubeform.dev/kubeform/apis/aws"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = schema.GroupVersion{Group: aws.GroupName, Version: "v1alpha1"}

var (
	// TODO: move SchemeBuilder with zz_generated.deepcopy.go to k8s.io/api.
	// localSchemeBuilder and AddToScheme will stay in k8s.io/kubernetes.
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	// We only register manually written functions here. The registration of the
	// generated functions takes place in the generated files. The separation
	// makes the code compile even when the generated files are missing.
	localSchemeBuilder.Register(addKnownTypes)
}

// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,

		&AcmCertificate{},
		&AcmCertificateList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&Alb{},
		&AlbList{},

		&AlbListener{},
		&AlbListenerList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&Ami{},
		&AmiList{},

		&AmiCopy{},
		&AmiCopyList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&BackupPlan{},
		&BackupPlanList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&BackupVault{},
		&BackupVaultList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&Codepipeline{},
		&CodepipelineList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&DaxCluster{},
		&DaxClusterList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&DbInstance{},
		&DbInstanceList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&DxConnection{},
		&DxConnectionList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&DxGateway{},
		&DxGatewayList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&DxLag{},
		&DxLagList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&EcsCluster{},
		&EcsClusterList{},

		&EcsService{},
		&EcsServiceList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&Eip{},
		&EipList{},

		&EipAssociation{},
		&EipAssociationList{},

		&EksCluster{},
		&EksClusterList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&Elb{},
		&ElbList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&EmrCluster{},
		&EmrClusterList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&FlowLog{},
		&FlowLogList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&GlueJob{},
		&GlueJobList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&IamGroup{},
		&IamGroupList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&IamPolicy{},
		&IamPolicyList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&IamRole{},
		&IamRoleList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&IamUser{},
		&IamUserList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&Instance{},
		&InstanceList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&IotCertificate{},
		&IotCertificateList{},

		&IotPolicy{},
		&IotPolicyList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&IotThing{},
		&IotThingList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&IotThingType{},
		&IotThingTypeList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&KeyPair{},
		&KeyPairList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&KmsAlias{},
		&KmsAliasList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&KmsGrant{},
		&KmsGrantList{},

		&KmsKey{},
		&KmsKeyList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&Lb{},
		&LbList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&LbListener{},
		&LbListenerList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&LbSSLNegotiationPolicy{},
		&LbSSLNegotiationPolicyList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&MqBroker{},
		&MqBrokerList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&NatGateway{},
		&NatGatewayList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&NetworkACL{},
		&NetworkACLList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&PinpointApp{},
		&PinpointAppList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&RdsCluster{},
		&RdsClusterList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&Route{},
		&RouteList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&Route53Record{},
		&Route53RecordList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&RouteTable{},
		&RouteTableList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&S3Bucket{},
		&S3BucketList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&SesTemplate{},
		&SesTemplateList{},

		&SfnActivity{},
		&SfnActivityList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&SnsTopic{},
		&SnsTopicList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&SqsQueue{},
		&SqsQueueList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&SsmActivation{},
		&SsmActivationList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&SsmParameter{},
		&SsmParameterList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&Subnet{},
		&SubnetList{},

		&SwfDomain{},
		&SwfDomainList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&TransferServer{},
		&TransferServerList{},

		&TransferUser{},
		&TransferUserList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&Vpc{},
		&VpcList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&WafIpset{},
		&WafIpsetList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&WafRule{},
		&WafRuleList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&WafWebACL{},
		&WafWebACLList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
