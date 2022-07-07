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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressSelector) DeepCopyInto(out *EgressSelector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressSelector.
func (in *EgressSelector) DeepCopy() *EgressSelector {
	if in == nil {
		return nil
	}
	out := new(EgressSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderMatch) DeepCopyInto(out *HeaderMatch) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderMatch.
func (in *HeaderMatch) DeepCopy() *HeaderMatch {
	if in == nil {
		return nil
	}
	out := new(HeaderMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpGlobalRateLimit) DeepCopyInto(out *HttpGlobalRateLimit) {
	*out = *in
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]*RateLimitMatch, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(RateLimitMatch)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(RateLimitService)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpGlobalRateLimit.
func (in *HttpGlobalRateLimit) DeepCopy() *HttpGlobalRateLimit {
	if in == nil {
		return nil
	}
	out := new(HttpGlobalRateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpLocalRateLimit) DeepCopyInto(out *HttpLocalRateLimit) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]*RateLimitRule, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(RateLimitRule)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpLocalRateLimit.
func (in *HttpLocalRateLimit) DeepCopy() *HttpLocalRateLimit {
	if in == nil {
		return nil
	}
	out := new(HttpLocalRateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressSelector) DeepCopyInto(out *IngressSelector) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(uint32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressSelector.
func (in *IngressSelector) DeepCopy() *IngressSelector {
	if in == nil {
		return nil
	}
	out := new(IngressSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MethodMatch) DeepCopyInto(out *MethodMatch) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MethodMatch.
func (in *MethodMatch) DeepCopy() *MethodMatch {
	if in == nil {
		return nil
	}
	out := new(MethodMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimit) DeepCopyInto(out *RateLimit) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimit.
func (in *RateLimit) DeepCopy() *RateLimit {
	if in == nil {
		return nil
	}
	out := new(RateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RateLimit) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitList) DeepCopyInto(out *RateLimitList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RateLimit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitList.
func (in *RateLimitList) DeepCopy() *RateLimitList {
	if in == nil {
		return nil
	}
	out := new(RateLimitList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RateLimitList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitMatch) DeepCopyInto(out *RateLimitMatch) {
	*out = *in
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(UrlMatch)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(SourceMatch)
		**out = **in
	}
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(MethodMatch)
		**out = **in
	}
	if in.Header != nil {
		in, out := &in.Header, &out.Header
		*out = new(HeaderMatch)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitMatch.
func (in *RateLimitMatch) DeepCopy() *RateLimitMatch {
	if in == nil {
		return nil
	}
	out := new(RateLimitMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitRule) DeepCopyInto(out *RateLimitRule) {
	*out = *in
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]*RateLimitMatch, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(RateLimitMatch)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(RatelimitPolicy)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitRule.
func (in *RateLimitRule) DeepCopy() *RateLimitRule {
	if in == nil {
		return nil
	}
	out := new(RateLimitRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitService) DeepCopyInto(out *RateLimitService) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitService.
func (in *RateLimitService) DeepCopy() *RateLimitService {
	if in == nil {
		return nil
	}
	out := new(RateLimitService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitSpec) DeepCopyInto(out *RateLimitSpec) {
	*out = *in
	if in.WorkloadSelector != nil {
		in, out := &in.WorkloadSelector, &out.WorkloadSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(IngressSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = new(EgressSelector)
		**out = **in
	}
	if in.HttpLocalRateLimit != nil {
		in, out := &in.HttpLocalRateLimit, &out.HttpLocalRateLimit
		*out = new(HttpLocalRateLimit)
		(*in).DeepCopyInto(*out)
	}
	if in.HttpGlobalRateLimit != nil {
		in, out := &in.HttpGlobalRateLimit, &out.HttpGlobalRateLimit
		*out = new(HttpGlobalRateLimit)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitSpec.
func (in *RateLimitSpec) DeepCopy() *RateLimitSpec {
	if in == nil {
		return nil
	}
	out := new(RateLimitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitStatus) DeepCopyInto(out *RateLimitStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitStatus.
func (in *RateLimitStatus) DeepCopy() *RateLimitStatus {
	if in == nil {
		return nil
	}
	out := new(RateLimitStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RatelimitPolicy) DeepCopyInto(out *RatelimitPolicy) {
	*out = *in
	out.Interval = in.Interval
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RatelimitPolicy.
func (in *RatelimitPolicy) DeepCopy() *RatelimitPolicy {
	if in == nil {
		return nil
	}
	out := new(RatelimitPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceMatch) DeepCopyInto(out *SourceMatch) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceMatch.
func (in *SourceMatch) DeepCopy() *SourceMatch {
	if in == nil {
		return nil
	}
	out := new(SourceMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UrlMatch) DeepCopyInto(out *UrlMatch) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UrlMatch.
func (in *UrlMatch) DeepCopy() *UrlMatch {
	if in == nil {
		return nil
	}
	out := new(UrlMatch)
	in.DeepCopyInto(out)
	return out
}
