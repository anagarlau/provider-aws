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

type AgentUpdateStatus string

const (
	AgentUpdateStatus_PENDING  AgentUpdateStatus = "PENDING"
	AgentUpdateStatus_STAGING  AgentUpdateStatus = "STAGING"
	AgentUpdateStatus_STAGED   AgentUpdateStatus = "STAGED"
	AgentUpdateStatus_UPDATING AgentUpdateStatus = "UPDATING"
	AgentUpdateStatus_UPDATED  AgentUpdateStatus = "UPDATED"
	AgentUpdateStatus_FAILED   AgentUpdateStatus = "FAILED"
)

type ApplicationProtocol string

const (
	ApplicationProtocol_http  ApplicationProtocol = "http"
	ApplicationProtocol_http2 ApplicationProtocol = "http2"
	ApplicationProtocol_grpc  ApplicationProtocol = "grpc"
)

type AssignPublicIP string

const (
	AssignPublicIP_ENABLED  AssignPublicIP = "ENABLED"
	AssignPublicIP_DISABLED AssignPublicIP = "DISABLED"
)

type CPUArchitecture string

const (
	CPUArchitecture_X86_64 CPUArchitecture = "X86_64"
	CPUArchitecture_ARM64  CPUArchitecture = "ARM64"
)

type CapacityProviderField string

const (
	CapacityProviderField_TAGS CapacityProviderField = "TAGS"
)

type CapacityProviderStatus string

const (
	CapacityProviderStatus_ACTIVE   CapacityProviderStatus = "ACTIVE"
	CapacityProviderStatus_INACTIVE CapacityProviderStatus = "INACTIVE"
)

type CapacityProviderUpdateStatus string

const (
	CapacityProviderUpdateStatus_DELETE_IN_PROGRESS CapacityProviderUpdateStatus = "DELETE_IN_PROGRESS"
	CapacityProviderUpdateStatus_DELETE_COMPLETE    CapacityProviderUpdateStatus = "DELETE_COMPLETE"
	CapacityProviderUpdateStatus_DELETE_FAILED      CapacityProviderUpdateStatus = "DELETE_FAILED"
	CapacityProviderUpdateStatus_UPDATE_IN_PROGRESS CapacityProviderUpdateStatus = "UPDATE_IN_PROGRESS"
	CapacityProviderUpdateStatus_UPDATE_COMPLETE    CapacityProviderUpdateStatus = "UPDATE_COMPLETE"
	CapacityProviderUpdateStatus_UPDATE_FAILED      CapacityProviderUpdateStatus = "UPDATE_FAILED"
)

type ClusterField string

const (
	ClusterField_ATTACHMENTS    ClusterField = "ATTACHMENTS"
	ClusterField_CONFIGURATIONS ClusterField = "CONFIGURATIONS"
	ClusterField_SETTINGS       ClusterField = "SETTINGS"
	ClusterField_STATISTICS     ClusterField = "STATISTICS"
	ClusterField_TAGS           ClusterField = "TAGS"
)

type ClusterSettingName string

const (
	ClusterSettingName_containerInsights ClusterSettingName = "containerInsights"
)

type Compatibility string

const (
	Compatibility_EC2      Compatibility = "EC2"
	Compatibility_FARGATE  Compatibility = "FARGATE"
	Compatibility_EXTERNAL Compatibility = "EXTERNAL"
)

type Connectivity string

const (
	Connectivity_CONNECTED    Connectivity = "CONNECTED"
	Connectivity_DISCONNECTED Connectivity = "DISCONNECTED"
)

type ContainerCondition string

const (
	ContainerCondition_START    ContainerCondition = "START"
	ContainerCondition_COMPLETE ContainerCondition = "COMPLETE"
	ContainerCondition_SUCCESS  ContainerCondition = "SUCCESS"
	ContainerCondition_HEALTHY  ContainerCondition = "HEALTHY"
)

type ContainerInstanceField string

const (
	ContainerInstanceField_TAGS                      ContainerInstanceField = "TAGS"
	ContainerInstanceField_CONTAINER_INSTANCE_HEALTH ContainerInstanceField = "CONTAINER_INSTANCE_HEALTH"
)

type ContainerInstanceStatus string

const (
	ContainerInstanceStatus_ACTIVE              ContainerInstanceStatus = "ACTIVE"
	ContainerInstanceStatus_DRAINING            ContainerInstanceStatus = "DRAINING"
	ContainerInstanceStatus_REGISTERING         ContainerInstanceStatus = "REGISTERING"
	ContainerInstanceStatus_DEREGISTERING       ContainerInstanceStatus = "DEREGISTERING"
	ContainerInstanceStatus_REGISTRATION_FAILED ContainerInstanceStatus = "REGISTRATION_FAILED"
)

type DeploymentControllerType string

const (
	DeploymentControllerType_ECS         DeploymentControllerType = "ECS"
	DeploymentControllerType_CODE_DEPLOY DeploymentControllerType = "CODE_DEPLOY"
	DeploymentControllerType_EXTERNAL    DeploymentControllerType = "EXTERNAL"
)

type DeploymentRolloutState string

const (
	DeploymentRolloutState_COMPLETED   DeploymentRolloutState = "COMPLETED"
	DeploymentRolloutState_FAILED      DeploymentRolloutState = "FAILED"
	DeploymentRolloutState_IN_PROGRESS DeploymentRolloutState = "IN_PROGRESS"
)

type DesiredStatus string

const (
	DesiredStatus_RUNNING DesiredStatus = "RUNNING"
	DesiredStatus_PENDING DesiredStatus = "PENDING"
	DesiredStatus_STOPPED DesiredStatus = "STOPPED"
)

type DeviceCgroupPermission string

const (
	DeviceCgroupPermission_read  DeviceCgroupPermission = "read"
	DeviceCgroupPermission_write DeviceCgroupPermission = "write"
	DeviceCgroupPermission_mknod DeviceCgroupPermission = "mknod"
)

type EFSAuthorizationConfigIAM string

const (
	EFSAuthorizationConfigIAM_ENABLED  EFSAuthorizationConfigIAM = "ENABLED"
	EFSAuthorizationConfigIAM_DISABLED EFSAuthorizationConfigIAM = "DISABLED"
)

type EFSTransitEncryption string

const (
	EFSTransitEncryption_ENABLED  EFSTransitEncryption = "ENABLED"
	EFSTransitEncryption_DISABLED EFSTransitEncryption = "DISABLED"
)

type EnvironmentFileType string

const (
	EnvironmentFileType_s3 EnvironmentFileType = "s3"
)

type ExecuteCommandLogging string

const (
	ExecuteCommandLogging_NONE     ExecuteCommandLogging = "NONE"
	ExecuteCommandLogging_DEFAULT  ExecuteCommandLogging = "DEFAULT"
	ExecuteCommandLogging_OVERRIDE ExecuteCommandLogging = "OVERRIDE"
)

type FirelensConfigurationType string

const (
	FirelensConfigurationType_fluentd   FirelensConfigurationType = "fluentd"
	FirelensConfigurationType_fluentbit FirelensConfigurationType = "fluentbit"
)

type HealthStatus string

const (
	HealthStatus_HEALTHY   HealthStatus = "HEALTHY"
	HealthStatus_UNHEALTHY HealthStatus = "UNHEALTHY"
	HealthStatus_UNKNOWN   HealthStatus = "UNKNOWN"
)

type IPCMode string

const (
	IPCMode_host IPCMode = "host"
	IPCMode_task IPCMode = "task"
	IPCMode_none IPCMode = "none"
)

type InstanceHealthCheckState string

const (
	InstanceHealthCheckState_OK                InstanceHealthCheckState = "OK"
	InstanceHealthCheckState_IMPAIRED          InstanceHealthCheckState = "IMPAIRED"
	InstanceHealthCheckState_INSUFFICIENT_DATA InstanceHealthCheckState = "INSUFFICIENT_DATA"
	InstanceHealthCheckState_INITIALIZING      InstanceHealthCheckState = "INITIALIZING"
)

type InstanceHealthCheckType string

const (
	InstanceHealthCheckType_CONTAINER_RUNTIME InstanceHealthCheckType = "CONTAINER_RUNTIME"
)

type LaunchType string

const (
	LaunchType_EC2      LaunchType = "EC2"
	LaunchType_FARGATE  LaunchType = "FARGATE"
	LaunchType_EXTERNAL LaunchType = "EXTERNAL"
)

type LogDriver string

const (
	LogDriver_json_file   LogDriver = "json-file"
	LogDriver_syslog      LogDriver = "syslog"
	LogDriver_journald    LogDriver = "journald"
	LogDriver_gelf        LogDriver = "gelf"
	LogDriver_fluentd     LogDriver = "fluentd"
	LogDriver_awslogs     LogDriver = "awslogs"
	LogDriver_splunk      LogDriver = "splunk"
	LogDriver_awsfirelens LogDriver = "awsfirelens"
)

type ManagedAgentName string

const (
	ManagedAgentName_ExecuteCommandAgent ManagedAgentName = "ExecuteCommandAgent"
)

type ManagedScalingStatus string

const (
	ManagedScalingStatus_ENABLED  ManagedScalingStatus = "ENABLED"
	ManagedScalingStatus_DISABLED ManagedScalingStatus = "DISABLED"
)

type ManagedTerminationProtection string

const (
	ManagedTerminationProtection_ENABLED  ManagedTerminationProtection = "ENABLED"
	ManagedTerminationProtection_DISABLED ManagedTerminationProtection = "DISABLED"
)

type NetworkMode string

const (
	NetworkMode_bridge NetworkMode = "bridge"
	NetworkMode_host   NetworkMode = "host"
	NetworkMode_awsvpc NetworkMode = "awsvpc"
	NetworkMode_none   NetworkMode = "none"
)

type OSFamily string

const (
	OSFamily_WINDOWS_SERVER_2019_FULL OSFamily = "WINDOWS_SERVER_2019_FULL"
	OSFamily_WINDOWS_SERVER_2019_CORE OSFamily = "WINDOWS_SERVER_2019_CORE"
	OSFamily_WINDOWS_SERVER_2016_FULL OSFamily = "WINDOWS_SERVER_2016_FULL"
	OSFamily_WINDOWS_SERVER_2004_CORE OSFamily = "WINDOWS_SERVER_2004_CORE"
	OSFamily_WINDOWS_SERVER_2022_CORE OSFamily = "WINDOWS_SERVER_2022_CORE"
	OSFamily_WINDOWS_SERVER_2022_FULL OSFamily = "WINDOWS_SERVER_2022_FULL"
	OSFamily_WINDOWS_SERVER_20H2_CORE OSFamily = "WINDOWS_SERVER_20H2_CORE"
	OSFamily_LINUX                    OSFamily = "LINUX"
)

type PIDMode string

const (
	PIDMode_host PIDMode = "host"
	PIDMode_task PIDMode = "task"
)

type PlacementConstraintType string

const (
	PlacementConstraintType_distinctInstance PlacementConstraintType = "distinctInstance"
	PlacementConstraintType_memberOf         PlacementConstraintType = "memberOf"
)

type PlacementStrategyType string

const (
	PlacementStrategyType_random  PlacementStrategyType = "random"
	PlacementStrategyType_spread  PlacementStrategyType = "spread"
	PlacementStrategyType_binpack PlacementStrategyType = "binpack"
)

type PlatformDeviceType string

const (
	PlatformDeviceType_GPU PlatformDeviceType = "GPU"
)

type PropagateTags string

const (
	PropagateTags_TASK_DEFINITION PropagateTags = "TASK_DEFINITION"
	PropagateTags_SERVICE         PropagateTags = "SERVICE"
	PropagateTags_NONE            PropagateTags = "NONE"
)

type ProxyConfigurationType string

const (
	ProxyConfigurationType_APPMESH ProxyConfigurationType = "APPMESH"
)

type ResourceType string

const (
	ResourceType_GPU                  ResourceType = "GPU"
	ResourceType_InferenceAccelerator ResourceType = "InferenceAccelerator"
)

type ScaleUnit string

const (
	ScaleUnit_PERCENT ScaleUnit = "PERCENT"
)

type SchedulingStrategy string

const (
	SchedulingStrategy_REPLICA SchedulingStrategy = "REPLICA"
	SchedulingStrategy_DAEMON  SchedulingStrategy = "DAEMON"
)

type Scope string

const (
	Scope_task   Scope = "task"
	Scope_shared Scope = "shared"
)

type ServiceField string

const (
	ServiceField_TAGS ServiceField = "TAGS"
)

type SettingName string

const (
	SettingName_serviceLongArnFormat            SettingName = "serviceLongArnFormat"
	SettingName_taskLongArnFormat               SettingName = "taskLongArnFormat"
	SettingName_containerInstanceLongArnFormat  SettingName = "containerInstanceLongArnFormat"
	SettingName_awsvpcTrunking                  SettingName = "awsvpcTrunking"
	SettingName_containerInsights               SettingName = "containerInsights"
	SettingName_fargateFIPSMode                 SettingName = "fargateFIPSMode"
	SettingName_tagResourceAuthorization        SettingName = "tagResourceAuthorization"
	SettingName_fargateTaskRetirementWaitPeriod SettingName = "fargateTaskRetirementWaitPeriod"
)

type SortOrder string

const (
	SortOrder_ASC  SortOrder = "ASC"
	SortOrder_DESC SortOrder = "DESC"
)

type StabilityStatus string

const (
	StabilityStatus_STEADY_STATE StabilityStatus = "STEADY_STATE"
	StabilityStatus_STABILIZING  StabilityStatus = "STABILIZING"
)

type TargetType string

const (
	TargetType_container_instance TargetType = "container-instance"
)

type TaskDefinitionFamilyStatus string

const (
	TaskDefinitionFamilyStatus_ACTIVE   TaskDefinitionFamilyStatus = "ACTIVE"
	TaskDefinitionFamilyStatus_INACTIVE TaskDefinitionFamilyStatus = "INACTIVE"
	TaskDefinitionFamilyStatus_ALL      TaskDefinitionFamilyStatus = "ALL"
)

type TaskDefinitionField string

const (
	TaskDefinitionField_TAGS TaskDefinitionField = "TAGS"
)

type TaskDefinitionPlacementConstraintType string

const (
	TaskDefinitionPlacementConstraintType_memberOf TaskDefinitionPlacementConstraintType = "memberOf"
)

type TaskDefinitionStatus_SDK string

const (
	TaskDefinitionStatus_SDK_ACTIVE             TaskDefinitionStatus_SDK = "ACTIVE"
	TaskDefinitionStatus_SDK_INACTIVE           TaskDefinitionStatus_SDK = "INACTIVE"
	TaskDefinitionStatus_SDK_DELETE_IN_PROGRESS TaskDefinitionStatus_SDK = "DELETE_IN_PROGRESS"
)

type TaskField string

const (
	TaskField_TAGS TaskField = "TAGS"
)

type TaskSetField string

const (
	TaskSetField_TAGS TaskSetField = "TAGS"
)

type TaskStopCode string

const (
	TaskStopCode_TaskFailedToStart         TaskStopCode = "TaskFailedToStart"
	TaskStopCode_EssentialContainerExited  TaskStopCode = "EssentialContainerExited"
	TaskStopCode_UserInitiated             TaskStopCode = "UserInitiated"
	TaskStopCode_ServiceSchedulerInitiated TaskStopCode = "ServiceSchedulerInitiated"
	TaskStopCode_SpotInterruption          TaskStopCode = "SpotInterruption"
	TaskStopCode_TerminationNotice         TaskStopCode = "TerminationNotice"
)

type TransportProtocol string

const (
	TransportProtocol_tcp TransportProtocol = "tcp"
	TransportProtocol_udp TransportProtocol = "udp"
)

type UlimitName string

const (
	UlimitName_core       UlimitName = "core"
	UlimitName_cpu        UlimitName = "cpu"
	UlimitName_data       UlimitName = "data"
	UlimitName_fsize      UlimitName = "fsize"
	UlimitName_locks      UlimitName = "locks"
	UlimitName_memlock    UlimitName = "memlock"
	UlimitName_msgqueue   UlimitName = "msgqueue"
	UlimitName_nice       UlimitName = "nice"
	UlimitName_nofile     UlimitName = "nofile"
	UlimitName_nproc      UlimitName = "nproc"
	UlimitName_rss        UlimitName = "rss"
	UlimitName_rtprio     UlimitName = "rtprio"
	UlimitName_rttime     UlimitName = "rttime"
	UlimitName_sigpending UlimitName = "sigpending"
	UlimitName_stack      UlimitName = "stack"
)
