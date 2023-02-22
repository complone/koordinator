//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Koordinator Authors.

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

package v1alpha2

import (
	unsafe "unsafe"

	config "github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	v1alpha1 "k8s.io/component-base/config/v1alpha1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*DefaultEvictorArgs)(nil), (*config.DefaultEvictorArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_DefaultEvictorArgs_To_config_DefaultEvictorArgs(a.(*DefaultEvictorArgs), b.(*config.DefaultEvictorArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.DefaultEvictorArgs)(nil), (*DefaultEvictorArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_DefaultEvictorArgs_To_v1alpha2_DefaultEvictorArgs(a.(*config.DefaultEvictorArgs), b.(*DefaultEvictorArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DeschedulerProfile)(nil), (*config.DeschedulerProfile)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_DeschedulerProfile_To_config_DeschedulerProfile(a.(*DeschedulerProfile), b.(*config.DeschedulerProfile), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.DeschedulerProfile)(nil), (*DeschedulerProfile)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_DeschedulerProfile_To_v1alpha2_DeschedulerProfile(a.(*config.DeschedulerProfile), b.(*DeschedulerProfile), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LoadAnomalyCondition)(nil), (*config.LoadAnomalyCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_LoadAnomalyCondition_To_config_LoadAnomalyCondition(a.(*LoadAnomalyCondition), b.(*config.LoadAnomalyCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LoadAnomalyCondition)(nil), (*LoadAnomalyCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LoadAnomalyCondition_To_v1alpha2_LoadAnomalyCondition(a.(*config.LoadAnomalyCondition), b.(*LoadAnomalyCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LowNodeLoadArgs)(nil), (*config.LowNodeLoadArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_LowNodeLoadArgs_To_config_LowNodeLoadArgs(a.(*LowNodeLoadArgs), b.(*config.LowNodeLoadArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LowNodeLoadArgs)(nil), (*LowNodeLoadArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LowNodeLoadArgs_To_v1alpha2_LowNodeLoadArgs(a.(*config.LowNodeLoadArgs), b.(*LowNodeLoadArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LowNodeLoadPodSelector)(nil), (*config.LowNodeLoadPodSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_LowNodeLoadPodSelector_To_config_LowNodeLoadPodSelector(a.(*LowNodeLoadPodSelector), b.(*config.LowNodeLoadPodSelector), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LowNodeLoadPodSelector)(nil), (*LowNodeLoadPodSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LowNodeLoadPodSelector_To_v1alpha2_LowNodeLoadPodSelector(a.(*config.LowNodeLoadPodSelector), b.(*LowNodeLoadPodSelector), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MigrationControllerArgs)(nil), (*config.MigrationControllerArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MigrationControllerArgs_To_config_MigrationControllerArgs(a.(*MigrationControllerArgs), b.(*config.MigrationControllerArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.MigrationControllerArgs)(nil), (*MigrationControllerArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MigrationControllerArgs_To_v1alpha2_MigrationControllerArgs(a.(*config.MigrationControllerArgs), b.(*MigrationControllerArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MigrationObjectLimiter)(nil), (*config.MigrationObjectLimiter)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MigrationObjectLimiter_To_config_MigrationObjectLimiter(a.(*MigrationObjectLimiter), b.(*config.MigrationObjectLimiter), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.MigrationObjectLimiter)(nil), (*MigrationObjectLimiter)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MigrationObjectLimiter_To_v1alpha2_MigrationObjectLimiter(a.(*config.MigrationObjectLimiter), b.(*MigrationObjectLimiter), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Namespaces)(nil), (*config.Namespaces)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_Namespaces_To_config_Namespaces(a.(*Namespaces), b.(*config.Namespaces), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Namespaces)(nil), (*Namespaces)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Namespaces_To_v1alpha2_Namespaces(a.(*config.Namespaces), b.(*Namespaces), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Plugin)(nil), (*config.Plugin)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_Plugin_To_config_Plugin(a.(*Plugin), b.(*config.Plugin), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Plugin)(nil), (*Plugin)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Plugin_To_v1alpha2_Plugin(a.(*config.Plugin), b.(*Plugin), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PluginConfig)(nil), (*config.PluginConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_PluginConfig_To_config_PluginConfig(a.(*PluginConfig), b.(*config.PluginConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PluginConfig)(nil), (*PluginConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PluginConfig_To_v1alpha2_PluginConfig(a.(*config.PluginConfig), b.(*PluginConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PluginSet)(nil), (*config.PluginSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_PluginSet_To_config_PluginSet(a.(*PluginSet), b.(*config.PluginSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PluginSet)(nil), (*PluginSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PluginSet_To_v1alpha2_PluginSet(a.(*config.PluginSet), b.(*PluginSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Plugins)(nil), (*config.Plugins)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_Plugins_To_config_Plugins(a.(*Plugins), b.(*config.Plugins), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Plugins)(nil), (*Plugins)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Plugins_To_v1alpha2_Plugins(a.(*config.Plugins), b.(*Plugins), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PriorityThreshold)(nil), (*config.PriorityThreshold)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_PriorityThreshold_To_config_PriorityThreshold(a.(*PriorityThreshold), b.(*config.PriorityThreshold), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PriorityThreshold)(nil), (*PriorityThreshold)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PriorityThreshold_To_v1alpha2_PriorityThreshold(a.(*config.PriorityThreshold), b.(*PriorityThreshold), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RemovePodsViolatingNodeAffinityArgs)(nil), (*config.RemovePodsViolatingNodeAffinityArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_RemovePodsViolatingNodeAffinityArgs_To_config_RemovePodsViolatingNodeAffinityArgs(a.(*RemovePodsViolatingNodeAffinityArgs), b.(*config.RemovePodsViolatingNodeAffinityArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.RemovePodsViolatingNodeAffinityArgs)(nil), (*RemovePodsViolatingNodeAffinityArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_RemovePodsViolatingNodeAffinityArgs_To_v1alpha2_RemovePodsViolatingNodeAffinityArgs(a.(*config.RemovePodsViolatingNodeAffinityArgs), b.(*RemovePodsViolatingNodeAffinityArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.DeschedulerConfiguration)(nil), (*DeschedulerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_DeschedulerConfiguration_To_v1alpha2_DeschedulerConfiguration(a.(*config.DeschedulerConfiguration), b.(*DeschedulerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*DeschedulerConfiguration)(nil), (*config.DeschedulerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_DeschedulerConfiguration_To_config_DeschedulerConfiguration(a.(*DeschedulerConfiguration), b.(*config.DeschedulerConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha2_DefaultEvictorArgs_To_config_DefaultEvictorArgs(in *DefaultEvictorArgs, out *config.DefaultEvictorArgs, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.DryRun, &out.DryRun, s); err != nil {
		return err
	}
	out.MaxNoOfPodsToEvictPerNode = (*int)(unsafe.Pointer(in.MaxNoOfPodsToEvictPerNode))
	out.MaxNoOfPodsToEvictPerNamespace = (*int)(unsafe.Pointer(in.MaxNoOfPodsToEvictPerNamespace))
	out.EvictFailedBarePods = in.EvictFailedBarePods
	out.EvictLocalStoragePods = in.EvictLocalStoragePods
	out.EvictSystemCriticalPods = in.EvictSystemCriticalPods
	out.IgnorePvcPods = in.IgnorePvcPods
	out.NodeFit = in.NodeFit
	out.PriorityThreshold = (*config.PriorityThreshold)(unsafe.Pointer(in.PriorityThreshold))
	out.LabelSelector = (*v1.LabelSelector)(unsafe.Pointer(in.LabelSelector))
	return nil
}

// Convert_v1alpha2_DefaultEvictorArgs_To_config_DefaultEvictorArgs is an autogenerated conversion function.
func Convert_v1alpha2_DefaultEvictorArgs_To_config_DefaultEvictorArgs(in *DefaultEvictorArgs, out *config.DefaultEvictorArgs, s conversion.Scope) error {
	return autoConvert_v1alpha2_DefaultEvictorArgs_To_config_DefaultEvictorArgs(in, out, s)
}

func autoConvert_config_DefaultEvictorArgs_To_v1alpha2_DefaultEvictorArgs(in *config.DefaultEvictorArgs, out *DefaultEvictorArgs, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.DryRun, &out.DryRun, s); err != nil {
		return err
	}
	out.MaxNoOfPodsToEvictPerNode = (*int)(unsafe.Pointer(in.MaxNoOfPodsToEvictPerNode))
	out.MaxNoOfPodsToEvictPerNamespace = (*int)(unsafe.Pointer(in.MaxNoOfPodsToEvictPerNamespace))
	out.EvictFailedBarePods = in.EvictFailedBarePods
	out.EvictLocalStoragePods = in.EvictLocalStoragePods
	out.EvictSystemCriticalPods = in.EvictSystemCriticalPods
	out.IgnorePvcPods = in.IgnorePvcPods
	out.NodeFit = in.NodeFit
	out.PriorityThreshold = (*PriorityThreshold)(unsafe.Pointer(in.PriorityThreshold))
	out.LabelSelector = (*v1.LabelSelector)(unsafe.Pointer(in.LabelSelector))
	return nil
}

// Convert_config_DefaultEvictorArgs_To_v1alpha2_DefaultEvictorArgs is an autogenerated conversion function.
func Convert_config_DefaultEvictorArgs_To_v1alpha2_DefaultEvictorArgs(in *config.DefaultEvictorArgs, out *DefaultEvictorArgs, s conversion.Scope) error {
	return autoConvert_config_DefaultEvictorArgs_To_v1alpha2_DefaultEvictorArgs(in, out, s)
}

func autoConvert_v1alpha2_DeschedulerConfiguration_To_config_DeschedulerConfiguration(in *DeschedulerConfiguration, out *config.DeschedulerConfiguration, s conversion.Scope) error {
	if err := v1alpha1.Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	if err := v1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	if err := v1alpha1.Convert_v1alpha1_DebuggingConfiguration_To_config_DebuggingConfiguration(&in.DebuggingConfiguration, &out.DebuggingConfiguration, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_string_To_string(&in.HealthzBindAddress, &out.HealthzBindAddress, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_string_To_string(&in.MetricsBindAddress, &out.MetricsBindAddress, s); err != nil {
		return err
	}
	out.DeschedulingInterval = in.DeschedulingInterval
	out.DryRun = in.DryRun
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]config.DeschedulerProfile, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_DeschedulerProfile_To_config_DeschedulerProfile(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Profiles = nil
	}
	out.NodeSelector = (*v1.LabelSelector)(unsafe.Pointer(in.NodeSelector))
	return nil
}

func autoConvert_config_DeschedulerConfiguration_To_v1alpha2_DeschedulerConfiguration(in *config.DeschedulerConfiguration, out *DeschedulerConfiguration, s conversion.Scope) error {
	if err := v1alpha1.Convert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	if err := v1alpha1.Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	if err := v1alpha1.Convert_config_DebuggingConfiguration_To_v1alpha1_DebuggingConfiguration(&in.DebuggingConfiguration, &out.DebuggingConfiguration, s); err != nil {
		return err
	}
	if err := v1.Convert_string_To_Pointer_string(&in.HealthzBindAddress, &out.HealthzBindAddress, s); err != nil {
		return err
	}
	if err := v1.Convert_string_To_Pointer_string(&in.MetricsBindAddress, &out.MetricsBindAddress, s); err != nil {
		return err
	}
	out.DeschedulingInterval = in.DeschedulingInterval
	out.DryRun = in.DryRun
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]DeschedulerProfile, len(*in))
		for i := range *in {
			if err := Convert_config_DeschedulerProfile_To_v1alpha2_DeschedulerProfile(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Profiles = nil
	}
	out.NodeSelector = (*v1.LabelSelector)(unsafe.Pointer(in.NodeSelector))
	return nil
}

func autoConvert_v1alpha2_DeschedulerProfile_To_config_DeschedulerProfile(in *DeschedulerProfile, out *config.DeschedulerProfile, s conversion.Scope) error {
	out.Name = in.Name
	if in.PluginConfig != nil {
		in, out := &in.PluginConfig, &out.PluginConfig
		*out = make([]config.PluginConfig, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_PluginConfig_To_config_PluginConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.PluginConfig = nil
	}
	out.Plugins = (*config.Plugins)(unsafe.Pointer(in.Plugins))
	return nil
}

// Convert_v1alpha2_DeschedulerProfile_To_config_DeschedulerProfile is an autogenerated conversion function.
func Convert_v1alpha2_DeschedulerProfile_To_config_DeschedulerProfile(in *DeschedulerProfile, out *config.DeschedulerProfile, s conversion.Scope) error {
	return autoConvert_v1alpha2_DeschedulerProfile_To_config_DeschedulerProfile(in, out, s)
}

func autoConvert_config_DeschedulerProfile_To_v1alpha2_DeschedulerProfile(in *config.DeschedulerProfile, out *DeschedulerProfile, s conversion.Scope) error {
	out.Name = in.Name
	if in.PluginConfig != nil {
		in, out := &in.PluginConfig, &out.PluginConfig
		*out = make([]PluginConfig, len(*in))
		for i := range *in {
			if err := Convert_config_PluginConfig_To_v1alpha2_PluginConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.PluginConfig = nil
	}
	out.Plugins = (*Plugins)(unsafe.Pointer(in.Plugins))
	return nil
}

// Convert_config_DeschedulerProfile_To_v1alpha2_DeschedulerProfile is an autogenerated conversion function.
func Convert_config_DeschedulerProfile_To_v1alpha2_DeschedulerProfile(in *config.DeschedulerProfile, out *DeschedulerProfile, s conversion.Scope) error {
	return autoConvert_config_DeschedulerProfile_To_v1alpha2_DeschedulerProfile(in, out, s)
}

func autoConvert_v1alpha2_LoadAnomalyCondition_To_config_LoadAnomalyCondition(in *LoadAnomalyCondition, out *config.LoadAnomalyCondition, s conversion.Scope) error {
	if err := v1.Convert_Pointer_v1_Duration_To_v1_Duration(&in.Timeout, &out.Timeout, s); err != nil {
		return err
	}
	out.ConsecutiveAbnormalities = in.ConsecutiveAbnormalities
	return nil
}

// Convert_v1alpha2_LoadAnomalyCondition_To_config_LoadAnomalyCondition is an autogenerated conversion function.
func Convert_v1alpha2_LoadAnomalyCondition_To_config_LoadAnomalyCondition(in *LoadAnomalyCondition, out *config.LoadAnomalyCondition, s conversion.Scope) error {
	return autoConvert_v1alpha2_LoadAnomalyCondition_To_config_LoadAnomalyCondition(in, out, s)
}

func autoConvert_config_LoadAnomalyCondition_To_v1alpha2_LoadAnomalyCondition(in *config.LoadAnomalyCondition, out *LoadAnomalyCondition, s conversion.Scope) error {
	if err := v1.Convert_v1_Duration_To_Pointer_v1_Duration(&in.Timeout, &out.Timeout, s); err != nil {
		return err
	}
	out.ConsecutiveAbnormalities = in.ConsecutiveAbnormalities
	return nil
}

// Convert_config_LoadAnomalyCondition_To_v1alpha2_LoadAnomalyCondition is an autogenerated conversion function.
func Convert_config_LoadAnomalyCondition_To_v1alpha2_LoadAnomalyCondition(in *config.LoadAnomalyCondition, out *LoadAnomalyCondition, s conversion.Scope) error {
	return autoConvert_config_LoadAnomalyCondition_To_v1alpha2_LoadAnomalyCondition(in, out, s)
}

func autoConvert_v1alpha2_LowNodeLoadArgs_To_config_LowNodeLoadArgs(in *LowNodeLoadArgs, out *config.LowNodeLoadArgs, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.Paused, &out.Paused, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_bool_To_bool(&in.DryRun, &out.DryRun, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int32_To_int32(&in.NumberOfNodes, &out.NumberOfNodes, s); err != nil {
		return err
	}
	out.EvictableNamespaces = (*config.Namespaces)(unsafe.Pointer(in.EvictableNamespaces))
	out.NodeSelector = (*v1.LabelSelector)(unsafe.Pointer(in.NodeSelector))
	out.PodSelectors = *(*[]config.LowNodeLoadPodSelector)(unsafe.Pointer(&in.PodSelectors))
	if err := v1.Convert_Pointer_bool_To_bool(&in.NodeFit, &out.NodeFit, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_bool_To_bool(&in.UseDeviationThresholds, &out.UseDeviationThresholds, s); err != nil {
		return err
	}
	out.HighThresholds = *(*config.ResourceThresholds)(unsafe.Pointer(&in.HighThresholds))
	out.LowThresholds = *(*config.ResourceThresholds)(unsafe.Pointer(&in.LowThresholds))
	if in.AnomalyCondition != nil {
		in, out := &in.AnomalyCondition, &out.AnomalyCondition
		*out = new(config.LoadAnomalyCondition)
		if err := Convert_v1alpha2_LoadAnomalyCondition_To_config_LoadAnomalyCondition(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.AnomalyCondition = nil
	}
	return nil
}

// Convert_v1alpha2_LowNodeLoadArgs_To_config_LowNodeLoadArgs is an autogenerated conversion function.
func Convert_v1alpha2_LowNodeLoadArgs_To_config_LowNodeLoadArgs(in *LowNodeLoadArgs, out *config.LowNodeLoadArgs, s conversion.Scope) error {
	return autoConvert_v1alpha2_LowNodeLoadArgs_To_config_LowNodeLoadArgs(in, out, s)
}

func autoConvert_config_LowNodeLoadArgs_To_v1alpha2_LowNodeLoadArgs(in *config.LowNodeLoadArgs, out *LowNodeLoadArgs, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.Paused, &out.Paused, s); err != nil {
		return err
	}
	if err := v1.Convert_bool_To_Pointer_bool(&in.DryRun, &out.DryRun, s); err != nil {
		return err
	}
	if err := v1.Convert_int32_To_Pointer_int32(&in.NumberOfNodes, &out.NumberOfNodes, s); err != nil {
		return err
	}
	out.EvictableNamespaces = (*Namespaces)(unsafe.Pointer(in.EvictableNamespaces))
	out.NodeSelector = (*v1.LabelSelector)(unsafe.Pointer(in.NodeSelector))
	out.PodSelectors = *(*[]LowNodeLoadPodSelector)(unsafe.Pointer(&in.PodSelectors))
	if err := v1.Convert_bool_To_Pointer_bool(&in.NodeFit, &out.NodeFit, s); err != nil {
		return err
	}
	if err := v1.Convert_bool_To_Pointer_bool(&in.UseDeviationThresholds, &out.UseDeviationThresholds, s); err != nil {
		return err
	}
	out.HighThresholds = *(*ResourceThresholds)(unsafe.Pointer(&in.HighThresholds))
	out.LowThresholds = *(*ResourceThresholds)(unsafe.Pointer(&in.LowThresholds))
	if in.AnomalyCondition != nil {
		in, out := &in.AnomalyCondition, &out.AnomalyCondition
		*out = new(LoadAnomalyCondition)
		if err := Convert_config_LoadAnomalyCondition_To_v1alpha2_LoadAnomalyCondition(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.AnomalyCondition = nil
	}
	return nil
}

// Convert_config_LowNodeLoadArgs_To_v1alpha2_LowNodeLoadArgs is an autogenerated conversion function.
func Convert_config_LowNodeLoadArgs_To_v1alpha2_LowNodeLoadArgs(in *config.LowNodeLoadArgs, out *LowNodeLoadArgs, s conversion.Scope) error {
	return autoConvert_config_LowNodeLoadArgs_To_v1alpha2_LowNodeLoadArgs(in, out, s)
}

func autoConvert_v1alpha2_LowNodeLoadPodSelector_To_config_LowNodeLoadPodSelector(in *LowNodeLoadPodSelector, out *config.LowNodeLoadPodSelector, s conversion.Scope) error {
	out.Name = in.Name
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	return nil
}

// Convert_v1alpha2_LowNodeLoadPodSelector_To_config_LowNodeLoadPodSelector is an autogenerated conversion function.
func Convert_v1alpha2_LowNodeLoadPodSelector_To_config_LowNodeLoadPodSelector(in *LowNodeLoadPodSelector, out *config.LowNodeLoadPodSelector, s conversion.Scope) error {
	return autoConvert_v1alpha2_LowNodeLoadPodSelector_To_config_LowNodeLoadPodSelector(in, out, s)
}

func autoConvert_config_LowNodeLoadPodSelector_To_v1alpha2_LowNodeLoadPodSelector(in *config.LowNodeLoadPodSelector, out *LowNodeLoadPodSelector, s conversion.Scope) error {
	out.Name = in.Name
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	return nil
}

// Convert_config_LowNodeLoadPodSelector_To_v1alpha2_LowNodeLoadPodSelector is an autogenerated conversion function.
func Convert_config_LowNodeLoadPodSelector_To_v1alpha2_LowNodeLoadPodSelector(in *config.LowNodeLoadPodSelector, out *LowNodeLoadPodSelector, s conversion.Scope) error {
	return autoConvert_config_LowNodeLoadPodSelector_To_v1alpha2_LowNodeLoadPodSelector(in, out, s)
}

func autoConvert_v1alpha2_MigrationControllerArgs_To_config_MigrationControllerArgs(in *MigrationControllerArgs, out *config.MigrationControllerArgs, s conversion.Scope) error {
	out.DryRun = in.DryRun
	if err := v1.Convert_Pointer_int32_To_int32(&in.MaxConcurrentReconciles, &out.MaxConcurrentReconciles, s); err != nil {
		return err
	}
	out.EvictFailedBarePods = in.EvictFailedBarePods
	out.EvictLocalStoragePods = in.EvictLocalStoragePods
	out.EvictSystemCriticalPods = in.EvictSystemCriticalPods
	out.IgnorePvcPods = in.IgnorePvcPods
	out.LabelSelector = (*v1.LabelSelector)(unsafe.Pointer(in.LabelSelector))
	out.Namespaces = (*config.Namespaces)(unsafe.Pointer(in.Namespaces))
	out.MaxMigratingPerNode = (*int32)(unsafe.Pointer(in.MaxMigratingPerNode))
	out.MaxMigratingPerNamespace = (*int32)(unsafe.Pointer(in.MaxMigratingPerNamespace))
	out.MaxMigratingPerWorkload = (*intstr.IntOrString)(unsafe.Pointer(in.MaxMigratingPerWorkload))
	out.MaxUnavailablePerWorkload = (*intstr.IntOrString)(unsafe.Pointer(in.MaxUnavailablePerWorkload))
	out.ObjectLimiters = *(*config.ObjectLimiterMap)(unsafe.Pointer(&in.ObjectLimiters))
	out.DefaultJobMode = in.DefaultJobMode
	if err := v1.Convert_Pointer_v1_Duration_To_v1_Duration(&in.DefaultJobTTL, &out.DefaultJobTTL, s); err != nil {
		return err
	}
	out.EvictQPS = (*config.Float64OrString)(unsafe.Pointer(in.EvictQPS))
	if err := v1.Convert_Pointer_int32_To_int32(&in.EvictBurst, &out.EvictBurst, s); err != nil {
		return err
	}
	out.EvictionPolicy = in.EvictionPolicy
	out.DefaultDeleteOptions = (*v1.DeleteOptions)(unsafe.Pointer(in.DefaultDeleteOptions))
	return nil
}

// Convert_v1alpha2_MigrationControllerArgs_To_config_MigrationControllerArgs is an autogenerated conversion function.
func Convert_v1alpha2_MigrationControllerArgs_To_config_MigrationControllerArgs(in *MigrationControllerArgs, out *config.MigrationControllerArgs, s conversion.Scope) error {
	return autoConvert_v1alpha2_MigrationControllerArgs_To_config_MigrationControllerArgs(in, out, s)
}

func autoConvert_config_MigrationControllerArgs_To_v1alpha2_MigrationControllerArgs(in *config.MigrationControllerArgs, out *MigrationControllerArgs, s conversion.Scope) error {
	out.DryRun = in.DryRun
	if err := v1.Convert_int32_To_Pointer_int32(&in.MaxConcurrentReconciles, &out.MaxConcurrentReconciles, s); err != nil {
		return err
	}
	out.EvictFailedBarePods = in.EvictFailedBarePods
	out.EvictLocalStoragePods = in.EvictLocalStoragePods
	out.EvictSystemCriticalPods = in.EvictSystemCriticalPods
	out.IgnorePvcPods = in.IgnorePvcPods
	out.LabelSelector = (*v1.LabelSelector)(unsafe.Pointer(in.LabelSelector))
	out.Namespaces = (*Namespaces)(unsafe.Pointer(in.Namespaces))
	out.MaxMigratingPerNode = (*int32)(unsafe.Pointer(in.MaxMigratingPerNode))
	out.MaxMigratingPerNamespace = (*int32)(unsafe.Pointer(in.MaxMigratingPerNamespace))
	out.MaxMigratingPerWorkload = (*intstr.IntOrString)(unsafe.Pointer(in.MaxMigratingPerWorkload))
	out.MaxUnavailablePerWorkload = (*intstr.IntOrString)(unsafe.Pointer(in.MaxUnavailablePerWorkload))
	out.ObjectLimiters = *(*ObjectLimiterMap)(unsafe.Pointer(&in.ObjectLimiters))
	out.DefaultJobMode = in.DefaultJobMode
	if err := v1.Convert_v1_Duration_To_Pointer_v1_Duration(&in.DefaultJobTTL, &out.DefaultJobTTL, s); err != nil {
		return err
	}
	out.EvictQPS = (*config.Float64OrString)(unsafe.Pointer(in.EvictQPS))
	if err := v1.Convert_int32_To_Pointer_int32(&in.EvictBurst, &out.EvictBurst, s); err != nil {
		return err
	}
	out.EvictionPolicy = in.EvictionPolicy
	out.DefaultDeleteOptions = (*v1.DeleteOptions)(unsafe.Pointer(in.DefaultDeleteOptions))
	return nil
}

// Convert_config_MigrationControllerArgs_To_v1alpha2_MigrationControllerArgs is an autogenerated conversion function.
func Convert_config_MigrationControllerArgs_To_v1alpha2_MigrationControllerArgs(in *config.MigrationControllerArgs, out *MigrationControllerArgs, s conversion.Scope) error {
	return autoConvert_config_MigrationControllerArgs_To_v1alpha2_MigrationControllerArgs(in, out, s)
}

func autoConvert_v1alpha2_MigrationObjectLimiter_To_config_MigrationObjectLimiter(in *MigrationObjectLimiter, out *config.MigrationObjectLimiter, s conversion.Scope) error {
	out.Duration = in.Duration
	out.MaxMigrating = (*intstr.IntOrString)(unsafe.Pointer(in.MaxMigrating))
	return nil
}

// Convert_v1alpha2_MigrationObjectLimiter_To_config_MigrationObjectLimiter is an autogenerated conversion function.
func Convert_v1alpha2_MigrationObjectLimiter_To_config_MigrationObjectLimiter(in *MigrationObjectLimiter, out *config.MigrationObjectLimiter, s conversion.Scope) error {
	return autoConvert_v1alpha2_MigrationObjectLimiter_To_config_MigrationObjectLimiter(in, out, s)
}

func autoConvert_config_MigrationObjectLimiter_To_v1alpha2_MigrationObjectLimiter(in *config.MigrationObjectLimiter, out *MigrationObjectLimiter, s conversion.Scope) error {
	out.Duration = in.Duration
	out.MaxMigrating = (*intstr.IntOrString)(unsafe.Pointer(in.MaxMigrating))
	return nil
}

// Convert_config_MigrationObjectLimiter_To_v1alpha2_MigrationObjectLimiter is an autogenerated conversion function.
func Convert_config_MigrationObjectLimiter_To_v1alpha2_MigrationObjectLimiter(in *config.MigrationObjectLimiter, out *MigrationObjectLimiter, s conversion.Scope) error {
	return autoConvert_config_MigrationObjectLimiter_To_v1alpha2_MigrationObjectLimiter(in, out, s)
}

func autoConvert_v1alpha2_Namespaces_To_config_Namespaces(in *Namespaces, out *config.Namespaces, s conversion.Scope) error {
	out.Include = *(*[]string)(unsafe.Pointer(&in.Include))
	out.Exclude = *(*[]string)(unsafe.Pointer(&in.Exclude))
	return nil
}

// Convert_v1alpha2_Namespaces_To_config_Namespaces is an autogenerated conversion function.
func Convert_v1alpha2_Namespaces_To_config_Namespaces(in *Namespaces, out *config.Namespaces, s conversion.Scope) error {
	return autoConvert_v1alpha2_Namespaces_To_config_Namespaces(in, out, s)
}

func autoConvert_config_Namespaces_To_v1alpha2_Namespaces(in *config.Namespaces, out *Namespaces, s conversion.Scope) error {
	out.Include = *(*[]string)(unsafe.Pointer(&in.Include))
	out.Exclude = *(*[]string)(unsafe.Pointer(&in.Exclude))
	return nil
}

// Convert_config_Namespaces_To_v1alpha2_Namespaces is an autogenerated conversion function.
func Convert_config_Namespaces_To_v1alpha2_Namespaces(in *config.Namespaces, out *Namespaces, s conversion.Scope) error {
	return autoConvert_config_Namespaces_To_v1alpha2_Namespaces(in, out, s)
}

func autoConvert_v1alpha2_Plugin_To_config_Plugin(in *Plugin, out *config.Plugin, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_v1alpha2_Plugin_To_config_Plugin is an autogenerated conversion function.
func Convert_v1alpha2_Plugin_To_config_Plugin(in *Plugin, out *config.Plugin, s conversion.Scope) error {
	return autoConvert_v1alpha2_Plugin_To_config_Plugin(in, out, s)
}

func autoConvert_config_Plugin_To_v1alpha2_Plugin(in *config.Plugin, out *Plugin, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_config_Plugin_To_v1alpha2_Plugin is an autogenerated conversion function.
func Convert_config_Plugin_To_v1alpha2_Plugin(in *config.Plugin, out *Plugin, s conversion.Scope) error {
	return autoConvert_config_Plugin_To_v1alpha2_Plugin(in, out, s)
}

func autoConvert_v1alpha2_PluginConfig_To_config_PluginConfig(in *PluginConfig, out *config.PluginConfig, s conversion.Scope) error {
	out.Name = in.Name
	if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.Args, &out.Args, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_PluginConfig_To_config_PluginConfig is an autogenerated conversion function.
func Convert_v1alpha2_PluginConfig_To_config_PluginConfig(in *PluginConfig, out *config.PluginConfig, s conversion.Scope) error {
	return autoConvert_v1alpha2_PluginConfig_To_config_PluginConfig(in, out, s)
}

func autoConvert_config_PluginConfig_To_v1alpha2_PluginConfig(in *config.PluginConfig, out *PluginConfig, s conversion.Scope) error {
	out.Name = in.Name
	if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.Args, &out.Args, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_PluginConfig_To_v1alpha2_PluginConfig is an autogenerated conversion function.
func Convert_config_PluginConfig_To_v1alpha2_PluginConfig(in *config.PluginConfig, out *PluginConfig, s conversion.Scope) error {
	return autoConvert_config_PluginConfig_To_v1alpha2_PluginConfig(in, out, s)
}

func autoConvert_v1alpha2_PluginSet_To_config_PluginSet(in *PluginSet, out *config.PluginSet, s conversion.Scope) error {
	out.Enabled = *(*[]config.Plugin)(unsafe.Pointer(&in.Enabled))
	out.Disabled = *(*[]config.Plugin)(unsafe.Pointer(&in.Disabled))
	return nil
}

// Convert_v1alpha2_PluginSet_To_config_PluginSet is an autogenerated conversion function.
func Convert_v1alpha2_PluginSet_To_config_PluginSet(in *PluginSet, out *config.PluginSet, s conversion.Scope) error {
	return autoConvert_v1alpha2_PluginSet_To_config_PluginSet(in, out, s)
}

func autoConvert_config_PluginSet_To_v1alpha2_PluginSet(in *config.PluginSet, out *PluginSet, s conversion.Scope) error {
	out.Enabled = *(*[]Plugin)(unsafe.Pointer(&in.Enabled))
	out.Disabled = *(*[]Plugin)(unsafe.Pointer(&in.Disabled))
	return nil
}

// Convert_config_PluginSet_To_v1alpha2_PluginSet is an autogenerated conversion function.
func Convert_config_PluginSet_To_v1alpha2_PluginSet(in *config.PluginSet, out *PluginSet, s conversion.Scope) error {
	return autoConvert_config_PluginSet_To_v1alpha2_PluginSet(in, out, s)
}

func autoConvert_v1alpha2_Plugins_To_config_Plugins(in *Plugins, out *config.Plugins, s conversion.Scope) error {
	if err := Convert_v1alpha2_PluginSet_To_config_PluginSet(&in.Deschedule, &out.Deschedule, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_PluginSet_To_config_PluginSet(&in.Balance, &out.Balance, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_PluginSet_To_config_PluginSet(&in.Evictor, &out.Evictor, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_Plugins_To_config_Plugins is an autogenerated conversion function.
func Convert_v1alpha2_Plugins_To_config_Plugins(in *Plugins, out *config.Plugins, s conversion.Scope) error {
	return autoConvert_v1alpha2_Plugins_To_config_Plugins(in, out, s)
}

func autoConvert_config_Plugins_To_v1alpha2_Plugins(in *config.Plugins, out *Plugins, s conversion.Scope) error {
	if err := Convert_config_PluginSet_To_v1alpha2_PluginSet(&in.Deschedule, &out.Deschedule, s); err != nil {
		return err
	}
	if err := Convert_config_PluginSet_To_v1alpha2_PluginSet(&in.Balance, &out.Balance, s); err != nil {
		return err
	}
	if err := Convert_config_PluginSet_To_v1alpha2_PluginSet(&in.Evictor, &out.Evictor, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_Plugins_To_v1alpha2_Plugins is an autogenerated conversion function.
func Convert_config_Plugins_To_v1alpha2_Plugins(in *config.Plugins, out *Plugins, s conversion.Scope) error {
	return autoConvert_config_Plugins_To_v1alpha2_Plugins(in, out, s)
}

func autoConvert_v1alpha2_PriorityThreshold_To_config_PriorityThreshold(in *PriorityThreshold, out *config.PriorityThreshold, s conversion.Scope) error {
	out.Value = (*int32)(unsafe.Pointer(in.Value))
	out.Name = in.Name
	return nil
}

// Convert_v1alpha2_PriorityThreshold_To_config_PriorityThreshold is an autogenerated conversion function.
func Convert_v1alpha2_PriorityThreshold_To_config_PriorityThreshold(in *PriorityThreshold, out *config.PriorityThreshold, s conversion.Scope) error {
	return autoConvert_v1alpha2_PriorityThreshold_To_config_PriorityThreshold(in, out, s)
}

func autoConvert_config_PriorityThreshold_To_v1alpha2_PriorityThreshold(in *config.PriorityThreshold, out *PriorityThreshold, s conversion.Scope) error {
	out.Value = (*int32)(unsafe.Pointer(in.Value))
	out.Name = in.Name
	return nil
}

// Convert_config_PriorityThreshold_To_v1alpha2_PriorityThreshold is an autogenerated conversion function.
func Convert_config_PriorityThreshold_To_v1alpha2_PriorityThreshold(in *config.PriorityThreshold, out *PriorityThreshold, s conversion.Scope) error {
	return autoConvert_config_PriorityThreshold_To_v1alpha2_PriorityThreshold(in, out, s)
}

func autoConvert_v1alpha2_RemovePodsViolatingNodeAffinityArgs_To_config_RemovePodsViolatingNodeAffinityArgs(in *RemovePodsViolatingNodeAffinityArgs, out *config.RemovePodsViolatingNodeAffinityArgs, s conversion.Scope) error {
	out.Namespaces = (*config.Namespaces)(unsafe.Pointer(in.Namespaces))
	out.LabelSelector = (*v1.LabelSelector)(unsafe.Pointer(in.LabelSelector))
	out.NodeAffinityType = *(*[]string)(unsafe.Pointer(&in.NodeAffinityType))
	return nil
}

// Convert_v1alpha2_RemovePodsViolatingNodeAffinityArgs_To_config_RemovePodsViolatingNodeAffinityArgs is an autogenerated conversion function.
func Convert_v1alpha2_RemovePodsViolatingNodeAffinityArgs_To_config_RemovePodsViolatingNodeAffinityArgs(in *RemovePodsViolatingNodeAffinityArgs, out *config.RemovePodsViolatingNodeAffinityArgs, s conversion.Scope) error {
	return autoConvert_v1alpha2_RemovePodsViolatingNodeAffinityArgs_To_config_RemovePodsViolatingNodeAffinityArgs(in, out, s)
}

func autoConvert_config_RemovePodsViolatingNodeAffinityArgs_To_v1alpha2_RemovePodsViolatingNodeAffinityArgs(in *config.RemovePodsViolatingNodeAffinityArgs, out *RemovePodsViolatingNodeAffinityArgs, s conversion.Scope) error {
	out.Namespaces = (*Namespaces)(unsafe.Pointer(in.Namespaces))
	out.LabelSelector = (*v1.LabelSelector)(unsafe.Pointer(in.LabelSelector))
	out.NodeAffinityType = *(*[]string)(unsafe.Pointer(&in.NodeAffinityType))
	return nil
}

// Convert_config_RemovePodsViolatingNodeAffinityArgs_To_v1alpha2_RemovePodsViolatingNodeAffinityArgs is an autogenerated conversion function.
func Convert_config_RemovePodsViolatingNodeAffinityArgs_To_v1alpha2_RemovePodsViolatingNodeAffinityArgs(in *config.RemovePodsViolatingNodeAffinityArgs, out *RemovePodsViolatingNodeAffinityArgs, s conversion.Scope) error {
	return autoConvert_config_RemovePodsViolatingNodeAffinityArgs_To_v1alpha2_RemovePodsViolatingNodeAffinityArgs(in, out, s)
}
