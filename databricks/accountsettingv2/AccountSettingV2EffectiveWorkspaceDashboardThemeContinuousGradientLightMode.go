// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2


type AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightMode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/account_setting_v2#color AccountSettingV2#color}.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/account_setting_v2#color_scheme AccountSettingV2#color_scheme}.
	ColorScheme *string `field:"optional" json:"colorScheme" yaml:"colorScheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/account_setting_v2#custom_diverging AccountSettingV2#custom_diverging}.
	CustomDiverging *AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomDiverging `field:"optional" json:"customDiverging" yaml:"customDiverging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/account_setting_v2#custom_sequential AccountSettingV2#custom_sequential}.
	CustomSequential *AccountSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientLightModeCustomSequential `field:"optional" json:"customSequential" yaml:"customSequential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/account_setting_v2#from_categorical_index AccountSettingV2#from_categorical_index}.
	FromCategoricalIndex *float64 `field:"optional" json:"fromCategoricalIndex" yaml:"fromCategoricalIndex"`
}

