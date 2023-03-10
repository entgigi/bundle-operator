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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntandoBundleInstanceV2) DeepCopyInto(out *EntandoBundleInstanceV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntandoBundleInstanceV2.
func (in *EntandoBundleInstanceV2) DeepCopy() *EntandoBundleInstanceV2 {
	if in == nil {
		return nil
	}
	out := new(EntandoBundleInstanceV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EntandoBundleInstanceV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntandoBundleInstanceV2List) DeepCopyInto(out *EntandoBundleInstanceV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EntandoBundleInstanceV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntandoBundleInstanceV2List.
func (in *EntandoBundleInstanceV2List) DeepCopy() *EntandoBundleInstanceV2List {
	if in == nil {
		return nil
	}
	out := new(EntandoBundleInstanceV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EntandoBundleInstanceV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntandoBundleInstanceV2Spec) DeepCopyInto(out *EntandoBundleInstanceV2Spec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntandoBundleInstanceV2Spec.
func (in *EntandoBundleInstanceV2Spec) DeepCopy() *EntandoBundleInstanceV2Spec {
	if in == nil {
		return nil
	}
	out := new(EntandoBundleInstanceV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntandoBundleInstanceV2Status) DeepCopyInto(out *EntandoBundleInstanceV2Status) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntandoBundleInstanceV2Status.
func (in *EntandoBundleInstanceV2Status) DeepCopy() *EntandoBundleInstanceV2Status {
	if in == nil {
		return nil
	}
	out := new(EntandoBundleInstanceV2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntandoBundleTag) DeepCopyInto(out *EntandoBundleTag) {
	*out = *in
	if in.SignatureInfo != nil {
		in, out := &in.SignatureInfo, &out.SignatureInfo
		*out = make([]SignatureInfo, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntandoBundleTag.
func (in *EntandoBundleTag) DeepCopy() *EntandoBundleTag {
	if in == nil {
		return nil
	}
	out := new(EntandoBundleTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntandoBundleV2) DeepCopyInto(out *EntandoBundleV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntandoBundleV2.
func (in *EntandoBundleV2) DeepCopy() *EntandoBundleV2 {
	if in == nil {
		return nil
	}
	out := new(EntandoBundleV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EntandoBundleV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntandoBundleV2List) DeepCopyInto(out *EntandoBundleV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EntandoBundleV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntandoBundleV2List.
func (in *EntandoBundleV2List) DeepCopy() *EntandoBundleV2List {
	if in == nil {
		return nil
	}
	out := new(EntandoBundleV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EntandoBundleV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntandoBundleV2Spec) DeepCopyInto(out *EntandoBundleV2Spec) {
	*out = *in
	if in.TagList != nil {
		in, out := &in.TagList, &out.TagList
		*out = make([]EntandoBundleTag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntandoBundleV2Spec.
func (in *EntandoBundleV2Spec) DeepCopy() *EntandoBundleV2Spec {
	if in == nil {
		return nil
	}
	out := new(EntandoBundleV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntandoBundleV2Status) DeepCopyInto(out *EntandoBundleV2Status) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntandoBundleV2Status.
func (in *EntandoBundleV2Status) DeepCopy() *EntandoBundleV2Status {
	if in == nil {
		return nil
	}
	out := new(EntandoBundleV2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureInfo) DeepCopyInto(out *SignatureInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureInfo.
func (in *SignatureInfo) DeepCopy() *SignatureInfo {
	if in == nil {
		return nil
	}
	out := new(SignatureInfo)
	in.DeepCopyInto(out)
	return out
}
