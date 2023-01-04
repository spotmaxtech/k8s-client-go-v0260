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

package v1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "github.com/spotmaxtech/k8s-client-go-v0260/tools/clientcmd/api"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AuthInfo)(nil), (*api.AuthInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AuthInfo_To_api_AuthInfo(a.(*AuthInfo), b.(*api.AuthInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.AuthInfo)(nil), (*AuthInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_AuthInfo_To_v1_AuthInfo(a.(*api.AuthInfo), b.(*AuthInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AuthProviderConfig)(nil), (*api.AuthProviderConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AuthProviderConfig_To_api_AuthProviderConfig(a.(*AuthProviderConfig), b.(*api.AuthProviderConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.AuthProviderConfig)(nil), (*AuthProviderConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_AuthProviderConfig_To_v1_AuthProviderConfig(a.(*api.AuthProviderConfig), b.(*AuthProviderConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Cluster)(nil), (*api.Cluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Cluster_To_api_Cluster(a.(*Cluster), b.(*api.Cluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.Cluster)(nil), (*Cluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_Cluster_To_v1_Cluster(a.(*api.Cluster), b.(*Cluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Config)(nil), (*api.Config)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Config_To_api_Config(a.(*Config), b.(*api.Config), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.Config)(nil), (*Config)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_Config_To_v1_Config(a.(*api.Config), b.(*Config), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Context)(nil), (*api.Context)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Context_To_api_Context(a.(*Context), b.(*api.Context), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.Context)(nil), (*Context)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_Context_To_v1_Context(a.(*api.Context), b.(*Context), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ExecConfig)(nil), (*api.ExecConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ExecConfig_To_api_ExecConfig(a.(*ExecConfig), b.(*api.ExecConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.ExecConfig)(nil), (*ExecConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_ExecConfig_To_v1_ExecConfig(a.(*api.ExecConfig), b.(*ExecConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ExecEnvVar)(nil), (*api.ExecEnvVar)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ExecEnvVar_To_api_ExecEnvVar(a.(*ExecEnvVar), b.(*api.ExecEnvVar), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.ExecEnvVar)(nil), (*ExecEnvVar)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_ExecEnvVar_To_v1_ExecEnvVar(a.(*api.ExecEnvVar), b.(*ExecEnvVar), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Preferences)(nil), (*api.Preferences)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Preferences_To_api_Preferences(a.(*Preferences), b.(*api.Preferences), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.Preferences)(nil), (*Preferences)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_Preferences_To_v1_Preferences(a.(*api.Preferences), b.(*Preferences), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*map[string]*api.AuthInfo)(nil), (*[]NamedAuthInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_Map_string_To_Pointer_api_AuthInfo_To_Slice_v1_NamedAuthInfo(a.(*map[string]*api.AuthInfo), b.(*[]NamedAuthInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*map[string]*api.Cluster)(nil), (*[]NamedCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_Map_string_To_Pointer_api_Cluster_To_Slice_v1_NamedCluster(a.(*map[string]*api.Cluster), b.(*[]NamedCluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*map[string]*api.Context)(nil), (*[]NamedContext)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_Map_string_To_Pointer_api_Context_To_Slice_v1_NamedContext(a.(*map[string]*api.Context), b.(*[]NamedContext), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*map[string]runtime.Object)(nil), (*[]NamedExtension)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_Map_string_To_runtime_Object_To_Slice_v1_NamedExtension(a.(*map[string]runtime.Object), b.(*[]NamedExtension), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*[]NamedAuthInfo)(nil), (*map[string]*api.AuthInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_Slice_v1_NamedAuthInfo_To_Map_string_To_Pointer_api_AuthInfo(a.(*[]NamedAuthInfo), b.(*map[string]*api.AuthInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*[]NamedCluster)(nil), (*map[string]*api.Cluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_Slice_v1_NamedCluster_To_Map_string_To_Pointer_api_Cluster(a.(*[]NamedCluster), b.(*map[string]*api.Cluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*[]NamedContext)(nil), (*map[string]*api.Context)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_Slice_v1_NamedContext_To_Map_string_To_Pointer_api_Context(a.(*[]NamedContext), b.(*map[string]*api.Context), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*[]NamedExtension)(nil), (*map[string]runtime.Object)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_Slice_v1_NamedExtension_To_Map_string_To_runtime_Object(a.(*[]NamedExtension), b.(*map[string]runtime.Object), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_AuthInfo_To_api_AuthInfo(in *AuthInfo, out *api.AuthInfo, s conversion.Scope) error {
	out.ClientCertificate = in.ClientCertificate
	out.ClientCertificateData = *(*[]byte)(unsafe.Pointer(&in.ClientCertificateData))
	out.ClientKey = in.ClientKey
	out.ClientKeyData = *(*[]byte)(unsafe.Pointer(&in.ClientKeyData))
	out.Token = in.Token
	out.TokenFile = in.TokenFile
	out.Impersonate = in.Impersonate
	out.ImpersonateUID = in.ImpersonateUID
	out.ImpersonateGroups = *(*[]string)(unsafe.Pointer(&in.ImpersonateGroups))
	out.ImpersonateUserExtra = *(*map[string][]string)(unsafe.Pointer(&in.ImpersonateUserExtra))
	out.Username = in.Username
	out.Password = in.Password
	out.AuthProvider = (*api.AuthProviderConfig)(unsafe.Pointer(in.AuthProvider))
	if in.Exec != nil {
		in, out := &in.Exec, &out.Exec
		*out = new(api.ExecConfig)
		if err := Convert_v1_ExecConfig_To_api_ExecConfig(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Exec = nil
	}
	if err := Convert_Slice_v1_NamedExtension_To_Map_string_To_runtime_Object(&in.Extensions, &out.Extensions, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_AuthInfo_To_api_AuthInfo is an autogenerated conversion function.
func Convert_v1_AuthInfo_To_api_AuthInfo(in *AuthInfo, out *api.AuthInfo, s conversion.Scope) error {
	return autoConvert_v1_AuthInfo_To_api_AuthInfo(in, out, s)
}

func autoConvert_api_AuthInfo_To_v1_AuthInfo(in *api.AuthInfo, out *AuthInfo, s conversion.Scope) error {
	// INFO: in.LocationOfOrigin opted out of conversion generation
	out.ClientCertificate = in.ClientCertificate
	out.ClientCertificateData = *(*[]byte)(unsafe.Pointer(&in.ClientCertificateData))
	out.ClientKey = in.ClientKey
	out.ClientKeyData = *(*[]byte)(unsafe.Pointer(&in.ClientKeyData))
	out.Token = in.Token
	out.TokenFile = in.TokenFile
	out.Impersonate = in.Impersonate
	out.ImpersonateUID = in.ImpersonateUID
	out.ImpersonateGroups = *(*[]string)(unsafe.Pointer(&in.ImpersonateGroups))
	out.ImpersonateUserExtra = *(*map[string][]string)(unsafe.Pointer(&in.ImpersonateUserExtra))
	out.Username = in.Username
	out.Password = in.Password
	out.AuthProvider = (*AuthProviderConfig)(unsafe.Pointer(in.AuthProvider))
	if in.Exec != nil {
		in, out := &in.Exec, &out.Exec
		*out = new(ExecConfig)
		if err := Convert_api_ExecConfig_To_v1_ExecConfig(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Exec = nil
	}
	if err := Convert_Map_string_To_runtime_Object_To_Slice_v1_NamedExtension(&in.Extensions, &out.Extensions, s); err != nil {
		return err
	}
	return nil
}

// Convert_api_AuthInfo_To_v1_AuthInfo is an autogenerated conversion function.
func Convert_api_AuthInfo_To_v1_AuthInfo(in *api.AuthInfo, out *AuthInfo, s conversion.Scope) error {
	return autoConvert_api_AuthInfo_To_v1_AuthInfo(in, out, s)
}

func autoConvert_v1_AuthProviderConfig_To_api_AuthProviderConfig(in *AuthProviderConfig, out *api.AuthProviderConfig, s conversion.Scope) error {
	out.Name = in.Name
	out.Config = *(*map[string]string)(unsafe.Pointer(&in.Config))
	return nil
}

// Convert_v1_AuthProviderConfig_To_api_AuthProviderConfig is an autogenerated conversion function.
func Convert_v1_AuthProviderConfig_To_api_AuthProviderConfig(in *AuthProviderConfig, out *api.AuthProviderConfig, s conversion.Scope) error {
	return autoConvert_v1_AuthProviderConfig_To_api_AuthProviderConfig(in, out, s)
}

func autoConvert_api_AuthProviderConfig_To_v1_AuthProviderConfig(in *api.AuthProviderConfig, out *AuthProviderConfig, s conversion.Scope) error {
	out.Name = in.Name
	out.Config = *(*map[string]string)(unsafe.Pointer(&in.Config))
	return nil
}

// Convert_api_AuthProviderConfig_To_v1_AuthProviderConfig is an autogenerated conversion function.
func Convert_api_AuthProviderConfig_To_v1_AuthProviderConfig(in *api.AuthProviderConfig, out *AuthProviderConfig, s conversion.Scope) error {
	return autoConvert_api_AuthProviderConfig_To_v1_AuthProviderConfig(in, out, s)
}

func autoConvert_v1_Cluster_To_api_Cluster(in *Cluster, out *api.Cluster, s conversion.Scope) error {
	out.Server = in.Server
	out.TLSServerName = in.TLSServerName
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CertificateAuthority = in.CertificateAuthority
	out.CertificateAuthorityData = *(*[]byte)(unsafe.Pointer(&in.CertificateAuthorityData))
	out.ProxyURL = in.ProxyURL
	out.DisableCompression = in.DisableCompression
	if err := Convert_Slice_v1_NamedExtension_To_Map_string_To_runtime_Object(&in.Extensions, &out.Extensions, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Cluster_To_api_Cluster is an autogenerated conversion function.
func Convert_v1_Cluster_To_api_Cluster(in *Cluster, out *api.Cluster, s conversion.Scope) error {
	return autoConvert_v1_Cluster_To_api_Cluster(in, out, s)
}

func autoConvert_api_Cluster_To_v1_Cluster(in *api.Cluster, out *Cluster, s conversion.Scope) error {
	// INFO: in.LocationOfOrigin opted out of conversion generation
	out.Server = in.Server
	out.TLSServerName = in.TLSServerName
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CertificateAuthority = in.CertificateAuthority
	out.CertificateAuthorityData = *(*[]byte)(unsafe.Pointer(&in.CertificateAuthorityData))
	out.ProxyURL = in.ProxyURL
	out.DisableCompression = in.DisableCompression
	if err := Convert_Map_string_To_runtime_Object_To_Slice_v1_NamedExtension(&in.Extensions, &out.Extensions, s); err != nil {
		return err
	}
	return nil
}

// Convert_api_Cluster_To_v1_Cluster is an autogenerated conversion function.
func Convert_api_Cluster_To_v1_Cluster(in *api.Cluster, out *Cluster, s conversion.Scope) error {
	return autoConvert_api_Cluster_To_v1_Cluster(in, out, s)
}

func autoConvert_v1_Config_To_api_Config(in *Config, out *api.Config, s conversion.Scope) error {
	// INFO: in.Kind opted out of conversion generation
	// INFO: in.APIVersion opted out of conversion generation
	if err := Convert_v1_Preferences_To_api_Preferences(&in.Preferences, &out.Preferences, s); err != nil {
		return err
	}
	if err := Convert_Slice_v1_NamedCluster_To_Map_string_To_Pointer_api_Cluster(&in.Clusters, &out.Clusters, s); err != nil {
		return err
	}
	if err := Convert_Slice_v1_NamedAuthInfo_To_Map_string_To_Pointer_api_AuthInfo(&in.AuthInfos, &out.AuthInfos, s); err != nil {
		return err
	}
	if err := Convert_Slice_v1_NamedContext_To_Map_string_To_Pointer_api_Context(&in.Contexts, &out.Contexts, s); err != nil {
		return err
	}
	out.CurrentContext = in.CurrentContext
	if err := Convert_Slice_v1_NamedExtension_To_Map_string_To_runtime_Object(&in.Extensions, &out.Extensions, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Config_To_api_Config is an autogenerated conversion function.
func Convert_v1_Config_To_api_Config(in *Config, out *api.Config, s conversion.Scope) error {
	return autoConvert_v1_Config_To_api_Config(in, out, s)
}

func autoConvert_api_Config_To_v1_Config(in *api.Config, out *Config, s conversion.Scope) error {
	// INFO: in.Kind opted out of conversion generation
	// INFO: in.APIVersion opted out of conversion generation
	if err := Convert_api_Preferences_To_v1_Preferences(&in.Preferences, &out.Preferences, s); err != nil {
		return err
	}
	if err := Convert_Map_string_To_Pointer_api_Cluster_To_Slice_v1_NamedCluster(&in.Clusters, &out.Clusters, s); err != nil {
		return err
	}
	if err := Convert_Map_string_To_Pointer_api_AuthInfo_To_Slice_v1_NamedAuthInfo(&in.AuthInfos, &out.AuthInfos, s); err != nil {
		return err
	}
	if err := Convert_Map_string_To_Pointer_api_Context_To_Slice_v1_NamedContext(&in.Contexts, &out.Contexts, s); err != nil {
		return err
	}
	out.CurrentContext = in.CurrentContext
	if err := Convert_Map_string_To_runtime_Object_To_Slice_v1_NamedExtension(&in.Extensions, &out.Extensions, s); err != nil {
		return err
	}
	return nil
}

// Convert_api_Config_To_v1_Config is an autogenerated conversion function.
func Convert_api_Config_To_v1_Config(in *api.Config, out *Config, s conversion.Scope) error {
	return autoConvert_api_Config_To_v1_Config(in, out, s)
}

func autoConvert_v1_Context_To_api_Context(in *Context, out *api.Context, s conversion.Scope) error {
	out.Cluster = in.Cluster
	out.AuthInfo = in.AuthInfo
	out.Namespace = in.Namespace
	if err := Convert_Slice_v1_NamedExtension_To_Map_string_To_runtime_Object(&in.Extensions, &out.Extensions, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Context_To_api_Context is an autogenerated conversion function.
func Convert_v1_Context_To_api_Context(in *Context, out *api.Context, s conversion.Scope) error {
	return autoConvert_v1_Context_To_api_Context(in, out, s)
}

func autoConvert_api_Context_To_v1_Context(in *api.Context, out *Context, s conversion.Scope) error {
	// INFO: in.LocationOfOrigin opted out of conversion generation
	out.Cluster = in.Cluster
	out.AuthInfo = in.AuthInfo
	out.Namespace = in.Namespace
	if err := Convert_Map_string_To_runtime_Object_To_Slice_v1_NamedExtension(&in.Extensions, &out.Extensions, s); err != nil {
		return err
	}
	return nil
}

// Convert_api_Context_To_v1_Context is an autogenerated conversion function.
func Convert_api_Context_To_v1_Context(in *api.Context, out *Context, s conversion.Scope) error {
	return autoConvert_api_Context_To_v1_Context(in, out, s)
}

func autoConvert_v1_ExecConfig_To_api_ExecConfig(in *ExecConfig, out *api.ExecConfig, s conversion.Scope) error {
	out.Command = in.Command
	out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
	out.Env = *(*[]api.ExecEnvVar)(unsafe.Pointer(&in.Env))
	out.APIVersion = in.APIVersion
	out.InstallHint = in.InstallHint
	out.ProvideClusterInfo = in.ProvideClusterInfo
	out.InteractiveMode = api.ExecInteractiveMode(in.InteractiveMode)
	return nil
}

// Convert_v1_ExecConfig_To_api_ExecConfig is an autogenerated conversion function.
func Convert_v1_ExecConfig_To_api_ExecConfig(in *ExecConfig, out *api.ExecConfig, s conversion.Scope) error {
	return autoConvert_v1_ExecConfig_To_api_ExecConfig(in, out, s)
}

func autoConvert_api_ExecConfig_To_v1_ExecConfig(in *api.ExecConfig, out *ExecConfig, s conversion.Scope) error {
	out.Command = in.Command
	out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
	out.Env = *(*[]ExecEnvVar)(unsafe.Pointer(&in.Env))
	out.APIVersion = in.APIVersion
	out.InstallHint = in.InstallHint
	out.ProvideClusterInfo = in.ProvideClusterInfo
	// INFO: in.Config opted out of conversion generation
	out.InteractiveMode = ExecInteractiveMode(in.InteractiveMode)
	// INFO: in.StdinUnavailable opted out of conversion generation
	// INFO: in.StdinUnavailableMessage opted out of conversion generation
	return nil
}

// Convert_api_ExecConfig_To_v1_ExecConfig is an autogenerated conversion function.
func Convert_api_ExecConfig_To_v1_ExecConfig(in *api.ExecConfig, out *ExecConfig, s conversion.Scope) error {
	return autoConvert_api_ExecConfig_To_v1_ExecConfig(in, out, s)
}

func autoConvert_v1_ExecEnvVar_To_api_ExecEnvVar(in *ExecEnvVar, out *api.ExecEnvVar, s conversion.Scope) error {
	out.Name = in.Name
	out.Value = in.Value
	return nil
}

// Convert_v1_ExecEnvVar_To_api_ExecEnvVar is an autogenerated conversion function.
func Convert_v1_ExecEnvVar_To_api_ExecEnvVar(in *ExecEnvVar, out *api.ExecEnvVar, s conversion.Scope) error {
	return autoConvert_v1_ExecEnvVar_To_api_ExecEnvVar(in, out, s)
}

func autoConvert_api_ExecEnvVar_To_v1_ExecEnvVar(in *api.ExecEnvVar, out *ExecEnvVar, s conversion.Scope) error {
	out.Name = in.Name
	out.Value = in.Value
	return nil
}

// Convert_api_ExecEnvVar_To_v1_ExecEnvVar is an autogenerated conversion function.
func Convert_api_ExecEnvVar_To_v1_ExecEnvVar(in *api.ExecEnvVar, out *ExecEnvVar, s conversion.Scope) error {
	return autoConvert_api_ExecEnvVar_To_v1_ExecEnvVar(in, out, s)
}

func autoConvert_v1_Preferences_To_api_Preferences(in *Preferences, out *api.Preferences, s conversion.Scope) error {
	out.Colors = in.Colors
	if err := Convert_Slice_v1_NamedExtension_To_Map_string_To_runtime_Object(&in.Extensions, &out.Extensions, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Preferences_To_api_Preferences is an autogenerated conversion function.
func Convert_v1_Preferences_To_api_Preferences(in *Preferences, out *api.Preferences, s conversion.Scope) error {
	return autoConvert_v1_Preferences_To_api_Preferences(in, out, s)
}

func autoConvert_api_Preferences_To_v1_Preferences(in *api.Preferences, out *Preferences, s conversion.Scope) error {
	out.Colors = in.Colors
	if err := Convert_Map_string_To_runtime_Object_To_Slice_v1_NamedExtension(&in.Extensions, &out.Extensions, s); err != nil {
		return err
	}
	return nil
}

// Convert_api_Preferences_To_v1_Preferences is an autogenerated conversion function.
func Convert_api_Preferences_To_v1_Preferences(in *api.Preferences, out *Preferences, s conversion.Scope) error {
	return autoConvert_api_Preferences_To_v1_Preferences(in, out, s)
}
