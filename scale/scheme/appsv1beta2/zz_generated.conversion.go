//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package appsv1beta2

import (
	v1beta2 "k8s.io/api/apps/v1beta2"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	scheme "github.com/spotmaxtech/k8s-client-go-v0260/scale/scheme"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1beta2.Scale)(nil), (*scheme.Scale)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_Scale_To_scheme_Scale(a.(*v1beta2.Scale), b.(*scheme.Scale), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*scheme.Scale)(nil), (*v1beta2.Scale)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_scheme_Scale_To_v1beta2_Scale(a.(*scheme.Scale), b.(*v1beta2.Scale), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.ScaleSpec)(nil), (*scheme.ScaleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_ScaleSpec_To_scheme_ScaleSpec(a.(*v1beta2.ScaleSpec), b.(*scheme.ScaleSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*scheme.ScaleSpec)(nil), (*v1beta2.ScaleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_scheme_ScaleSpec_To_v1beta2_ScaleSpec(a.(*scheme.ScaleSpec), b.(*v1beta2.ScaleSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*scheme.ScaleStatus)(nil), (*v1beta2.ScaleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_scheme_ScaleStatus_To_v1beta2_ScaleStatus(a.(*scheme.ScaleStatus), b.(*v1beta2.ScaleStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta2.ScaleStatus)(nil), (*scheme.ScaleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_ScaleStatus_To_scheme_ScaleStatus(a.(*v1beta2.ScaleStatus), b.(*scheme.ScaleStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta2_Scale_To_scheme_Scale(in *v1beta2.Scale, out *scheme.Scale, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta2_ScaleSpec_To_scheme_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta2_ScaleStatus_To_scheme_ScaleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta2_Scale_To_scheme_Scale is an autogenerated conversion function.
func Convert_v1beta2_Scale_To_scheme_Scale(in *v1beta2.Scale, out *scheme.Scale, s conversion.Scope) error {
	return autoConvert_v1beta2_Scale_To_scheme_Scale(in, out, s)
}

func autoConvert_scheme_Scale_To_v1beta2_Scale(in *scheme.Scale, out *v1beta2.Scale, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_scheme_ScaleSpec_To_v1beta2_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_scheme_ScaleStatus_To_v1beta2_ScaleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_scheme_Scale_To_v1beta2_Scale is an autogenerated conversion function.
func Convert_scheme_Scale_To_v1beta2_Scale(in *scheme.Scale, out *v1beta2.Scale, s conversion.Scope) error {
	return autoConvert_scheme_Scale_To_v1beta2_Scale(in, out, s)
}

func autoConvert_v1beta2_ScaleSpec_To_scheme_ScaleSpec(in *v1beta2.ScaleSpec, out *scheme.ScaleSpec, s conversion.Scope) error {
	out.Replicas = in.Replicas
	return nil
}

// Convert_v1beta2_ScaleSpec_To_scheme_ScaleSpec is an autogenerated conversion function.
func Convert_v1beta2_ScaleSpec_To_scheme_ScaleSpec(in *v1beta2.ScaleSpec, out *scheme.ScaleSpec, s conversion.Scope) error {
	return autoConvert_v1beta2_ScaleSpec_To_scheme_ScaleSpec(in, out, s)
}

func autoConvert_scheme_ScaleSpec_To_v1beta2_ScaleSpec(in *scheme.ScaleSpec, out *v1beta2.ScaleSpec, s conversion.Scope) error {
	out.Replicas = in.Replicas
	return nil
}

// Convert_scheme_ScaleSpec_To_v1beta2_ScaleSpec is an autogenerated conversion function.
func Convert_scheme_ScaleSpec_To_v1beta2_ScaleSpec(in *scheme.ScaleSpec, out *v1beta2.ScaleSpec, s conversion.Scope) error {
	return autoConvert_scheme_ScaleSpec_To_v1beta2_ScaleSpec(in, out, s)
}

func autoConvert_v1beta2_ScaleStatus_To_scheme_ScaleStatus(in *v1beta2.ScaleStatus, out *scheme.ScaleStatus, s conversion.Scope) error {
	out.Replicas = in.Replicas
	// WARNING: in.Selector requires manual conversion: inconvertible types (map[string]string vs *k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector)
	// WARNING: in.TargetSelector requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_scheme_ScaleStatus_To_v1beta2_ScaleStatus(in *scheme.ScaleStatus, out *v1beta2.ScaleStatus, s conversion.Scope) error {
	out.Replicas = in.Replicas
	// WARNING: in.Selector requires manual conversion: inconvertible types (*k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector vs map[string]string)
	return nil
}
