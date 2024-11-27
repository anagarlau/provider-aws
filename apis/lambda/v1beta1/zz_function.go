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

package v1beta1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// FunctionParameters defines the desired state of Function
type FunctionParameters struct {
	// Region is which region the Function will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The instruction set architecture that the function supports. Enter a string
	// array with one of the valid values (arm64 or x86_64). The default value is
	// x86_64.
	Architectures []*string `json:"architectures,omitempty"`
	// To enable code signing for this function, specify the ARN of a code-signing
	// configuration. A code-signing configuration includes a set of signing profiles,
	// which define the trusted publishers for this function.
	CodeSigningConfigARN *string `json:"codeSigningConfigARN,omitempty"`
	// A dead-letter queue configuration that specifies the queue or topic where
	// Lambda sends asynchronous events when they fail processing. For more information,
	// see Dead-letter queues (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-dlq).
	DeadLetterConfig *DeadLetterConfig `json:"deadLetterConfig,omitempty"`
	// A description of the function.
	Description *string `json:"description,omitempty"`
	// Environment variables that are accessible from function code during execution.
	Environment *Environment `json:"environment,omitempty"`
	// The size of the function's /tmp directory in MB. The default value is 512,
	// but can be any whole number between 512 and 10,240 MB. For more information,
	// see Configuring ephemeral storage (console) (https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-ephemeral-storage).
	EphemeralStorage *EphemeralStorage `json:"ephemeralStorage,omitempty"`
	// Connection settings for an Amazon EFS file system.
	FileSystemConfigs []*FileSystemConfig `json:"fileSystemConfigs,omitempty"`
	// The name of the method within your code that Lambda calls to run your function.
	// Handler is required if the deployment package is a .zip file archive. The
	// format includes the file name. It can also include namespaces and other qualifiers,
	// depending on the runtime. For more information, see Lambda programming model
	// (https://docs.aws.amazon.com/lambda/latest/dg/foundation-progmodel.html).
	Handler *string `json:"handler,omitempty"`
	// Container image configuration values (https://docs.aws.amazon.com/lambda/latest/dg/configuration-images.html#configuration-images-settings)
	// that override the values in the container image Dockerfile.
	ImageConfig *ImageConfig `json:"imageConfig,omitempty"`
	// The ARN of the Key Management Service (KMS) customer managed key that's used
	// to encrypt your function's environment variables (https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-encryption).
	// When Lambda SnapStart (https://docs.aws.amazon.com/lambda/latest/dg/snapstart-security.html)
	// is activated, Lambda also uses this key is to encrypt your function's snapshot.
	// If you deploy your function using a container image, Lambda also uses this
	// key to encrypt your function when it's deployed. Note that this is not the
	// same key that's used to protect your container image in the Amazon Elastic
	// Container Registry (Amazon ECR). If you don't provide a customer managed
	// key, Lambda uses a default service key.
	KMSKeyARN *string `json:"kmsKeyARN,omitempty"`
	// A list of function layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
	// to add to the function's execution environment. Specify each layer by its
	// ARN, including the version.
	Layers []*string `json:"layers,omitempty"`
	// The amount of memory available to the function (https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-memory-console)
	// at runtime. Increasing the function memory also increases its CPU allocation.
	// The default value is 128 MB. The value can be any multiple of 1 MB.
	MemorySize *int64 `json:"memorySize,omitempty"`
	// The type of deployment package. Set to Image for container image and set
	// to Zip for .zip file archive.
	PackageType *string `json:"packageType,omitempty"`
	// Set to true to publish the first version of the function during creation.
	Publish *bool `json:"publish,omitempty"`
	// The identifier of the function's runtime (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html).
	// Runtime is required if the deployment package is a .zip file archive.
	//
	// The following list includes deprecated runtimes. For more information, see
	// Runtime deprecation policy (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtime-support-policy).
	Runtime *string `json:"runtime,omitempty"`
	// The function's SnapStart (https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html)
	// setting.
	SnapStart *SnapStart `json:"snapStart,omitempty"`
	// A list of tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html)
	// to apply to the function.
	Tags map[string]*string `json:"tags,omitempty"`
	// The amount of time (in seconds) that Lambda allows a function to run before
	// stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds.
	// For more information, see Lambda execution environment (https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html).
	Timeout *int64 `json:"timeout,omitempty"`
	// Set Mode to Active to sample and trace a subset of incoming requests with
	// X-Ray (https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html).
	TracingConfig            *TracingConfig `json:"tracingConfig,omitempty"`
	CustomFunctionParameters `json:",inline"`
}

// FunctionSpec defines the desired state of Function
type FunctionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FunctionParameters `json:"forProvider"`
}

// FunctionObservation defines the observed state of Function
type FunctionObservation struct {
	// The SHA256 hash of the function's deployment package.
	CodeSHA256 *string `json:"codeSHA256,omitempty"`
	// The size of the function's deployment package, in bytes.
	CodeSize *int64 `json:"codeSize,omitempty"`
	// The function's Amazon Resource Name (ARN).
	FunctionARN *string `json:"functionARN,omitempty"`
	// The name of the function.
	FunctionName *string `json:"functionName,omitempty"`
	// The function's image configuration values.
	ImageConfigResponse *ImageConfigResponse `json:"imageConfigResponse,omitempty"`
	// The date and time that the function was last updated, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModified *string `json:"lastModified,omitempty"`
	// The status of the last update that was performed on the function. This is
	// first set to Successful after function creation completes.
	LastUpdateStatus *string `json:"lastUpdateStatus,omitempty"`
	// The reason for the last update that was performed on the function.
	LastUpdateStatusReason *string `json:"lastUpdateStatusReason,omitempty"`
	// The reason code for the last update that was performed on the function.
	LastUpdateStatusReasonCode *string `json:"lastUpdateStatusReasonCode,omitempty"`
	// For Lambda@Edge functions, the ARN of the main function.
	MasterARN *string `json:"masterARN,omitempty"`
	// The latest updated revision of the function or alias.
	RevisionID *string `json:"revisionID,omitempty"`
	// The function's execution role.
	Role *string `json:"role,omitempty"`
	// The ARN of the runtime and any errors that occured.
	RuntimeVersionConfig *RuntimeVersionConfig `json:"runtimeVersionConfig,omitempty"`
	// The ARN of the signing job.
	SigningJobARN *string `json:"signingJobARN,omitempty"`
	// The ARN of the signing profile version.
	SigningProfileVersionARN *string `json:"signingProfileVersionARN,omitempty"`
	// The current state of the function. When the state is Inactive, you can reactivate
	// the function by invoking it.
	State *string `json:"state,omitempty"`
	// The reason for the function's current state.
	StateReason *string `json:"stateReason,omitempty"`
	// The reason code for the function's current state. When the code is Creating,
	// you can't invoke or modify the function.
	StateReasonCode *string `json:"stateReasonCode,omitempty"`
	// The version of the Lambda function.
	Version *string `json:"version,omitempty"`
	// The function's networking configuration.
	VPCConfig *VPCConfigResponse `json:"vpcConfig,omitempty"`
}

// FunctionStatus defines the observed state of Function.
type FunctionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Function is the Schema for the Functions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionSpec   `json:"spec"`
	Status            FunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionList contains a list of Functions
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Function `json:"items"`
}

// Repository type metadata.
var (
	FunctionKind             = "Function"
	FunctionGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FunctionKind}.String()
	FunctionKindAPIVersion   = FunctionKind + "." + GroupVersion.String()
	FunctionGroupVersionKind = GroupVersion.WithKind(FunctionKind)
)

func init() {
	SchemeBuilder.Register(&Function{}, &FunctionList{})
}
