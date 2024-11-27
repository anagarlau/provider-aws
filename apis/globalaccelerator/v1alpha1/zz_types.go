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
type AcceleratorAttributes struct {
	FlowLogsEnabled *bool `json:"flowLogsEnabled,omitempty"`

	FlowLogsS3Bucket *string `json:"flowLogsS3Bucket,omitempty"`

	FlowLogsS3Prefix *string `json:"flowLogsS3Prefix,omitempty"`
}

// +kubebuilder:skipversion
type AcceleratorEvent struct {
	Message *string `json:"message,omitempty"`

	Timestamp *metav1.Time `json:"timestamp,omitempty"`
}

// +kubebuilder:skipversion
type Accelerator_SDK struct {
	AcceleratorARN *string `json:"acceleratorARN,omitempty"`

	CreatedTime *metav1.Time `json:"createdTime,omitempty"`

	DNSName *string `json:"dnsName,omitempty"`

	DualStackDNSName *string `json:"dualStackDNSName,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	Events []*AcceleratorEvent `json:"events,omitempty"`

	IPAddressType *string `json:"ipAddressType,omitempty"`

	IPSets []*IPSet `json:"ipSets,omitempty"`

	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`
}

// +kubebuilder:skipversion
type Attachment struct {
	AttachmentARN *string `json:"attachmentARN,omitempty"`

	CreatedTime *metav1.Time `json:"createdTime,omitempty"`

	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
}

// +kubebuilder:skipversion
type ByoipCIDR struct {
	CIDR *string `json:"cidr,omitempty"`
}

// +kubebuilder:skipversion
type ByoipCIDREvent struct {
	Message *string `json:"message,omitempty"`

	Timestamp *metav1.Time `json:"timestamp,omitempty"`
}

// +kubebuilder:skipversion
type CIDRAuthorizationContext struct {
	Message *string `json:"message,omitempty"`

	Signature *string `json:"signature,omitempty"`
}

// +kubebuilder:skipversion
type CrossAccountResource struct {
	AttachmentARN *string `json:"attachmentARN,omitempty"`

	EndpointID *string `json:"endpointID,omitempty"`
}

// +kubebuilder:skipversion
type CustomRoutingAccelerator struct {
	AcceleratorARN *string `json:"acceleratorARN,omitempty"`

	CreatedTime *metav1.Time `json:"createdTime,omitempty"`

	DNSName *string `json:"dnsName,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	IPAddressType *string `json:"ipAddressType,omitempty"`

	IPSets []*IPSet `json:"ipSets,omitempty"`

	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`

	Name *string `json:"name,omitempty"`
}

// +kubebuilder:skipversion
type CustomRoutingAcceleratorAttributes struct {
	FlowLogsEnabled *bool `json:"flowLogsEnabled,omitempty"`

	FlowLogsS3Bucket *string `json:"flowLogsS3Bucket,omitempty"`

	FlowLogsS3Prefix *string `json:"flowLogsS3Prefix,omitempty"`
}

// +kubebuilder:skipversion
type CustomRoutingDestinationConfiguration struct {
	FromPort *int64 `json:"fromPort,omitempty"`

	ToPort *int64 `json:"toPort,omitempty"`
}

// +kubebuilder:skipversion
type CustomRoutingDestinationDescription struct {
	FromPort *int64 `json:"fromPort,omitempty"`

	ToPort *int64 `json:"toPort,omitempty"`
}

// +kubebuilder:skipversion
type CustomRoutingEndpointConfiguration struct {
	AttachmentARN *string `json:"attachmentARN,omitempty"`

	EndpointID *string `json:"endpointID,omitempty"`
}

// +kubebuilder:skipversion
type CustomRoutingEndpointDescription struct {
	EndpointID *string `json:"endpointID,omitempty"`
}

// +kubebuilder:skipversion
type CustomRoutingEndpointGroup struct {
	EndpointGroupARN *string `json:"endpointGroupARN,omitempty"`

	EndpointGroupRegion *string `json:"endpointGroupRegion,omitempty"`
}

// +kubebuilder:skipversion
type CustomRoutingListener struct {
	ListenerARN *string `json:"listenerARN,omitempty"`

	PortRanges []*PortRange `json:"portRanges,omitempty"`
}

// +kubebuilder:skipversion
type DestinationPortMapping struct {
	AcceleratorARN *string `json:"acceleratorARN,omitempty"`

	EndpointGroupARN *string `json:"endpointGroupARN,omitempty"`

	EndpointGroupRegion *string `json:"endpointGroupRegion,omitempty"`

	EndpointID *string `json:"endpointID,omitempty"`

	IPAddressType *string `json:"ipAddressType,omitempty"`
}

// +kubebuilder:skipversion
type EndpointConfiguration struct {
	AttachmentARN *string `json:"attachmentARN,omitempty"`

	ClientIPPreservationEnabled *bool `json:"clientIPPreservationEnabled,omitempty"`

	EndpointID *string `json:"endpointID,omitempty"`

	Weight *int64 `json:"weight,omitempty"`
}

// +kubebuilder:skipversion
type EndpointDescription struct {
	ClientIPPreservationEnabled *bool `json:"clientIPPreservationEnabled,omitempty"`

	EndpointID *string `json:"endpointID,omitempty"`

	HealthReason *string `json:"healthReason,omitempty"`

	HealthState *string `json:"healthState,omitempty"`

	Weight *int64 `json:"weight,omitempty"`
}

// +kubebuilder:skipversion
type EndpointGroup_SDK struct {
	EndpointDescriptions []*EndpointDescription `json:"endpointDescriptions,omitempty"`

	EndpointGroupARN *string `json:"endpointGroupARN,omitempty"`

	EndpointGroupRegion *string `json:"endpointGroupRegion,omitempty"`

	HealthCheckIntervalSeconds *int64 `json:"healthCheckIntervalSeconds,omitempty"`

	HealthCheckPath *string `json:"healthCheckPath,omitempty"`

	HealthCheckPort *int64 `json:"healthCheckPort,omitempty"`

	HealthCheckProtocol *string `json:"healthCheckProtocol,omitempty"`

	PortOverrides []*PortOverride `json:"portOverrides,omitempty"`

	ThresholdCount *int64 `json:"thresholdCount,omitempty"`

	TrafficDialPercentage *float64 `json:"trafficDialPercentage,omitempty"`
}

// +kubebuilder:skipversion
type EndpointIdentifier struct {
	ClientIPPreservationEnabled *bool `json:"clientIPPreservationEnabled,omitempty"`

	EndpointID *string `json:"endpointID,omitempty"`
}

// +kubebuilder:skipversion
type IPSet struct {
	IPAddressFamily *string `json:"ipAddressFamily,omitempty"`

	IPAddresses []*string `json:"ipAddresses,omitempty"`

	IPFamily *string `json:"ipFamily,omitempty"`
}

// +kubebuilder:skipversion
type Listener_SDK struct {
	ClientAffinity *string `json:"clientAffinity,omitempty"`

	ListenerARN *string `json:"listenerARN,omitempty"`

	PortRanges []*PortRange `json:"portRanges,omitempty"`

	Protocol *string `json:"protocol,omitempty"`
}

// +kubebuilder:skipversion
type PortMapping struct {
	AcceleratorPort *int64 `json:"acceleratorPort,omitempty"`

	EndpointGroupARN *string `json:"endpointGroupARN,omitempty"`

	EndpointID *string `json:"endpointID,omitempty"`
}

// +kubebuilder:skipversion
type PortOverride struct {
	EndpointPort *int64 `json:"endpointPort,omitempty"`

	ListenerPort *int64 `json:"listenerPort,omitempty"`
}

// +kubebuilder:skipversion
type PortRange struct {
	FromPort *int64 `json:"fromPort,omitempty"`

	ToPort *int64 `json:"toPort,omitempty"`
}

// +kubebuilder:skipversion
type Resource struct {
	EndpointID *string `json:"endpointID,omitempty"`

	Region *string `json:"region,omitempty"`
}

// +kubebuilder:skipversion
type SocketAddress struct {
	IPAddress *string `json:"ipAddress,omitempty"`

	Port *int64 `json:"port,omitempty"`
}

// +kubebuilder:skipversion
type Tag struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}
