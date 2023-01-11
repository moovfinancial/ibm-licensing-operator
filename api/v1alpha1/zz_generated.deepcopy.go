//go:build !ignore_autogenerated
// +build !ignore_autogenerated

//
// Copyright 2023 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
/*


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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/IBM/ibm-licensing-operator/api/v1alpha1/features"
	v1 "github.com/openshift/api/route/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Container) DeepCopyInto(out *Container) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Container.
func (in *Container) DeepCopy() *Container {
	if in == nil {
		return nil
	}
	out := new(Container)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Features) DeepCopyInto(out *Features) {
	*out = *in
	if in.HyperThreading != nil {
		in, out := &in.HyperThreading, &out.HyperThreading
		*out = new(features.HyperThreading)
		**out = **in
	}
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(features.Auth)
		**out = **in
	}
	if in.PrometheusQuerySource != nil {
		in, out := &in.PrometheusQuerySource, &out.PrometheusQuerySource
		*out = new(features.PrometheusQuerySource)
		(*in).DeepCopyInto(*out)
	}
	if in.Alerting != nil {
		in, out := &in.Alerting, &out.Alerting
		*out = new(features.Alerting)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceScopeEnabled != nil {
		in, out := &in.NamespaceScopeEnabled, &out.NamespaceScopeEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Features.
func (in *Features) DeepCopy() *Features {
	if in == nil {
		return nil
	}
	out := new(Features)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicenseServiceBaseSpec) DeepCopyInto(out *IBMLicenseServiceBaseSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RouteOptions != nil {
		in, out := &in.RouteOptions, &out.RouteOptions
		*out = new(IBMLicenseServiceRouteOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicenseServiceBaseSpec.
func (in *IBMLicenseServiceBaseSpec) DeepCopy() *IBMLicenseServiceBaseSpec {
	if in == nil {
		return nil
	}
	out := new(IBMLicenseServiceBaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicenseServiceReporter) DeepCopyInto(out *IBMLicenseServiceReporter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicenseServiceReporter.
func (in *IBMLicenseServiceReporter) DeepCopy() *IBMLicenseServiceReporter {
	if in == nil {
		return nil
	}
	out := new(IBMLicenseServiceReporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IBMLicenseServiceReporter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicenseServiceReporterList) DeepCopyInto(out *IBMLicenseServiceReporterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IBMLicenseServiceReporter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicenseServiceReporterList.
func (in *IBMLicenseServiceReporterList) DeepCopy() *IBMLicenseServiceReporterList {
	if in == nil {
		return nil
	}
	out := new(IBMLicenseServiceReporterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IBMLicenseServiceReporterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicenseServiceReporterSpec) DeepCopyInto(out *IBMLicenseServiceReporterSpec) {
	*out = *in
	if in.EnvVariable != nil {
		in, out := &in.EnvVariable, &out.EnvVariable
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.ReceiverContainer.DeepCopyInto(&out.ReceiverContainer)
	in.ReporterUIContainer.DeepCopyInto(&out.ReporterUIContainer)
	in.DatabaseContainer.DeepCopyInto(&out.DatabaseContainer)
	in.IBMLicenseServiceBaseSpec.DeepCopyInto(&out.IBMLicenseServiceBaseSpec)
	out.Capacity = in.Capacity.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicenseServiceReporterSpec.
func (in *IBMLicenseServiceReporterSpec) DeepCopy() *IBMLicenseServiceReporterSpec {
	if in == nil {
		return nil
	}
	out := new(IBMLicenseServiceReporterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicenseServiceReporterStatus) DeepCopyInto(out *IBMLicenseServiceReporterStatus) {
	*out = *in
	if in.LicensingReporterPods != nil {
		in, out := &in.LicensingReporterPods, &out.LicensingReporterPods
		*out = make([]corev1.PodStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicenseServiceReporterStatus.
func (in *IBMLicenseServiceReporterStatus) DeepCopy() *IBMLicenseServiceReporterStatus {
	if in == nil {
		return nil
	}
	out := new(IBMLicenseServiceReporterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicenseServiceRouteOptions) DeepCopyInto(out *IBMLicenseServiceRouteOptions) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(v1.TLSConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicenseServiceRouteOptions.
func (in *IBMLicenseServiceRouteOptions) DeepCopy() *IBMLicenseServiceRouteOptions {
	if in == nil {
		return nil
	}
	out := new(IBMLicenseServiceRouteOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicensing) DeepCopyInto(out *IBMLicensing) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicensing.
func (in *IBMLicensing) DeepCopy() *IBMLicensing {
	if in == nil {
		return nil
	}
	out := new(IBMLicensing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IBMLicensing) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicensingFeaturesStatus) DeepCopyInto(out *IBMLicensingFeaturesStatus) {
	*out = *in
	if in.RHMPEnabled != nil {
		in, out := &in.RHMPEnabled, &out.RHMPEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicensingFeaturesStatus.
func (in *IBMLicensingFeaturesStatus) DeepCopy() *IBMLicensingFeaturesStatus {
	if in == nil {
		return nil
	}
	out := new(IBMLicensingFeaturesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicensingIngressOptions) DeepCopyInto(out *IBMLicensingIngressOptions) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = make([]networkingv1.IngressTLS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicensingIngressOptions.
func (in *IBMLicensingIngressOptions) DeepCopy() *IBMLicensingIngressOptions {
	if in == nil {
		return nil
	}
	out := new(IBMLicensingIngressOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicensingList) DeepCopyInto(out *IBMLicensingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IBMLicensing, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicensingList.
func (in *IBMLicensingList) DeepCopy() *IBMLicensingList {
	if in == nil {
		return nil
	}
	out := new(IBMLicensingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IBMLicensingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicensingMetadata) DeepCopyInto(out *IBMLicensingMetadata) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicensingMetadata.
func (in *IBMLicensingMetadata) DeepCopy() *IBMLicensingMetadata {
	if in == nil {
		return nil
	}
	out := new(IBMLicensingMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IBMLicensingMetadata) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicensingMetadataCondition) DeepCopyInto(out *IBMLicensingMetadataCondition) {
	*out = *in
	if in.Annotation != nil {
		in, out := &in.Annotation, &out.Annotation
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicensingMetadataCondition.
func (in *IBMLicensingMetadataCondition) DeepCopy() *IBMLicensingMetadataCondition {
	if in == nil {
		return nil
	}
	out := new(IBMLicensingMetadataCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicensingMetadataList) DeepCopyInto(out *IBMLicensingMetadataList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IBMLicensingMetadata, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicensingMetadataList.
func (in *IBMLicensingMetadataList) DeepCopy() *IBMLicensingMetadataList {
	if in == nil {
		return nil
	}
	out := new(IBMLicensingMetadataList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IBMLicensingMetadataList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicensingMetadataSpec) DeepCopyInto(out *IBMLicensingMetadataSpec) {
	*out = *in
	in.Condition.DeepCopyInto(&out.Condition)
	if in.Extend != nil {
		in, out := &in.Extend, &out.Extend
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicensingMetadataSpec.
func (in *IBMLicensingMetadataSpec) DeepCopy() *IBMLicensingMetadataSpec {
	if in == nil {
		return nil
	}
	out := new(IBMLicensingMetadataSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicensingMetadataStatus) DeepCopyInto(out *IBMLicensingMetadataStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicensingMetadataStatus.
func (in *IBMLicensingMetadataStatus) DeepCopy() *IBMLicensingMetadataStatus {
	if in == nil {
		return nil
	}
	out := new(IBMLicensingMetadataStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicensingRouteOptions) DeepCopyInto(out *IBMLicensingRouteOptions) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(v1.TLSConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicensingRouteOptions.
func (in *IBMLicensingRouteOptions) DeepCopy() *IBMLicensingRouteOptions {
	if in == nil {
		return nil
	}
	out := new(IBMLicensingRouteOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicensingSecurityContext) DeepCopyInto(out *IBMLicensingSecurityContext) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicensingSecurityContext.
func (in *IBMLicensingSecurityContext) DeepCopy() *IBMLicensingSecurityContext {
	if in == nil {
		return nil
	}
	out := new(IBMLicensingSecurityContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicensingSenderSpec) DeepCopyInto(out *IBMLicensingSenderSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicensingSenderSpec.
func (in *IBMLicensingSenderSpec) DeepCopy() *IBMLicensingSenderSpec {
	if in == nil {
		return nil
	}
	out := new(IBMLicensingSenderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicensingSpec) DeepCopyInto(out *IBMLicensingSpec) {
	*out = *in
	if in.EnvVariable != nil {
		in, out := &in.EnvVariable, &out.EnvVariable
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Container.DeepCopyInto(&out.Container)
	in.IBMLicenseServiceBaseSpec.DeepCopyInto(&out.IBMLicenseServiceBaseSpec)
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(IBMLicensingSecurityContext)
		**out = **in
	}
	if in.RouteEnabled != nil {
		in, out := &in.RouteEnabled, &out.RouteEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RHMPEnabled != nil {
		in, out := &in.RHMPEnabled, &out.RHMPEnabled
		*out = new(bool)
		**out = **in
	}
	in.UsageContainer.DeepCopyInto(&out.UsageContainer)
	if in.ChargebackEnabled != nil {
		in, out := &in.ChargebackEnabled, &out.ChargebackEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ChargebackRetentionPeriod != nil {
		in, out := &in.ChargebackRetentionPeriod, &out.ChargebackRetentionPeriod
		*out = new(int)
		**out = **in
	}
	if in.IngressEnabled != nil {
		in, out := &in.IngressEnabled, &out.IngressEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IngressOptions != nil {
		in, out := &in.IngressOptions, &out.IngressOptions
		*out = new(IBMLicensingIngressOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Sender != nil {
		in, out := &in.Sender, &out.Sender
		*out = new(IBMLicensingSenderSpec)
		**out = **in
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(Features)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicensingSpec.
func (in *IBMLicensingSpec) DeepCopy() *IBMLicensingSpec {
	if in == nil {
		return nil
	}
	out := new(IBMLicensingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMLicensingStatus) DeepCopyInto(out *IBMLicensingStatus) {
	*out = *in
	if in.LicensingPods != nil {
		in, out := &in.LicensingPods, &out.LicensingPods
		*out = make([]corev1.PodStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Features.DeepCopyInto(&out.Features)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMLicensingStatus.
func (in *IBMLicensingStatus) DeepCopy() *IBMLicensingStatus {
	if in == nil {
		return nil
	}
	out := new(IBMLicensingStatus)
	in.DeepCopyInto(out)
	return out
}
