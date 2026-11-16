// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountsettingv2


type DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightMode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/account_setting_v2#color DataDatabricksAccountSettingV2#color}.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/account_setting_v2#color_scheme DataDatabricksAccountSettingV2#color_scheme}.
	ColorScheme *string `field:"optional" json:"colorScheme" yaml:"colorScheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/account_setting_v2#custom_diverging DataDatabricksAccountSettingV2#custom_diverging}.
	CustomDiverging *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDiverging `field:"optional" json:"customDiverging" yaml:"customDiverging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/account_setting_v2#custom_sequential DataDatabricksAccountSettingV2#custom_sequential}.
	CustomSequential *DataDatabricksAccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomSequential `field:"optional" json:"customSequential" yaml:"customSequential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/account_setting_v2#from_categorical_index DataDatabricksAccountSettingV2#from_categorical_index}.
	FromCategoricalIndex *float64 `field:"optional" json:"fromCategoricalIndex" yaml:"fromCategoricalIndex"`
}

