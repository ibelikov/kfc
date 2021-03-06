// +build !ignore_autogenerated

/*
Copyright The Kubeform Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	json "encoding/json"

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureAppService) DeepCopyInto(out *AzureAppService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureAppService.
func (in *AzureAppService) DeepCopy() *AzureAppService {
	if in == nil {
		return nil
	}
	out := new(AzureAppService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureAppService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureAppServiceList) DeepCopyInto(out *AzureAppServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureAppService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureAppServiceList.
func (in *AzureAppServiceList) DeepCopy() *AzureAppServiceList {
	if in == nil {
		return nil
	}
	out := new(AzureAppServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureAppServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureAppServiceOutput) DeepCopyInto(out *AzureAppServiceOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureAppServiceOutput.
func (in *AzureAppServiceOutput) DeepCopy() *AzureAppServiceOutput {
	if in == nil {
		return nil
	}
	out := new(AzureAppServiceOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureAppServiceSpec) DeepCopyInto(out *AzureAppServiceSpec) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	out.ProviderRef = in.ProviderRef
	if in.AlwaysOn != nil {
		in, out := &in.AlwaysOn, &out.AlwaysOn
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.AppServicePlanName != nil {
		in, out := &in.AppServicePlanName, &out.AppServicePlanName
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.AppSettings != nil {
		in, out := &in.AppSettings, &out.AppSettings
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.HttpsOnly != nil {
		in, out := &in.HttpsOnly, &out.HttpsOnly
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.RgLocation != nil {
		in, out := &in.RgLocation, &out.RgLocation
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.RgName != nil {
		in, out := &in.RgName, &out.RgName
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.Use32BitWorkerProcess != nil {
		in, out := &in.Use32BitWorkerProcess, &out.Use32BitWorkerProcess
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureAppServiceSpec.
func (in *AzureAppServiceSpec) DeepCopy() *AzureAppServiceSpec {
	if in == nil {
		return nil
	}
	out := new(AzureAppServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureAppServiceStatus) DeepCopyInto(out *AzureAppServiceStatus) {
	*out = *in
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = new(AzureAppServiceOutput)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureAppServiceStatus.
func (in *AzureAppServiceStatus) DeepCopy() *AzureAppServiceStatus {
	if in == nil {
		return nil
	}
	out := new(AzureAppServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleServiceAccount) DeepCopyInto(out *GoogleServiceAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleServiceAccount.
func (in *GoogleServiceAccount) DeepCopy() *GoogleServiceAccount {
	if in == nil {
		return nil
	}
	out := new(GoogleServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GoogleServiceAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleServiceAccountList) DeepCopyInto(out *GoogleServiceAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GoogleServiceAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleServiceAccountList.
func (in *GoogleServiceAccountList) DeepCopy() *GoogleServiceAccountList {
	if in == nil {
		return nil
	}
	out := new(GoogleServiceAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GoogleServiceAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleServiceAccountOutput) DeepCopyInto(out *GoogleServiceAccountOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleServiceAccountOutput.
func (in *GoogleServiceAccountOutput) DeepCopy() *GoogleServiceAccountOutput {
	if in == nil {
		return nil
	}
	out := new(GoogleServiceAccountOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleServiceAccountSpec) DeepCopyInto(out *GoogleServiceAccountSpec) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	out.ProviderRef = in.ProviderRef
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ProjectRoles != nil {
		in, out := &in.ProjectRoles, &out.ProjectRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleServiceAccountSpec.
func (in *GoogleServiceAccountSpec) DeepCopy() *GoogleServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := new(GoogleServiceAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleServiceAccountStatus) DeepCopyInto(out *GoogleServiceAccountStatus) {
	*out = *in
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = new(GoogleServiceAccountOutput)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleServiceAccountStatus.
func (in *GoogleServiceAccountStatus) DeepCopy() *GoogleServiceAccountStatus {
	if in == nil {
		return nil
	}
	out := new(GoogleServiceAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDS) DeepCopyInto(out *RDS) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDS.
func (in *RDS) DeepCopy() *RDS {
	if in == nil {
		return nil
	}
	out := new(RDS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDS) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSList) DeepCopyInto(out *RDSList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RDS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSList.
func (in *RDSList) DeepCopy() *RDSList {
	if in == nil {
		return nil
	}
	out := new(RDSList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDSList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSOutput) DeepCopyInto(out *RDSOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSOutput.
func (in *RDSOutput) DeepCopy() *RDSOutput {
	if in == nil {
		return nil
	}
	out := new(RDSOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSSpec) DeepCopyInto(out *RDSSpec) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	out.ProviderRef = in.ProviderRef
	if in.EnabledCloudwatchLogsExports != nil {
		in, out := &in.EnabledCloudwatchLogsExports, &out.EnabledCloudwatchLogsExports
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OptionGroupTimeouts != nil {
		in, out := &in.OptionGroupTimeouts, &out.OptionGroupTimeouts
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDS != nil {
		in, out := &in.SubnetIDS, &out.SubnetIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VpcSecurityGroupIDS != nil {
		in, out := &in.VpcSecurityGroupIDS, &out.VpcSecurityGroupIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSSpec.
func (in *RDSSpec) DeepCopy() *RDSSpec {
	if in == nil {
		return nil
	}
	out := new(RDSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSStatus) DeepCopyInto(out *RDSStatus) {
	*out = *in
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = new(RDSOutput)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSStatus.
func (in *RDSStatus) DeepCopy() *RDSStatus {
	if in == nil {
		return nil
	}
	out := new(RDSStatus)
	in.DeepCopyInto(out)
	return out
}
