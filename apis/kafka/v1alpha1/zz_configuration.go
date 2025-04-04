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
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ConfigurationParameters defines the desired state of Configuration
type ConfigurationParameters struct {
	// Region is which region the Configuration will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The description of the configuration.
	Description *string `json:"description,omitempty"`
	// The versions of Apache Kafka with which you can use this MSK configuration.
	KafkaVersions                 []*string `json:"kafkaVersions,omitempty"`
	CustomConfigurationParameters `json:",inline"`
}

// ConfigurationSpec defines the desired state of Configuration
type ConfigurationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ConfigurationParameters `json:"forProvider"`
}

// ConfigurationObservation defines the observed state of Configuration
type ConfigurationObservation struct {
	// The Amazon Resource Name (ARN) of the configuration.
	ARN *string `json:"arn,omitempty"`
	// The time when the configuration was created.
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
	// Latest revision of the configuration.
	LatestRevision *ConfigurationRevision `json:"latestRevision,omitempty"`
	// The name of the configuration. Configuration names are strings that match
	// the regex "^[0-9A-Za-z-]+$".
	Name *string `json:"name,omitempty"`
	// The state of the configuration. The possible states are ACTIVE, DELETING
	// and DELETE_FAILED.
	State *string `json:"state,omitempty"`

	CustomConfigurationObservation `json:",inline"`
}

// ConfigurationStatus defines the observed state of Configuration.
type ConfigurationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Configuration is the Schema for the Configurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Configuration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigurationSpec   `json:"spec"`
	Status            ConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationList contains a list of Configurations
type ConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Configuration `json:"items"`
}

// Repository type metadata.
var (
	ConfigurationKind             = "Configuration"
	ConfigurationGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConfigurationKind}.String()
	ConfigurationKindAPIVersion   = ConfigurationKind + "." + GroupVersion.String()
	ConfigurationGroupVersionKind = GroupVersion.WithKind(ConfigurationKind)
)

func init() {
	SchemeBuilder.Register(&Configuration{}, &ConfigurationList{})
}
