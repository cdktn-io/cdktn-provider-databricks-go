// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2


type WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkMode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#color WorkspaceSettingV2#color}.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#color_scheme WorkspaceSettingV2#color_scheme}.
	ColorScheme *string `field:"optional" json:"colorScheme" yaml:"colorScheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#custom_diverging WorkspaceSettingV2#custom_diverging}.
	CustomDiverging *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomDiverging `field:"optional" json:"customDiverging" yaml:"customDiverging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#custom_sequential WorkspaceSettingV2#custom_sequential}.
	CustomSequential *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientDarkModeCustomSequential `field:"optional" json:"customSequential" yaml:"customSequential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/workspace_setting_v2#from_categorical_index WorkspaceSettingV2#from_categorical_index}.
	FromCategoricalIndex *float64 `field:"optional" json:"fromCategoricalIndex" yaml:"fromCategoricalIndex"`
}

