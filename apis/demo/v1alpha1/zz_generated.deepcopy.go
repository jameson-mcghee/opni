// +build !ignore_autogenerated

/*
Copyright 2021.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartOptions) DeepCopyInto(out *ChartOptions) {
	*out = *in
	if in.Set != nil {
		in, out := &in.Set, &out.Set
		*out = make(map[string]intstr.IntOrString, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartOptions.
func (in *ChartOptions) DeepCopy() *ChartOptions {
	if in == nil {
		return nil
	}
	out := new(ChartOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentsSpec) DeepCopyInto(out *ComponentsSpec) {
	*out = *in
	out.Infra = in.Infra
	in.Opni.DeepCopyInto(&out.Opni)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentsSpec.
func (in *ComponentsSpec) DeepCopy() *ComponentsSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraStack) DeepCopyInto(out *InfraStack) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraStack.
func (in *InfraStack) DeepCopy() *InfraStack {
	if in == nil {
		return nil
	}
	out := new(InfraStack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpniDemo) DeepCopyInto(out *OpniDemo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpniDemo.
func (in *OpniDemo) DeepCopy() *OpniDemo {
	if in == nil {
		return nil
	}
	out := new(OpniDemo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpniDemo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpniDemoList) DeepCopyInto(out *OpniDemoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpniDemo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpniDemoList.
func (in *OpniDemoList) DeepCopy() *OpniDemoList {
	if in == nil {
		return nil
	}
	out := new(OpniDemoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpniDemoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpniDemoSpec) DeepCopyInto(out *OpniDemoSpec) {
	*out = *in
	in.Components.DeepCopyInto(&out.Components)
	if in.CreateKibanaDashboard != nil {
		in, out := &in.CreateKibanaDashboard, &out.CreateKibanaDashboard
		*out = new(bool)
		**out = **in
	}
	if in.LoggingCRDNamespace != nil {
		in, out := &in.LoggingCRDNamespace, &out.LoggingCRDNamespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpniDemoSpec.
func (in *OpniDemoSpec) DeepCopy() *OpniDemoSpec {
	if in == nil {
		return nil
	}
	out := new(OpniDemoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpniDemoStatus) DeepCopyInto(out *OpniDemoStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpniDemoStatus.
func (in *OpniDemoStatus) DeepCopy() *OpniDemoStatus {
	if in == nil {
		return nil
	}
	out := new(OpniDemoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpniStack) DeepCopyInto(out *OpniStack) {
	*out = *in
	in.RancherLogging.DeepCopyInto(&out.RancherLogging)
	in.Minio.DeepCopyInto(&out.Minio)
	in.Nats.DeepCopyInto(&out.Nats)
	in.Elastic.DeepCopyInto(&out.Elastic)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpniStack.
func (in *OpniStack) DeepCopy() *OpniStack {
	if in == nil {
		return nil
	}
	out := new(OpniStack)
	in.DeepCopyInto(out)
	return out
}