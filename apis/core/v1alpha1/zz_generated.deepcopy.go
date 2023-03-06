//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagConfiguration) DeepCopyInto(out *FeatureFlagConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagConfiguration.
func (in *FeatureFlagConfiguration) DeepCopy() *FeatureFlagConfiguration {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureFlagConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagConfigurationList) DeepCopyInto(out *FeatureFlagConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FeatureFlagConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagConfigurationList.
func (in *FeatureFlagConfigurationList) DeepCopy() *FeatureFlagConfigurationList {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureFlagConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagConfigurationSpec) DeepCopyInto(out *FeatureFlagConfigurationSpec) {
	*out = *in
	if in.ServiceProvider != nil {
		in, out := &in.ServiceProvider, &out.ServiceProvider
		*out = new(FeatureFlagServiceProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.SyncProvider != nil {
		in, out := &in.SyncProvider, &out.SyncProvider
		*out = new(FeatureFlagSyncProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.FlagDSpec != nil {
		in, out := &in.FlagDSpec, &out.FlagDSpec
		*out = new(FlagDSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagConfigurationSpec.
func (in *FeatureFlagConfigurationSpec) DeepCopy() *FeatureFlagConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagConfigurationStatus) DeepCopyInto(out *FeatureFlagConfigurationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagConfigurationStatus.
func (in *FeatureFlagConfigurationStatus) DeepCopy() *FeatureFlagConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagServiceProvider) DeepCopyInto(out *FeatureFlagServiceProvider) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagServiceProvider.
func (in *FeatureFlagServiceProvider) DeepCopy() *FeatureFlagServiceProvider {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagServiceProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagSyncProvider) DeepCopyInto(out *FeatureFlagSyncProvider) {
	*out = *in
	if in.HttpSyncConfiguration != nil {
		in, out := &in.HttpSyncConfiguration, &out.HttpSyncConfiguration
		*out = new(HttpSyncConfiguration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagSyncProvider.
func (in *FeatureFlagSyncProvider) DeepCopy() *FeatureFlagSyncProvider {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagSyncProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlagDSpec) DeepCopyInto(out *FlagDSpec) {
	*out = *in
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlagDSpec.
func (in *FlagDSpec) DeepCopy() *FlagDSpec {
	if in == nil {
		return nil
	}
	out := new(FlagDSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlagSourceConfiguration) DeepCopyInto(out *FlagSourceConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlagSourceConfiguration.
func (in *FlagSourceConfiguration) DeepCopy() *FlagSourceConfiguration {
	if in == nil {
		return nil
	}
	out := new(FlagSourceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlagSourceConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlagSourceConfigurationList) DeepCopyInto(out *FlagSourceConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlagSourceConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlagSourceConfigurationList.
func (in *FlagSourceConfigurationList) DeepCopy() *FlagSourceConfigurationList {
	if in == nil {
		return nil
	}
	out := new(FlagSourceConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlagSourceConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlagSourceConfigurationSpec) DeepCopyInto(out *FlagSourceConfigurationSpec) {
	*out = *in
	if in.SyncProviderArgs != nil {
		in, out := &in.SyncProviderArgs, &out.SyncProviderArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlagSourceConfigurationSpec.
func (in *FlagSourceConfigurationSpec) DeepCopy() *FlagSourceConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(FlagSourceConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlagSourceConfigurationStatus) DeepCopyInto(out *FlagSourceConfigurationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlagSourceConfigurationStatus.
func (in *FlagSourceConfigurationStatus) DeepCopy() *FlagSourceConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(FlagSourceConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpSyncConfiguration) DeepCopyInto(out *HttpSyncConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpSyncConfiguration.
func (in *HttpSyncConfiguration) DeepCopy() *HttpSyncConfiguration {
	if in == nil {
		return nil
	}
	out := new(HttpSyncConfiguration)
	in.DeepCopyInto(out)
	return out
}
