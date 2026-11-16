// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspacesettingv2


type DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightMode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/workspace_setting_v2#color DataDatabricksWorkspaceSettingV2#color}.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/workspace_setting_v2#color_scheme DataDatabricksWorkspaceSettingV2#color_scheme}.
	ColorScheme *string `field:"optional" json:"colorScheme" yaml:"colorScheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/workspace_setting_v2#custom_diverging DataDatabricksWorkspaceSettingV2#custom_diverging}.
	CustomDiverging *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomDiverging `field:"optional" json:"customDiverging" yaml:"customDiverging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/workspace_setting_v2#custom_sequential DataDatabricksWorkspaceSettingV2#custom_sequential}.
	CustomSequential *DataDatabricksWorkspaceSettingV2WorkspaceDashboardThemeContinuousGradientLightModeCustomSequential `field:"optional" json:"customSequential" yaml:"customSequential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/workspace_setting_v2#from_categorical_index DataDatabricksWorkspaceSettingV2#from_categorical_index}.
	FromCategoricalIndex *float64 `field:"optional" json:"fromCategoricalIndex" yaml:"fromCategoricalIndex"`
}

