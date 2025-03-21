/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

// +kubebuilder:skipversion
type AWSDomainInformation struct {
	// The name of an OpenSearch Service domain. Domain names are unique across
	// the domains owned by an account within an Amazon Web Services Region.
	DomainName *string `json:"domainName,omitempty"`
}

// +kubebuilder:skipversion
type AccessPoliciesStatus struct {
	// Access policy rules for an Amazon OpenSearch Service domain endpoint. For
	// more information, see Configuring access policies (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-access-policies).
	// The maximum size of a policy document is 100 KB.
	Options *string `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type AdvancedOptionsStatus struct {
	// Exposes native OpenSearch configuration values from opensearch.yml. The following
	// advanced options are available:
	//
	//    * Allows references to indexes in an HTTP request body. Must be false
	//    when configuring access to individual sub-resources. Default is true.
	//
	//    * Specifies the percentage of heap space allocated to field data. Default
	//    is unbounded.
	//
	// For more information, see Advanced cluster parameters (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options).
	Options map[string]*string `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type AdvancedSecurityOptions struct {
	AnonymousAuthDisableDate *metav1.Time `json:"anonymousAuthDisableDate,omitempty"`

	AnonymousAuthEnabled *bool `json:"anonymousAuthEnabled,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	InternalUserDatabaseEnabled *bool `json:"internalUserDatabaseEnabled,omitempty"`
	// Describes the SAML application configured for the domain.
	SAMLOptions *SAMLOptionsOutput `json:"sAMLOptions,omitempty"`
}

// +kubebuilder:skipversion
type AdvancedSecurityOptionsInput struct {
	AnonymousAuthEnabled *bool `json:"anonymousAuthEnabled,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	InternalUserDatabaseEnabled *bool `json:"internalUserDatabaseEnabled,omitempty"`
	// Credentials for the master user for a domain.
	MasterUserOptions *MasterUserOptions `json:"masterUserOptions,omitempty"`
	// The SAML authentication configuration for an Amazon OpenSearch Service domain.
	SAMLOptions *SAMLOptionsInput `json:"sAMLOptions,omitempty"`
}

// +kubebuilder:skipversion
type AdvancedSecurityOptionsStatus struct {
	// Container for fine-grained access control settings.
	Options *AdvancedSecurityOptions `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type AuthorizedPrincipal struct {
	Principal *string `json:"principal,omitempty"`
}

// +kubebuilder:skipversion
type AutoTuneMaintenanceSchedule struct {
	CronExpressionForRecurrence *string `json:"cronExpressionForRecurrence,omitempty"`
	// The duration of a maintenance schedule. For more information, see Auto-Tune
	// for Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html).
	Duration *Duration `json:"duration,omitempty"`

	StartAt *metav1.Time `json:"startAt,omitempty"`
}

// +kubebuilder:skipversion
type AutoTuneOptions struct {
	// The Auto-Tune desired state. Valid values are ENABLED and DISABLED.
	DesiredState *string `json:"desiredState,omitempty"`

	MaintenanceSchedules []*AutoTuneMaintenanceSchedule `json:"maintenanceSchedules,omitempty"`

	UseOffPeakWindow *bool `json:"useOffPeakWindow,omitempty"`
}

// +kubebuilder:skipversion
type AutoTuneOptionsInput struct {
	// The Auto-Tune desired state. Valid values are ENABLED and DISABLED.
	DesiredState *string `json:"desiredState,omitempty"`

	MaintenanceSchedules []*AutoTuneMaintenanceSchedule `json:"maintenanceSchedules,omitempty"`

	UseOffPeakWindow *bool `json:"useOffPeakWindow,omitempty"`
}

// +kubebuilder:skipversion
type AutoTuneOptionsOutput struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// The Auto-Tune state for the domain. For valid states see Auto-Tune for Amazon
	// OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html).
	State *string `json:"state,omitempty"`

	UseOffPeakWindow *bool `json:"useOffPeakWindow,omitempty"`
}

// +kubebuilder:skipversion
type AutoTuneStatus struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`

	PendingDeletion *bool `json:"pendingDeletion,omitempty"`
	// The Auto-Tune state for the domain. For valid states see Auto-Tune for Amazon
	// OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html).
	State *string `json:"state,omitempty"`
}

// +kubebuilder:skipversion
type ChangeProgressDetails struct {
	ChangeID *string `json:"changeID,omitempty"`

	Message *string `json:"message,omitempty"`
}

// +kubebuilder:skipversion
type ChangeProgressStatusDetails struct {
	ChangeID *string `json:"changeID,omitempty"`

	CompletedProperties []*string `json:"completedProperties,omitempty"`

	PendingProperties []*string `json:"pendingProperties,omitempty"`
}

// +kubebuilder:skipversion
type ClusterConfig struct {
	// Container for the parameters required to enable cold storage for an OpenSearch
	// Service domain. For more information, see Cold storage for Amazon OpenSearch
	// Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cold-storage.html).
	ColdStorageOptions *ColdStorageOptions `json:"coldStorageOptions,omitempty"`

	DedicatedMasterCount *int64 `json:"dedicatedMasterCount,omitempty"`

	DedicatedMasterEnabled *bool `json:"dedicatedMasterEnabled,omitempty"`

	DedicatedMasterType *string `json:"dedicatedMasterType,omitempty"`

	InstanceCount *int64 `json:"instanceCount,omitempty"`

	InstanceType *string `json:"instanceType,omitempty"`

	MultiAZWithStandbyEnabled *bool `json:"multiAZWithStandbyEnabled,omitempty"`

	WarmCount *int64 `json:"warmCount,omitempty"`

	WarmEnabled *bool `json:"warmEnabled,omitempty"`

	WarmType *string `json:"warmType,omitempty"`
	// The zone awareness configuration for an Amazon OpenSearch Service domain.
	ZoneAwarenessConfig *ZoneAwarenessConfig `json:"zoneAwarenessConfig,omitempty"`

	ZoneAwarenessEnabled *bool `json:"zoneAwarenessEnabled,omitempty"`
}

// +kubebuilder:skipversion
type ClusterConfigStatus struct {
	// Container for the cluster configuration of an OpenSearch Service domain.
	// For more information, see Creating and managing Amazon OpenSearch Service
	// domains (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html).
	Options *ClusterConfig `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type CognitoOptions struct {
	Enabled *bool `json:"enabled,omitempty"`

	IdentityPoolID *string `json:"identityPoolID,omitempty"`

	RoleARN *string `json:"roleARN,omitempty"`

	UserPoolID *string `json:"userPoolID,omitempty"`
}

// +kubebuilder:skipversion
type CognitoOptionsStatus struct {
	// Container for the parameters required to enable Cognito authentication for
	// an OpenSearch Service domain. For more information, see Configuring Amazon
	// Cognito authentication for OpenSearch Dashboards (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html).
	Options *CognitoOptions `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type ColdStorageOptions struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// +kubebuilder:skipversion
type CompatibleVersionsMap struct {
	SourceVersion *string `json:"sourceVersion,omitempty"`
}

// +kubebuilder:skipversion
type DomainConfig struct {
	// Container for information about a configuration change happening on a domain.
	ChangeProgressDetails *ChangeProgressDetails `json:"changeProgressDetails,omitempty"`
}

// +kubebuilder:skipversion
type DomainEndpointOptions struct {
	CustomEndpoint *string `json:"customEndpoint,omitempty"`
	// The Amazon Resource Name (ARN) of the domain. See Identifiers for IAM Entities
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/index.html) in Using AWS
	// Identity and Access Management for more information.
	CustomEndpointCertificateARN *string `json:"customEndpointCertificateARN,omitempty"`

	CustomEndpointEnabled *bool `json:"customEndpointEnabled,omitempty"`

	EnforceHTTPS *bool `json:"enforceHTTPS,omitempty"`

	TLSSecurityPolicy *string `json:"tlsSecurityPolicy,omitempty"`
}

// +kubebuilder:skipversion
type DomainEndpointOptionsStatus struct {
	// Options to configure a custom endpoint for an OpenSearch Service domain.
	Options *DomainEndpointOptions `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type DomainInfo struct {
	// The name of an OpenSearch Service domain. Domain names are unique across
	// the domains owned by an account within an Amazon Web Services Region.
	DomainName *string `json:"domainName,omitempty"`
}

// +kubebuilder:skipversion
type DomainMaintenanceDetails struct {
	// The name of an OpenSearch Service domain. Domain names are unique across
	// the domains owned by an account within an Amazon Web Services Region.
	DomainName *string `json:"domainName,omitempty"`
}

// +kubebuilder:skipversion
type DomainNodesStatus struct {
	InstanceType *string `json:"instanceType,omitempty"`
	// The type of EBS volume that a domain uses. For more information, see Configuring
	// EBS-based storage (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/opensearch-createupdatedomains.html#opensearch-createdomain-configure-ebs).
	StorageVolumeType *string `json:"storageVolumeType,omitempty"`
}

// +kubebuilder:skipversion
type DomainPackageDetails struct {
	// The name of an OpenSearch Service domain. Domain names are unique across
	// the domains owned by an account within an Amazon Web Services Region.
	DomainName *string `json:"domainName,omitempty"`
}

// +kubebuilder:skipversion
type DomainStatus_SDK struct {
	// The Amazon Resource Name (ARN) of the domain. See Identifiers for IAM Entities
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/index.html) in Using AWS
	// Identity and Access Management for more information.
	ARN *string `json:"arn,omitempty"`
	// Access policy rules for an Amazon OpenSearch Service domain endpoint. For
	// more information, see Configuring access policies (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-access-policies).
	// The maximum size of a policy document is 100 KB.
	AccessPolicies *string `json:"accessPolicies,omitempty"`
	// Exposes native OpenSearch configuration values from opensearch.yml. The following
	// advanced options are available:
	//
	//    * Allows references to indexes in an HTTP request body. Must be false
	//    when configuring access to individual sub-resources. Default is true.
	//
	//    * Specifies the percentage of heap space allocated to field data. Default
	//    is unbounded.
	//
	// For more information, see Advanced cluster parameters (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options).
	AdvancedOptions map[string]*string `json:"advancedOptions,omitempty"`
	// Container for fine-grained access control settings.
	AdvancedSecurityOptions *AdvancedSecurityOptions `json:"advancedSecurityOptions,omitempty"`
	// The Auto-Tune settings for a domain, displayed when enabling or disabling
	// Auto-Tune.
	AutoTuneOptions *AutoTuneOptionsOutput `json:"autoTuneOptions,omitempty"`
	// Container for information about a configuration change happening on a domain.
	ChangeProgressDetails *ChangeProgressDetails `json:"changeProgressDetails,omitempty"`
	// Container for the cluster configuration of an OpenSearch Service domain.
	// For more information, see Creating and managing Amazon OpenSearch Service
	// domains (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html).
	ClusterConfig *ClusterConfig `json:"clusterConfig,omitempty"`
	// Container for the parameters required to enable Cognito authentication for
	// an OpenSearch Service domain. For more information, see Configuring Amazon
	// Cognito authentication for OpenSearch Dashboards (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html).
	CognitoOptions *CognitoOptions `json:"cognitoOptions,omitempty"`

	Created *bool `json:"created,omitempty"`

	Deleted *bool `json:"deleted,omitempty"`
	// Options to configure a custom endpoint for an OpenSearch Service domain.
	DomainEndpointOptions *DomainEndpointOptions `json:"domainEndpointOptions,omitempty"`
	// Unique identifier for an OpenSearch Service domain.
	DomainID *string `json:"domainID,omitempty"`
	// The name of an OpenSearch Service domain. Domain names are unique across
	// the domains owned by an account within an Amazon Web Services Region.
	DomainName *string `json:"domainName,omitempty"`
	// Container for the parameters required to enable EBS-based storage for an
	// OpenSearch Service domain.
	EBSOptions *EBSOptions `json:"ebsOptions,omitempty"`
	// Specifies whether the domain should encrypt data at rest, and if so, the
	// Key Management Service (KMS) key to use. Can be used only to create a new
	// domain, not update an existing one.
	EncryptionAtRestOptions *EncryptionAtRestOptions `json:"encryptionAtRestOptions,omitempty"`
	// The domain endpoint to which index and search requests are submitted. For
	// example, search-imdb-movies-oopcnjfn6ugo.eu-west-1.es.amazonaws.com or doc-imdb-movies-oopcnjfn6u.eu-west-1.es.amazonaws.com.
	Endpoint *string `json:"endpoint,omitempty"`
	// The domain endpoint to which index and search requests are submitted. For
	// example, search-imdb-movies-oopcnjfn6ugo.eu-west-1.es.amazonaws.com or doc-imdb-movies-oopcnjfn6u.eu-west-1.es.amazonaws.com.
	EndpointV2 *string `json:"endpointV2,omitempty"`

	Endpoints map[string]*string `json:"endpoints,omitempty"`

	EngineVersion *string `json:"engineVersion,omitempty"`

	IPAddressType *string `json:"iPAddressType,omitempty"`

	LogPublishingOptions map[string]*LogPublishingOption `json:"logPublishingOptions,omitempty"`
	// Enables or disables node-to-node encryption. For more information, see Node-to-node
	// encryption for Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ntn.html).
	NodeToNodeEncryptionOptions *NodeToNodeEncryptionOptions `json:"nodeToNodeEncryptionOptions,omitempty"`
	// Options for a domain's off-peak window (https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_OffPeakWindow.html),
	// during which OpenSearch Service can perform mandatory configuration changes
	// on the domain.
	OffPeakWindowOptions *OffPeakWindowOptions `json:"offPeakWindowOptions,omitempty"`

	Processing *bool `json:"processing,omitempty"`
	// The current status of the service software for an Amazon OpenSearch Service
	// domain. For more information, see Service software updates in Amazon OpenSearch
	// Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/service-software.html).
	ServiceSoftwareOptions *ServiceSoftwareOptions `json:"serviceSoftwareOptions,omitempty"`
	// The time, in UTC format, when OpenSearch Service takes a daily automated
	// snapshot of the specified domain. Default is 0 hours.
	SnapshotOptions *SnapshotOptions `json:"snapshotOptions,omitempty"`
	// Options for configuring service software updates for a domain.
	SoftwareUpdateOptions *SoftwareUpdateOptions `json:"softwareUpdateOptions,omitempty"`

	UpgradeProcessing *bool `json:"upgradeProcessing,omitempty"`
	// Information about the subnets and security groups for an Amazon OpenSearch
	// Service domain provisioned within a virtual private cloud (VPC). For more
	// information, see Launching your Amazon OpenSearch Service domains using a
	// VPC (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html).
	// This information only exists if the domain was created with VPCOptions.
	VPCOptions *VPCDerivedInfo `json:"vpcOptions,omitempty"`
}

// +kubebuilder:skipversion
type DryRunProgressStatus struct {
	CreationDate *string `json:"creationDate,omitempty"`

	DryRunID *string `json:"dryRunID,omitempty"`

	DryRunStatus *string `json:"dryRunStatus,omitempty"`

	UpdateDate *string `json:"updateDate,omitempty"`
}

// +kubebuilder:skipversion
type DryRunResults struct {
	Message *string `json:"message,omitempty"`
}

// +kubebuilder:skipversion
type Duration struct {
	// The unit of a maintenance schedule duration. Valid value is HOUR.
	Unit *string `json:"unit,omitempty"`
	// Integer that specifies the value of a maintenance schedule duration.
	Value *int64 `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type EBSOptions struct {
	EBSEnabled *bool `json:"ebsEnabled,omitempty"`

	IOPS *int64 `json:"iops,omitempty"`

	Throughput *int64 `json:"throughput,omitempty"`

	VolumeSize *int64 `json:"volumeSize,omitempty"`
	// The type of EBS volume that a domain uses. For more information, see Configuring
	// EBS-based storage (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/opensearch-createupdatedomains.html#opensearch-createdomain-configure-ebs).
	VolumeType *string `json:"volumeType,omitempty"`
}

// +kubebuilder:skipversion
type EBSOptionsStatus struct {
	// Container for the parameters required to enable EBS-based storage for an
	// OpenSearch Service domain.
	Options *EBSOptions `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type EncryptionAtRestOptions struct {
	Enabled *bool `json:"enabled,omitempty"`

	KMSKeyID *string `json:"kmsKeyID,omitempty"`
}

// +kubebuilder:skipversion
type EncryptionAtRestOptionsStatus struct {
	// Specifies whether the domain should encrypt data at rest, and if so, the
	// Key Management Service (KMS) key to use. Can be used only to create a new
	// domain, not update an existing one.
	Options *EncryptionAtRestOptions `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type IPAddressTypeStatus struct {
	Options *string `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type InstanceTypeDetails struct {
	AdvancedSecurityEnabled *bool `json:"advancedSecurityEnabled,omitempty"`

	AppLogsEnabled *bool `json:"appLogsEnabled,omitempty"`

	CognitoEnabled *bool `json:"cognitoEnabled,omitempty"`

	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty"`

	InstanceType *string `json:"instanceType,omitempty"`

	WarmEnabled *bool `json:"warmEnabled,omitempty"`
}

// +kubebuilder:skipversion
type LogPublishingOption struct {
	// ARN of the Cloudwatch log group to publish logs to.
	CloudWatchLogsLogGroupARN *string `json:"cloudWatchLogsLogGroupARN,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`
}

// +kubebuilder:skipversion
type LogPublishingOptionsStatus struct {
	Options map[string]*LogPublishingOption `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type MasterUserOptions struct {
	// The Amazon Resource Name (ARN) of the domain. See Identifiers for IAM Entities
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/index.html) in Using AWS
	// Identity and Access Management for more information.
	MasterUserARN *string `json:"masterUserARN,omitempty"`

	MasterUserName *string `json:"masterUserName,omitempty"`

	MasterUserPassword *string `json:"masterUserPassword,omitempty"`
}

// +kubebuilder:skipversion
type NodeToNodeEncryptionOptions struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// +kubebuilder:skipversion
type NodeToNodeEncryptionOptionsStatus struct {
	// Enables or disables node-to-node encryption. For more information, see Node-to-node
	// encryption for Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ntn.html).
	Options *NodeToNodeEncryptionOptions `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type OffPeakWindow struct {
	// The desired start time for an off-peak maintenance window (https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_OffPeakWindow.html).
	WindowStartTime *WindowStartTime `json:"windowStartTime,omitempty"`
}

// +kubebuilder:skipversion
type OffPeakWindowOptions struct {
	Enabled *bool `json:"enabled,omitempty"`
	// A custom 10-hour, low-traffic window during which OpenSearch Service can
	// perform mandatory configuration changes on the domain. These actions can
	// include scheduled service software updates and blue/green Auto-Tune enhancements.
	// OpenSearch Service will schedule these actions during the window that you
	// specify.
	//
	// If you don't specify a window start time, it defaults to 10:00 P.M. local
	// time.
	//
	// For more information, see Defining off-peak maintenance windows for Amazon
	// OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/off-peak.html).
	OffPeakWindow *OffPeakWindow `json:"offPeakWindow,omitempty"`
}

// +kubebuilder:skipversion
type OffPeakWindowOptionsStatus struct {
	// Options for a domain's off-peak window (https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_OffPeakWindow.html),
	// during which OpenSearch Service can perform mandatory configuration changes
	// on the domain.
	Options *OffPeakWindowOptions `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type OptionStatus struct {
	PendingDeletion *bool `json:"pendingDeletion,omitempty"`
}

// +kubebuilder:skipversion
type RecurringCharge struct {
	RecurringChargeFrequency *string `json:"recurringChargeFrequency,omitempty"`
}

// +kubebuilder:skipversion
type ReservedInstance struct {
	CurrencyCode *string `json:"currencyCode,omitempty"`

	InstanceType *string `json:"instanceType,omitempty"`

	ReservedInstanceID *string `json:"reservedInstanceID,omitempty"`

	ReservedInstanceOfferingID *string `json:"reservedInstanceOfferingID,omitempty"`

	State *string `json:"state,omitempty"`
}

// +kubebuilder:skipversion
type ReservedInstanceOffering struct {
	CurrencyCode *string `json:"currencyCode,omitempty"`

	InstanceType *string `json:"instanceType,omitempty"`

	ReservedInstanceOfferingID *string `json:"reservedInstanceOfferingID,omitempty"`
}

// +kubebuilder:skipversion
type SAMLIDp struct {
	EntityID *string `json:"entityID,omitempty"`

	MetadataContent *string `json:"metadataContent,omitempty"`
}

// +kubebuilder:skipversion
type SAMLOptionsInput struct {
	Enabled *bool `json:"enabled,omitempty"`
	// The SAML identity povider information.
	IDp *SAMLIDp `json:"idp,omitempty"`

	MasterBackendRole *string `json:"masterBackendRole,omitempty"`

	MasterUserName *string `json:"masterUserName,omitempty"`

	RolesKey *string `json:"rolesKey,omitempty"`

	SessionTimeoutMinutes *int64 `json:"sessionTimeoutMinutes,omitempty"`

	SubjectKey *string `json:"subjectKey,omitempty"`
}

// +kubebuilder:skipversion
type SAMLOptionsOutput struct {
	Enabled *bool `json:"enabled,omitempty"`
	// The SAML identity povider information.
	IDp *SAMLIDp `json:"idp,omitempty"`

	RolesKey *string `json:"rolesKey,omitempty"`

	SessionTimeoutMinutes *int64 `json:"sessionTimeoutMinutes,omitempty"`

	SubjectKey *string `json:"subjectKey,omitempty"`
}

// +kubebuilder:skipversion
type ScheduledAction struct {
	Cancellable *bool `json:"cancellable,omitempty"`

	Description *string `json:"description,omitempty"`

	ID *string `json:"id,omitempty"`

	Mandatory *bool `json:"mandatory,omitempty"`
}

// +kubebuilder:skipversion
type ServiceSoftwareOptions struct {
	AutomatedUpdateDate *metav1.Time `json:"automatedUpdateDate,omitempty"`

	Cancellable *bool `json:"cancellable,omitempty"`

	CurrentVersion *string `json:"currentVersion,omitempty"`

	Description *string `json:"description,omitempty"`

	NewVersion *string `json:"newVersion,omitempty"`

	OptionalDeployment *bool `json:"optionalDeployment,omitempty"`

	UpdateAvailable *bool `json:"updateAvailable,omitempty"`

	UpdateStatus *string `json:"updateStatus,omitempty"`
}

// +kubebuilder:skipversion
type SnapshotOptions struct {
	AutomatedSnapshotStartHour *int64 `json:"automatedSnapshotStartHour,omitempty"`
}

// +kubebuilder:skipversion
type SnapshotOptionsStatus struct {
	// The time, in UTC format, when OpenSearch Service takes a daily automated
	// snapshot of the specified domain. Default is 0 hours.
	Options *SnapshotOptions `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type SoftwareUpdateOptions struct {
	AutoSoftwareUpdateEnabled *bool `json:"autoSoftwareUpdateEnabled,omitempty"`
}

// +kubebuilder:skipversion
type SoftwareUpdateOptionsStatus struct {
	// Options for configuring service software updates for a domain.
	Options *SoftwareUpdateOptions `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type Tag struct {
	// A string between 1 to 128 characters that specifies the key for a tag. Tag
	// keys must be unique for the domain to which they're attached.
	Key *string `json:"key,omitempty"`
	// A string between 0 to 256 characters that specifies the value for a tag.
	// Tag values can be null and don't have to be unique in a tag set.
	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type VPCDerivedInfo struct {
	AvailabilityZones []*string `json:"availabilityZones,omitempty"`

	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`

	SubnetIDs []*string `json:"subnetIDs,omitempty"`

	VPCID *string `json:"vpcID,omitempty"`
}

// +kubebuilder:skipversion
type VPCDerivedInfoStatus struct {
	// Information about the subnets and security groups for an Amazon OpenSearch
	// Service domain provisioned within a virtual private cloud (VPC). For more
	// information, see Launching your Amazon OpenSearch Service domains using a
	// VPC (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html).
	// This information only exists if the domain was created with VPCOptions.
	Options *VPCDerivedInfo `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type VPCEndpoint struct {
	// Information about the subnets and security groups for an Amazon OpenSearch
	// Service domain provisioned within a virtual private cloud (VPC). For more
	// information, see Launching your Amazon OpenSearch Service domains using a
	// VPC (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html).
	// This information only exists if the domain was created with VPCOptions.
	VPCOptions *VPCDerivedInfo `json:"vpcOptions,omitempty"`
}

// +kubebuilder:skipversion
type VPCEndpointError struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// +kubebuilder:skipversion
type VPCEndpointSummary struct {
	VPCEndpointOwner *string `json:"vpcEndpointOwner,omitempty"`
}

// +kubebuilder:skipversion
type VPCOptions struct {
	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`

	SubnetIDs []*string `json:"subnetIDs,omitempty"`
}

// +kubebuilder:skipversion
type ValidationFailure struct {
	Code *string `json:"code,omitempty"`

	Message *string `json:"message,omitempty"`
}

// +kubebuilder:skipversion
type VersionStatus struct {
	Options *string `json:"options,omitempty"`
}

// +kubebuilder:skipversion
type WindowStartTime struct {
	Hours *int64 `json:"hours,omitempty"`

	Minutes *int64 `json:"minutes,omitempty"`
}

// +kubebuilder:skipversion
type ZoneAwarenessConfig struct {
	AvailabilityZoneCount *int64 `json:"availabilityZoneCount,omitempty"`
}
