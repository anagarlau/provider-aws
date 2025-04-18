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

package volume

import (
	"context"

	svcsdk "github.com/aws/aws-sdk-go/service/ec2"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/connection"
	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	ctrl "sigs.k8s.io/controller-runtime"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/ec2/v1alpha1"
	"github.com/crossplane-contrib/provider-aws/apis/v1alpha1"
	"github.com/crossplane-contrib/provider-aws/pkg/features"
	"github.com/crossplane-contrib/provider-aws/pkg/utils/pointer"
	custommanaged "github.com/crossplane-contrib/provider-aws/pkg/utils/reconciler/managed"
)

// SetupVolume adds a controller that reconciles Volume.
func SetupVolume(mgr ctrl.Manager, o controller.Options) error {
	name := managed.ControllerName(svcapitypes.VolumeGroupKind)
	opts := []option{
		func(e *external) {
			e.preCreate = preCreate
			e.postCreate = postCreate
			e.postObserve = postObserve
			e.filterList = filterList
		},
	}

	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.Features.Enabled(features.EnableAlphaExternalSecretStores) {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), v1alpha1.StoreConfigGroupVersionKind))
	}

	reconcilerOpts := []managed.ReconcilerOption{
		managed.WithInitializers(),
		managed.WithCriticalAnnotationUpdater(custommanaged.NewRetryingCriticalAnnotationUpdater(mgr.GetClient())),
		managed.WithTypedExternalConnector(&connector{kube: mgr.GetClient(), opts: opts}),
		managed.WithPollInterval(o.PollInterval),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithConnectionPublishers(cps...),
	}

	if o.Features.Enabled(features.EnableAlphaManagementPolicies) {
		reconcilerOpts = append(reconcilerOpts, managed.WithManagementPolicies())
	}

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(svcapitypes.VolumeGroupVersionKind),
		reconcilerOpts...)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		WithEventFilter(resource.DesiredStateChanged()).
		For(&svcapitypes.Volume{}).
		Complete(r)
}

func filterList(cr *svcapitypes.Volume, obj *svcsdk.DescribeVolumesOutput) *svcsdk.DescribeVolumesOutput {
	volumeIdentifier := pointer.ToOrNilIfZeroValue(meta.GetExternalName(cr))
	resp := &svcsdk.DescribeVolumesOutput{}
	for _, volume := range obj.Volumes {
		if pointer.StringValue(volume.VolumeId) == pointer.StringValue(volumeIdentifier) {
			resp.Volumes = append(resp.Volumes, volume)
			break
		}
	}
	return resp
}

func preCreate(_ context.Context, cr *svcapitypes.Volume, obj *svcsdk.CreateVolumeInput) error {
	obj.KmsKeyId = cr.Spec.ForProvider.KMSKeyID
	obj.ClientToken = pointer.ToOrNilIfZeroValue(string(cr.UID))
	return nil
}

func postCreate(_ context.Context, cr *svcapitypes.Volume, obj *svcsdk.Volume, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	if err != nil {
		return managed.ExternalCreation{}, err
	}
	meta.SetExternalName(cr, pointer.StringValue(obj.VolumeId))
	return managed.ExternalCreation{}, nil
}

func postObserve(_ context.Context, cr *svcapitypes.Volume, obj *svcsdk.DescribeVolumesOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	if err != nil {
		return managed.ExternalObservation{}, err
	}

	switch pointer.StringValue(obj.Volumes[0].State) {
	case string(svcapitypes.VolumeState_available):
		cr.SetConditions(xpv1.Available())
	case string(svcapitypes.VolumeState_creating):
		cr.SetConditions(xpv1.Creating())
	case string(svcapitypes.VolumeState_error):
		cr.SetConditions(xpv1.Unavailable())
	case string(svcapitypes.VolumeState_deleting):
		cr.SetConditions(xpv1.Deleting())
	}

	obs.ConnectionDetails = managed.ConnectionDetails{
		"volumeID": []byte(pointer.StringValue(obj.Volumes[0].VolumeId)),
	}
	return obs, nil
}
