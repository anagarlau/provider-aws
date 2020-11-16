/*
Copyright 2020 The Crossplane Authors.

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

package domainname

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigatewayv2/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.
// TODO(muvaf): We can generate one-time boilerplate for these hooks but currently
// ACK doesn't support not generating if file exists.
// GenerateGetDomainNamesInput returns input for read
// operation.
func GenerateGetDomainNamesInput(cr *svcapitypes.DomainName) *svcsdk.GetDomainNamesInput {
	res := preGenerateGetDomainNamesInput(cr, &svcsdk.GetDomainNamesInput{})

	return postGenerateGetDomainNamesInput(cr, res)
}

// GenerateDomainName returns the current state in the form of *svcapitypes.DomainName.
func GenerateDomainName(resp *svcsdk.GetDomainNamesOutput) *svcapitypes.DomainName {
	cr := &svcapitypes.DomainName{}

	found := false
	for _, elem := range resp.Items {
		if elem.ApiMappingSelectionExpression != nil {
			cr.Status.AtProvider.APIMappingSelectionExpression = elem.ApiMappingSelectionExpression
		}
		if elem.DomainName != nil {
			cr.Status.AtProvider.DomainName = elem.DomainName
		}
		if elem.DomainNameConfigurations != nil {
			f2 := []*svcapitypes.DomainNameConfiguration{}
			for _, f2iter := range elem.DomainNameConfigurations {
				f2elem := &svcapitypes.DomainNameConfiguration{}
				if f2iter.ApiGatewayDomainName != nil {
					f2elem.APIGatewayDomainName = f2iter.ApiGatewayDomainName
				}
				if f2iter.CertificateArn != nil {
					f2elem.CertificateARN = f2iter.CertificateArn
				}
				if f2iter.CertificateName != nil {
					f2elem.CertificateName = f2iter.CertificateName
				}
				if f2iter.CertificateUploadDate != nil {
					f2elem.CertificateUploadDate = &metav1.Time{*f2iter.CertificateUploadDate}
				}
				if f2iter.DomainNameStatus != nil {
					f2elem.DomainNameStatus = f2iter.DomainNameStatus
				}
				if f2iter.DomainNameStatusMessage != nil {
					f2elem.DomainNameStatusMessage = f2iter.DomainNameStatusMessage
				}
				if f2iter.EndpointType != nil {
					f2elem.EndpointType = f2iter.EndpointType
				}
				if f2iter.HostedZoneId != nil {
					f2elem.HostedZoneID = f2iter.HostedZoneId
				}
				if f2iter.SecurityPolicy != nil {
					f2elem.SecurityPolicy = f2iter.SecurityPolicy
				}
				f2 = append(f2, f2elem)
			}
			cr.Spec.ForProvider.DomainNameConfigurations = f2
		}
		if elem.MutualTlsAuthentication != nil {
			f3 := &svcapitypes.MutualTLSAuthenticationInput{}
			if elem.MutualTlsAuthentication.TruststoreUri != nil {
				f3.TruststoreURI = elem.MutualTlsAuthentication.TruststoreUri
			}
			if elem.MutualTlsAuthentication.TruststoreVersion != nil {
				f3.TruststoreVersion = elem.MutualTlsAuthentication.TruststoreVersion
			}
			cr.Spec.ForProvider.MutualTLSAuthentication = f3
		}
		if elem.Tags != nil {
			f4 := map[string]*string{}
			for f4key, f4valiter := range elem.Tags {
				var f4val string
				f4val = *f4valiter
				f4[f4key] = &f4val
			}
			cr.Spec.ForProvider.Tags = f4
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateDomainNameInput returns a create input.
func GenerateCreateDomainNameInput(cr *svcapitypes.DomainName) *svcsdk.CreateDomainNameInput {
	res := preGenerateCreateDomainNameInput(cr, &svcsdk.CreateDomainNameInput{})

	if cr.Spec.ForProvider.DomainNameConfigurations != nil {
		f0 := []*svcsdk.DomainNameConfiguration{}
		for _, f0iter := range cr.Spec.ForProvider.DomainNameConfigurations {
			f0elem := &svcsdk.DomainNameConfiguration{}
			if f0iter.APIGatewayDomainName != nil {
				f0elem.SetApiGatewayDomainName(*f0iter.APIGatewayDomainName)
			}
			if f0iter.CertificateARN != nil {
				f0elem.SetCertificateArn(*f0iter.CertificateARN)
			}
			if f0iter.CertificateName != nil {
				f0elem.SetCertificateName(*f0iter.CertificateName)
			}
			if f0iter.CertificateUploadDate != nil {
				f0elem.SetCertificateUploadDate(f0iter.CertificateUploadDate.Time)
			}
			if f0iter.DomainNameStatus != nil {
				f0elem.SetDomainNameStatus(*f0iter.DomainNameStatus)
			}
			if f0iter.DomainNameStatusMessage != nil {
				f0elem.SetDomainNameStatusMessage(*f0iter.DomainNameStatusMessage)
			}
			if f0iter.EndpointType != nil {
				f0elem.SetEndpointType(*f0iter.EndpointType)
			}
			if f0iter.HostedZoneID != nil {
				f0elem.SetHostedZoneId(*f0iter.HostedZoneID)
			}
			if f0iter.SecurityPolicy != nil {
				f0elem.SetSecurityPolicy(*f0iter.SecurityPolicy)
			}
			f0 = append(f0, f0elem)
		}
		res.SetDomainNameConfigurations(f0)
	}
	if cr.Spec.ForProvider.MutualTLSAuthentication != nil {
		f1 := &svcsdk.MutualTlsAuthenticationInput{}
		if cr.Spec.ForProvider.MutualTLSAuthentication.TruststoreURI != nil {
			f1.SetTruststoreUri(*cr.Spec.ForProvider.MutualTLSAuthentication.TruststoreURI)
		}
		if cr.Spec.ForProvider.MutualTLSAuthentication.TruststoreVersion != nil {
			f1.SetTruststoreVersion(*cr.Spec.ForProvider.MutualTLSAuthentication.TruststoreVersion)
		}
		res.SetMutualTlsAuthentication(f1)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f2 := map[string]*string{}
		for f2key, f2valiter := range cr.Spec.ForProvider.Tags {
			var f2val string
			f2val = *f2valiter
			f2[f2key] = &f2val
		}
		res.SetTags(f2)
	}

	return postGenerateCreateDomainNameInput(cr, res)
}

// GenerateDeleteDomainNameInput returns a deletion input.
func GenerateDeleteDomainNameInput(cr *svcapitypes.DomainName) *svcsdk.DeleteDomainNameInput {
	res := preGenerateDeleteDomainNameInput(cr, &svcsdk.DeleteDomainNameInput{})

	return postGenerateDeleteDomainNameInput(cr, res)
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}