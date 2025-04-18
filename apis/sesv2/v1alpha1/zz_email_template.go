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

// EmailTemplateParameters defines the desired state of EmailTemplate
type EmailTemplateParameters struct {
	// Region is which region the EmailTemplate will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The content of the email template, composed of a subject line, an HTML part,
	// and a text-only part.
	// +kubebuilder:validation:Required
	TemplateContent               *EmailTemplateContent `json:"templateContent"`
	CustomEmailTemplateParameters `json:",inline"`
}

// EmailTemplateSpec defines the desired state of EmailTemplate
type EmailTemplateSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EmailTemplateParameters `json:"forProvider"`
}

// EmailTemplateObservation defines the observed state of EmailTemplate
type EmailTemplateObservation struct {
	CustomEmailTemplateObservation `json:",inline"`
}

// EmailTemplateStatus defines the observed state of EmailTemplate.
type EmailTemplateStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EmailTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EmailTemplate is the Schema for the EmailTemplates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EmailTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmailTemplateSpec   `json:"spec"`
	Status            EmailTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmailTemplateList contains a list of EmailTemplates
type EmailTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmailTemplate `json:"items"`
}

// Repository type metadata.
var (
	EmailTemplateKind             = "EmailTemplate"
	EmailTemplateGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EmailTemplateKind}.String()
	EmailTemplateKindAPIVersion   = EmailTemplateKind + "." + GroupVersion.String()
	EmailTemplateGroupVersionKind = GroupVersion.WithKind(EmailTemplateKind)
)

func init() {
	SchemeBuilder.Register(&EmailTemplate{}, &EmailTemplateList{})
}
