// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2


type WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradient struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#dark_mode WorkspaceSettingV2#dark_mode}.
	DarkMode *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkMode `field:"optional" json:"darkMode" yaml:"darkMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#light_mode WorkspaceSettingV2#light_mode}.
	LightMode *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightMode `field:"optional" json:"lightMode" yaml:"lightMode"`
}

