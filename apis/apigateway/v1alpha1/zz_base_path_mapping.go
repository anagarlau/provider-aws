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

// BasePathMappingParameters defines the desired state of BasePathMapping
type BasePathMappingParameters struct {
	// Region is which region the BasePathMapping will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The base path name that callers of the API must provide as part of the URL
	// after the domain name. This value must be unique for all of the mappings
	// across a single API. Specify '(none)' if you do not want callers to specify
	// a base path name after the domain name.
	BasePath *string `json:"basePath,omitempty"`
	// The domain name of the BasePathMapping resource to create.
	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName"`
	// The name of the API's stage that you want to use for this mapping. Specify
	// '(none)' if you want callers to explicitly specify the stage name after any
	// base path name.
	Stage                           *string `json:"stage,omitempty"`
	CustomBasePathMappingParameters `json:",inline"`
}

// BasePathMappingSpec defines the desired state of BasePathMapping
type BasePathMappingSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       BasePathMappingParameters `json:"forProvider"`
}

// BasePathMappingObservation defines the observed state of BasePathMapping
type BasePathMappingObservation struct {
	// The string identifier of the associated RestApi.
	RestAPIID *string `json:"restAPIID,omitempty"`

	CustomBasePathMappingObservation `json:",inline"`
}

// BasePathMappingStatus defines the observed state of BasePathMapping.
type BasePathMappingStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          BasePathMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BasePathMapping is the Schema for the BasePathMappings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BasePathMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BasePathMappingSpec   `json:"spec"`
	Status            BasePathMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BasePathMappingList contains a list of BasePathMappings
type BasePathMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BasePathMapping `json:"items"`
}

// Repository type metadata.
var (
	BasePathMappingKind             = "BasePathMapping"
	BasePathMappingGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BasePathMappingKind}.String()
	BasePathMappingKindAPIVersion   = BasePathMappingKind + "." + GroupVersion.String()
	BasePathMappingGroupVersionKind = GroupVersion.WithKind(BasePathMappingKind)
)

func init() {
	SchemeBuilder.Register(&BasePathMapping{}, &BasePathMappingList{})
}
