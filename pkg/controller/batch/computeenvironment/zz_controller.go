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

package computeenvironment

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/batch"
	svcsdk "github.com/aws/aws-sdk-go/service/batch"
	svcsdkapi "github.com/aws/aws-sdk-go/service/batch/batchiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/batch/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an ComputeEnvironment resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create ComputeEnvironment in AWS"
	errUpdate        = "cannot update ComputeEnvironment in AWS"
	errDescribe      = "failed to describe ComputeEnvironment"
	errDelete        = "failed to delete ComputeEnvironment"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, cr *svcapitypes.ComputeEnvironment) (managed.TypedExternalClient[*svcapitypes.ComputeEnvironment], error) {
	sess, err := connectaws.GetConfigV1(ctx, c.kube, cr, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, cr *svcapitypes.ComputeEnvironment) (managed.ExternalObservation, error) {
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeComputeEnvironmentsInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeComputeEnvironmentsWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	resp = e.filterList(cr, resp)
	if len(resp.ComputeEnvironments) == 0 {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateComputeEnvironment(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
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

func (e *external) Create(ctx context.Context, cr *svcapitypes.ComputeEnvironment) (managed.ExternalCreation, error) {
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateComputeEnvironmentInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateComputeEnvironmentWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.ComputeEnvironmentArn != nil {
		cr.Status.AtProvider.ComputeEnvironmentARN = resp.ComputeEnvironmentArn
	} else {
		cr.Status.AtProvider.ComputeEnvironmentARN = nil
	}
	if resp.ComputeEnvironmentName != nil {
		cr.Status.AtProvider.ComputeEnvironmentName = resp.ComputeEnvironmentName
	} else {
		cr.Status.AtProvider.ComputeEnvironmentName = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, cr *svcapitypes.ComputeEnvironment) (managed.ExternalUpdate, error) {
	input := GenerateUpdateComputeEnvironmentInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateComputeEnvironmentWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, cr *svcapitypes.ComputeEnvironment) (managed.ExternalDelete, error) {
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteComputeEnvironmentInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return managed.ExternalDelete{}, errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return managed.ExternalDelete{}, nil
	}
	resp, err := e.client.DeleteComputeEnvironmentWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

func (e *external) Disconnect(ctx context.Context) error {
	// Unimplemented, required by newer versions of crossplane-runtime
	return nil
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.BatchAPI, opts []option) *external {
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
	client         svcsdkapi.BatchAPI
	preObserve     func(context.Context, *svcapitypes.ComputeEnvironment, *svcsdk.DescribeComputeEnvironmentsInput) error
	postObserve    func(context.Context, *svcapitypes.ComputeEnvironment, *svcsdk.DescribeComputeEnvironmentsOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	filterList     func(*svcapitypes.ComputeEnvironment, *svcsdk.DescribeComputeEnvironmentsOutput) *svcsdk.DescribeComputeEnvironmentsOutput
	lateInitialize func(*svcapitypes.ComputeEnvironmentParameters, *svcsdk.DescribeComputeEnvironmentsOutput) error
	isUpToDate     func(context.Context, *svcapitypes.ComputeEnvironment, *svcsdk.DescribeComputeEnvironmentsOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.ComputeEnvironment, *svcsdk.CreateComputeEnvironmentInput) error
	postCreate     func(context.Context, *svcapitypes.ComputeEnvironment, *svcsdk.CreateComputeEnvironmentOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.ComputeEnvironment, *svcsdk.DeleteComputeEnvironmentInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.ComputeEnvironment, *svcsdk.DeleteComputeEnvironmentOutput, error) (managed.ExternalDelete, error)
	preUpdate      func(context.Context, *svcapitypes.ComputeEnvironment, *svcsdk.UpdateComputeEnvironmentInput) error
	postUpdate     func(context.Context, *svcapitypes.ComputeEnvironment, *svcsdk.UpdateComputeEnvironmentOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.ComputeEnvironment, *svcsdk.DescribeComputeEnvironmentsInput) error {
	return nil
}
func nopPostObserve(_ context.Context, _ *svcapitypes.ComputeEnvironment, _ *svcsdk.DescribeComputeEnvironmentsOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopFilterList(_ *svcapitypes.ComputeEnvironment, list *svcsdk.DescribeComputeEnvironmentsOutput) *svcsdk.DescribeComputeEnvironmentsOutput {
	return list
}

func nopLateInitialize(*svcapitypes.ComputeEnvironmentParameters, *svcsdk.DescribeComputeEnvironmentsOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.ComputeEnvironment, *svcsdk.DescribeComputeEnvironmentsOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.ComputeEnvironment, *svcsdk.CreateComputeEnvironmentInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.ComputeEnvironment, _ *svcsdk.CreateComputeEnvironmentOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.ComputeEnvironment, *svcsdk.DeleteComputeEnvironmentInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.ComputeEnvironment, _ *svcsdk.DeleteComputeEnvironmentOutput, err error) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, err
}
func nopPreUpdate(context.Context, *svcapitypes.ComputeEnvironment, *svcsdk.UpdateComputeEnvironmentInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.ComputeEnvironment, _ *svcsdk.UpdateComputeEnvironmentOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
