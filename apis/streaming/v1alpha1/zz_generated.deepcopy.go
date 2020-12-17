// +build !ignore_autogenerated

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SRSCluster) DeepCopyInto(out *SRSCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SRSCluster.
func (in *SRSCluster) DeepCopy() *SRSCluster {
	if in == nil {
		return nil
	}
	out := new(SRSCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SRSCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SRSClusterList) DeepCopyInto(out *SRSClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SRSCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SRSClusterList.
func (in *SRSClusterList) DeepCopy() *SRSClusterList {
	if in == nil {
		return nil
	}
	out := new(SRSClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SRSClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SRSClusterNginx) DeepCopyInto(out *SRSClusterNginx) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SRSClusterNginx.
func (in *SRSClusterNginx) DeepCopy() *SRSClusterNginx {
	if in == nil {
		return nil
	}
	out := new(SRSClusterNginx)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SRSClusterNginx) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SRSClusterNginxList) DeepCopyInto(out *SRSClusterNginxList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SRSClusterNginx, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SRSClusterNginxList.
func (in *SRSClusterNginxList) DeepCopy() *SRSClusterNginxList {
	if in == nil {
		return nil
	}
	out := new(SRSClusterNginxList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SRSClusterNginxList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SRSClusterNginxSpec) DeepCopyInto(out *SRSClusterNginxSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SRSClusterNginxSpec.
func (in *SRSClusterNginxSpec) DeepCopy() *SRSClusterNginxSpec {
	if in == nil {
		return nil
	}
	out := new(SRSClusterNginxSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SRSClusterNginxStatus) DeepCopyInto(out *SRSClusterNginxStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SRSClusterNginxStatus.
func (in *SRSClusterNginxStatus) DeepCopy() *SRSClusterNginxStatus {
	if in == nil {
		return nil
	}
	out := new(SRSClusterNginxStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SRSClusterSpec) DeepCopyInto(out *SRSClusterSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(SRSConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SRSClusterSpec.
func (in *SRSClusterSpec) DeepCopy() *SRSClusterSpec {
	if in == nil {
		return nil
	}
	out := new(SRSClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SRSClusterStatus) DeepCopyInto(out *SRSClusterStatus) {
	*out = *in
	if in.PodNames != nil {
		in, out := &in.PodNames, &out.PodNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SRSClusterStatus.
func (in *SRSClusterStatus) DeepCopy() *SRSClusterStatus {
	if in == nil {
		return nil
	}
	out := new(SRSClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SRSConfigServerSpec) DeepCopyInto(out *SRSConfigServerSpec) {
	*out = *in
	if in.Listen != nil {
		in, out := &in.Listen, &out.Listen
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SRSConfigServerSpec.
func (in *SRSConfigServerSpec) DeepCopy() *SRSConfigServerSpec {
	if in == nil {
		return nil
	}
	out := new(SRSConfigServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SRSConfigSpec) DeepCopyInto(out *SRSConfigSpec) {
	*out = *in
	if in.MaxConnextions != nil {
		in, out := &in.MaxConnextions, &out.MaxConnextions
		*out = new(int32)
		**out = **in
	}
	if in.Listen != nil {
		in, out := &in.Listen, &out.Listen
		*out = new(int32)
		**out = **in
	}
	if in.API != nil {
		in, out := &in.API, &out.API
		*out = new(SRSConfigServerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Server != nil {
		in, out := &in.Server, &out.Server
		*out = new(SRSConfigServerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.VHosts != nil {
		in, out := &in.VHosts, &out.VHosts
		*out = make([]VHostSRSConfigSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Origins != nil {
		in, out := &in.Origins, &out.Origins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SRSConfigSpec.
func (in *SRSConfigSpec) DeepCopy() *SRSConfigSpec {
	if in == nil {
		return nil
	}
	out := new(SRSConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VHostSRSConfigSpec) DeepCopyInto(out *VHostSRSConfigSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.MixCorrect != nil {
		in, out := &in.MixCorrect, &out.MixCorrect
		*out = new(bool)
		**out = **in
	}
	if in.HLS != nil {
		in, out := &in.HLS, &out.HLS
		*out = new(VHostSRSHLSConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VHostSRSConfigSpec.
func (in *VHostSRSConfigSpec) DeepCopy() *VHostSRSConfigSpec {
	if in == nil {
		return nil
	}
	out := new(VHostSRSConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VHostSRSHLSConfigSpec) DeepCopyInto(out *VHostSRSHLSConfigSpec) {
	*out = *in
	if in.Fragment != nil {
		in, out := &in.Fragment, &out.Fragment
		*out = new(int32)
		**out = **in
	}
	if in.Window != nil {
		in, out := &in.Window, &out.Window
		*out = new(int32)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.M3U8File != nil {
		in, out := &in.M3U8File, &out.M3U8File
		*out = new(string)
		**out = **in
	}
	if in.TSFile != nil {
		in, out := &in.TSFile, &out.TSFile
		*out = new(string)
		**out = **in
	}
	if in.CleanUp != nil {
		in, out := &in.CleanUp, &out.CleanUp
		*out = new(bool)
		**out = **in
	}
	if in.NBNotify != nil {
		in, out := &in.NBNotify, &out.NBNotify
		*out = new(int32)
		**out = **in
	}
	if in.WaitFrame != nil {
		in, out := &in.WaitFrame, &out.WaitFrame
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VHostSRSHLSConfigSpec.
func (in *VHostSRSHLSConfigSpec) DeepCopy() *VHostSRSHLSConfigSpec {
	if in == nil {
		return nil
	}
	out := new(VHostSRSHLSConfigSpec)
	in.DeepCopyInto(out)
	return out
}
