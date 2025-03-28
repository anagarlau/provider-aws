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

package endpointgroup

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/globalaccelerator"
	svcsdk "github.com/aws/aws-sdk-go/service/globalaccelerator"
	svcsdkapi "github.com/aws/aws-sdk-go/service/globalaccelerator/globalacceleratoriface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/globalaccelerator/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an EndpointGroup resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create EndpointGroup in AWS"
	errUpdate        = "cannot update EndpointGroup in AWS"
	errDescribe      = "failed to describe EndpointGroup"
	errDelete        = "failed to delete EndpointGroup"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, cr *svcapitypes.EndpointGroup) (managed.TypedExternalClient[*svcapitypes.EndpointGroup], error) {
	sess, err := connectaws.GetConfigV1(ctx, c.kube, cr, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, cr *svcapitypes.EndpointGroup) (managed.ExternalObservation, error) {
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeEndpointGroupInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeEndpointGroupWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateEndpointGroup(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
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

func (e *external) Create(ctx context.Context, cr *svcapitypes.EndpointGroup) (managed.ExternalCreation, error) {
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateEndpointGroupInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateEndpointGroupWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.EndpointGroup.EndpointDescriptions != nil {
		f0 := []*svcapitypes.EndpointDescription{}
		for _, f0iter := range resp.EndpointGroup.EndpointDescriptions {
			f0elem := &svcapitypes.EndpointDescription{}
			if f0iter.ClientIPPreservationEnabled != nil {
				f0elem.ClientIPPreservationEnabled = f0iter.ClientIPPreservationEnabled
			}
			if f0iter.EndpointId != nil {
				f0elem.EndpointID = f0iter.EndpointId
			}
			if f0iter.HealthReason != nil {
				f0elem.HealthReason = f0iter.HealthReason
			}
			if f0iter.HealthState != nil {
				f0elem.HealthState = f0iter.HealthState
			}
			if f0iter.Weight != nil {
				f0elem.Weight = f0iter.Weight
			}
			f0 = append(f0, f0elem)
		}
		cr.Status.AtProvider.EndpointDescriptions = f0
	} else {
		cr.Status.AtProvider.EndpointDescriptions = nil
	}
	if resp.EndpointGroup.EndpointGroupArn != nil {
		cr.Status.AtProvider.EndpointGroupARN = resp.EndpointGroup.EndpointGroupArn
	} else {
		cr.Status.AtProvider.EndpointGroupARN = nil
	}
	if resp.EndpointGroup.EndpointGroupRegion != nil {
		cr.Spec.ForProvider.EndpointGroupRegion = resp.EndpointGroup.EndpointGroupRegion
	} else {
		cr.Spec.ForProvider.EndpointGroupRegion = nil
	}
	if resp.EndpointGroup.HealthCheckIntervalSeconds != nil {
		cr.Spec.ForProvider.HealthCheckIntervalSeconds = resp.EndpointGroup.HealthCheckIntervalSeconds
	} else {
		cr.Spec.ForProvider.HealthCheckIntervalSeconds = nil
	}
	if resp.EndpointGroup.HealthCheckPath != nil {
		cr.Spec.ForProvider.HealthCheckPath = resp.EndpointGroup.HealthCheckPath
	} else {
		cr.Spec.ForProvider.HealthCheckPath = nil
	}
	if resp.EndpointGroup.HealthCheckPort != nil {
		cr.Spec.ForProvider.HealthCheckPort = resp.EndpointGroup.HealthCheckPort
	} else {
		cr.Spec.ForProvider.HealthCheckPort = nil
	}
	if resp.EndpointGroup.HealthCheckProtocol != nil {
		cr.Spec.ForProvider.HealthCheckProtocol = resp.EndpointGroup.HealthCheckProtocol
	} else {
		cr.Spec.ForProvider.HealthCheckProtocol = nil
	}
	if resp.EndpointGroup.PortOverrides != nil {
		f7 := []*svcapitypes.PortOverride{}
		for _, f7iter := range resp.EndpointGroup.PortOverrides {
			f7elem := &svcapitypes.PortOverride{}
			if f7iter.EndpointPort != nil {
				f7elem.EndpointPort = f7iter.EndpointPort
			}
			if f7iter.ListenerPort != nil {
				f7elem.ListenerPort = f7iter.ListenerPort
			}
			f7 = append(f7, f7elem)
		}
		cr.Spec.ForProvider.PortOverrides = f7
	} else {
		cr.Spec.ForProvider.PortOverrides = nil
	}
	if resp.EndpointGroup.ThresholdCount != nil {
		cr.Spec.ForProvider.ThresholdCount = resp.EndpointGroup.ThresholdCount
	} else {
		cr.Spec.ForProvider.ThresholdCount = nil
	}
	if resp.EndpointGroup.TrafficDialPercentage != nil {
		cr.Spec.ForProvider.TrafficDialPercentage = resp.EndpointGroup.TrafficDialPercentage
	} else {
		cr.Spec.ForProvider.TrafficDialPercentage = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, cr *svcapitypes.EndpointGroup) (managed.ExternalUpdate, error) {
	input := GenerateUpdateEndpointGroupInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateEndpointGroupWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, cr *svcapitypes.EndpointGroup) (managed.ExternalDelete, error) {
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteEndpointGroupInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return managed.ExternalDelete{}, errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return managed.ExternalDelete{}, nil
	}
	resp, err := e.client.DeleteEndpointGroupWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

func (e *external) Disconnect(ctx context.Context) error {
	// Unimplemented, required by newer versions of crossplane-runtime
	return nil
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.GlobalAcceleratorAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
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
	client         svcsdkapi.GlobalAcceleratorAPI
	preObserve     func(context.Context, *svcapitypes.EndpointGroup, *svcsdk.DescribeEndpointGroupInput) error
	postObserve    func(context.Context, *svcapitypes.EndpointGroup, *svcsdk.DescribeEndpointGroupOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.EndpointGroupParameters, *svcsdk.DescribeEndpointGroupOutput) error
	isUpToDate     func(context.Context, *svcapitypes.EndpointGroup, *svcsdk.DescribeEndpointGroupOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.EndpointGroup, *svcsdk.CreateEndpointGroupInput) error
	postCreate     func(context.Context, *svcapitypes.EndpointGroup, *svcsdk.CreateEndpointGroupOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.EndpointGroup, *svcsdk.DeleteEndpointGroupInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.EndpointGroup, *svcsdk.DeleteEndpointGroupOutput, error) (managed.ExternalDelete, error)
	preUpdate      func(context.Context, *svcapitypes.EndpointGroup, *svcsdk.UpdateEndpointGroupInput) error
	postUpdate     func(context.Context, *svcapitypes.EndpointGroup, *svcsdk.UpdateEndpointGroupOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.EndpointGroup, *svcsdk.DescribeEndpointGroupInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.EndpointGroup, _ *svcsdk.DescribeEndpointGroupOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.EndpointGroupParameters, *svcsdk.DescribeEndpointGroupOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.EndpointGroup, *svcsdk.DescribeEndpointGroupOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.EndpointGroup, *svcsdk.CreateEndpointGroupInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.EndpointGroup, _ *svcsdk.CreateEndpointGroupOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.EndpointGroup, *svcsdk.DeleteEndpointGroupInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.EndpointGroup, _ *svcsdk.DeleteEndpointGroupOutput, err error) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, err
}
func nopPreUpdate(context.Context, *svcapitypes.EndpointGroup, *svcsdk.UpdateEndpointGroupInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.EndpointGroup, _ *svcsdk.UpdateEndpointGroupOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
