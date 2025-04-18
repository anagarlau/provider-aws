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

package vpcendpoint

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/ec2"
	svcsdk "github.com/aws/aws-sdk-go/service/ec2"
	svcsdkapi "github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/ec2/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an VPCEndpoint resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create VPCEndpoint in AWS"
	errUpdate        = "cannot update VPCEndpoint in AWS"
	errDescribe      = "failed to describe VPCEndpoint"
	errDelete        = "failed to delete VPCEndpoint"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, cr *svcapitypes.VPCEndpoint) (managed.TypedExternalClient[*svcapitypes.VPCEndpoint], error) {
	sess, err := connectaws.GetConfigV1(ctx, c.kube, cr, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, cr *svcapitypes.VPCEndpoint) (managed.ExternalObservation, error) {
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeVpcEndpointsInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeVpcEndpointsWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	resp = e.filterList(cr, resp)
	if len(resp.VpcEndpoints) == 0 {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateVPCEndpoint(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
	upToDate := true
	diff := ""
	if !meta.WasDeleted(cr) { // There is no need to run isUpToDate if the resource is deleted
		upToDate, diff, err = e.isUpToDate(ctx, cr, resp)
		if err != nil {
			return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
		}
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		Diff:                    diff,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, cr *svcapitypes.VPCEndpoint) (managed.ExternalCreation, error) {
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateVpcEndpointInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateVpcEndpointWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.VpcEndpoint.CreationTimestamp != nil {
		cr.Status.AtProvider.CreationTimestamp = &metav1.Time{*resp.VpcEndpoint.CreationTimestamp}
	} else {
		cr.Status.AtProvider.CreationTimestamp = nil
	}
	if resp.VpcEndpoint.DnsEntries != nil {
		f1 := []*svcapitypes.DNSEntry{}
		for _, f1iter := range resp.VpcEndpoint.DnsEntries {
			f1elem := &svcapitypes.DNSEntry{}
			if f1iter.DnsName != nil {
				f1elem.DNSName = f1iter.DnsName
			}
			if f1iter.HostedZoneId != nil {
				f1elem.HostedZoneID = f1iter.HostedZoneId
			}
			f1 = append(f1, f1elem)
		}
		cr.Status.AtProvider.DNSEntries = f1
	} else {
		cr.Status.AtProvider.DNSEntries = nil
	}
	if resp.VpcEndpoint.DnsOptions != nil {
		f2 := &svcapitypes.DNSOptionsSpecification{}
		if resp.VpcEndpoint.DnsOptions.DnsRecordIpType != nil {
			f2.DNSRecordIPType = resp.VpcEndpoint.DnsOptions.DnsRecordIpType
		}
		if resp.VpcEndpoint.DnsOptions.PrivateDnsOnlyForInboundResolverEndpoint != nil {
			f2.PrivateDNSOnlyForInboundResolverEndpoint = resp.VpcEndpoint.DnsOptions.PrivateDnsOnlyForInboundResolverEndpoint
		}
		cr.Spec.ForProvider.DNSOptions = f2
	} else {
		cr.Spec.ForProvider.DNSOptions = nil
	}
	if resp.VpcEndpoint.Groups != nil {
		f3 := []*svcapitypes.SecurityGroupIdentifier{}
		for _, f3iter := range resp.VpcEndpoint.Groups {
			f3elem := &svcapitypes.SecurityGroupIdentifier{}
			if f3iter.GroupId != nil {
				f3elem.GroupID = f3iter.GroupId
			}
			if f3iter.GroupName != nil {
				f3elem.GroupName = f3iter.GroupName
			}
			f3 = append(f3, f3elem)
		}
		cr.Status.AtProvider.Groups = f3
	} else {
		cr.Status.AtProvider.Groups = nil
	}
	if resp.VpcEndpoint.IpAddressType != nil {
		cr.Spec.ForProvider.IPAddressType = resp.VpcEndpoint.IpAddressType
	} else {
		cr.Spec.ForProvider.IPAddressType = nil
	}
	if resp.VpcEndpoint.LastError != nil {
		f5 := &svcapitypes.LastError{}
		if resp.VpcEndpoint.LastError.Code != nil {
			f5.Code = resp.VpcEndpoint.LastError.Code
		}
		if resp.VpcEndpoint.LastError.Message != nil {
			f5.Message = resp.VpcEndpoint.LastError.Message
		}
		cr.Status.AtProvider.LastError = f5
	} else {
		cr.Status.AtProvider.LastError = nil
	}
	if resp.VpcEndpoint.NetworkInterfaceIds != nil {
		f6 := []*string{}
		for _, f6iter := range resp.VpcEndpoint.NetworkInterfaceIds {
			var f6elem string
			f6elem = *f6iter
			f6 = append(f6, &f6elem)
		}
		cr.Status.AtProvider.NetworkInterfaceIDs = f6
	} else {
		cr.Status.AtProvider.NetworkInterfaceIDs = nil
	}
	if resp.VpcEndpoint.OwnerId != nil {
		cr.Status.AtProvider.OwnerID = resp.VpcEndpoint.OwnerId
	} else {
		cr.Status.AtProvider.OwnerID = nil
	}
	if resp.VpcEndpoint.PolicyDocument != nil {
		cr.Spec.ForProvider.PolicyDocument = resp.VpcEndpoint.PolicyDocument
	} else {
		cr.Spec.ForProvider.PolicyDocument = nil
	}
	if resp.VpcEndpoint.PrivateDnsEnabled != nil {
		cr.Spec.ForProvider.PrivateDNSEnabled = resp.VpcEndpoint.PrivateDnsEnabled
	} else {
		cr.Spec.ForProvider.PrivateDNSEnabled = nil
	}
	if resp.VpcEndpoint.RequesterManaged != nil {
		cr.Status.AtProvider.RequesterManaged = resp.VpcEndpoint.RequesterManaged
	} else {
		cr.Status.AtProvider.RequesterManaged = nil
	}
	if resp.VpcEndpoint.RouteTableIds != nil {
		f11 := []*string{}
		for _, f11iter := range resp.VpcEndpoint.RouteTableIds {
			var f11elem string
			f11elem = *f11iter
			f11 = append(f11, &f11elem)
		}
		cr.Status.AtProvider.RouteTableIDs = f11
	} else {
		cr.Status.AtProvider.RouteTableIDs = nil
	}
	if resp.VpcEndpoint.ServiceName != nil {
		cr.Spec.ForProvider.ServiceName = resp.VpcEndpoint.ServiceName
	} else {
		cr.Spec.ForProvider.ServiceName = nil
	}
	if resp.VpcEndpoint.State != nil {
		cr.Status.AtProvider.State = resp.VpcEndpoint.State
	} else {
		cr.Status.AtProvider.State = nil
	}
	if resp.VpcEndpoint.SubnetIds != nil {
		f14 := []*string{}
		for _, f14iter := range resp.VpcEndpoint.SubnetIds {
			var f14elem string
			f14elem = *f14iter
			f14 = append(f14, &f14elem)
		}
		cr.Status.AtProvider.SubnetIDs = f14
	} else {
		cr.Status.AtProvider.SubnetIDs = nil
	}
	if resp.VpcEndpoint.Tags != nil {
		f15 := []*svcapitypes.Tag{}
		for _, f15iter := range resp.VpcEndpoint.Tags {
			f15elem := &svcapitypes.Tag{}
			if f15iter.Key != nil {
				f15elem.Key = f15iter.Key
			}
			if f15iter.Value != nil {
				f15elem.Value = f15iter.Value
			}
			f15 = append(f15, f15elem)
		}
		cr.Status.AtProvider.Tags = f15
	} else {
		cr.Status.AtProvider.Tags = nil
	}
	if resp.VpcEndpoint.VpcEndpointId != nil {
		cr.Status.AtProvider.VPCEndpointID = resp.VpcEndpoint.VpcEndpointId
	} else {
		cr.Status.AtProvider.VPCEndpointID = nil
	}
	if resp.VpcEndpoint.VpcEndpointType != nil {
		cr.Spec.ForProvider.VPCEndpointType = resp.VpcEndpoint.VpcEndpointType
	} else {
		cr.Spec.ForProvider.VPCEndpointType = nil
	}
	if resp.VpcEndpoint.VpcId != nil {
		cr.Status.AtProvider.VPCID = resp.VpcEndpoint.VpcId
	} else {
		cr.Status.AtProvider.VPCID = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, cr *svcapitypes.VPCEndpoint) (managed.ExternalUpdate, error) {
	input := GenerateModifyVpcEndpointInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.ModifyVpcEndpointWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, cr *svcapitypes.VPCEndpoint) (managed.ExternalDelete, error) {
	cr.Status.SetConditions(xpv1.Deleting())
	return e.delete(ctx, cr)

}

func (e *external) Disconnect(ctx context.Context) error {
	// Unimplemented, required by newer versions of crossplane-runtime
	return nil
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.EC2API, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		filterList:     nopFilterList,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		delete:         nopDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.EC2API
	preObserve     func(context.Context, *svcapitypes.VPCEndpoint, *svcsdk.DescribeVpcEndpointsInput) error
	postObserve    func(context.Context, *svcapitypes.VPCEndpoint, *svcsdk.DescribeVpcEndpointsOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	filterList     func(*svcapitypes.VPCEndpoint, *svcsdk.DescribeVpcEndpointsOutput) *svcsdk.DescribeVpcEndpointsOutput
	lateInitialize func(*svcapitypes.VPCEndpointParameters, *svcsdk.DescribeVpcEndpointsOutput) error
	isUpToDate     func(context.Context, *svcapitypes.VPCEndpoint, *svcsdk.DescribeVpcEndpointsOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.VPCEndpoint, *svcsdk.CreateVpcEndpointInput) error
	postCreate     func(context.Context, *svcapitypes.VPCEndpoint, *svcsdk.CreateVpcEndpointOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	delete         func(context.Context, *svcapitypes.VPCEndpoint) (managed.ExternalDelete, error)
	preUpdate      func(context.Context, *svcapitypes.VPCEndpoint, *svcsdk.ModifyVpcEndpointInput) error
	postUpdate     func(context.Context, *svcapitypes.VPCEndpoint, *svcsdk.ModifyVpcEndpointOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.VPCEndpoint, *svcsdk.DescribeVpcEndpointsInput) error {
	return nil
}
func nopPostObserve(_ context.Context, _ *svcapitypes.VPCEndpoint, _ *svcsdk.DescribeVpcEndpointsOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopFilterList(_ *svcapitypes.VPCEndpoint, list *svcsdk.DescribeVpcEndpointsOutput) *svcsdk.DescribeVpcEndpointsOutput {
	return list
}

func nopLateInitialize(*svcapitypes.VPCEndpointParameters, *svcsdk.DescribeVpcEndpointsOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.VPCEndpoint, *svcsdk.DescribeVpcEndpointsOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.VPCEndpoint, *svcsdk.CreateVpcEndpointInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.VPCEndpoint, _ *svcsdk.CreateVpcEndpointOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopDelete(context.Context, *svcapitypes.VPCEndpoint) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, nil
}
func nopPreUpdate(context.Context, *svcapitypes.VPCEndpoint, *svcsdk.ModifyVpcEndpointInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.VPCEndpoint, _ *svcsdk.ModifyVpcEndpointOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
