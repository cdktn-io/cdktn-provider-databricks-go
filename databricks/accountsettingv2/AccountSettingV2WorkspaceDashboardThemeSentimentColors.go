// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2


type AccountSettingV2WorkspaceDashboardThemeSentimentColors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/account_setting_v2#negative AccountSettingV2#negative}.
	Negative *AccountSettingV2WorkspaceDashboardThemeSentimentColorsNegative `field:"optional" json:"negative" yaml:"negative"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/account_setting_v2#neutral AccountSettingV2#neutral}.
	Neutral *AccountSettingV2WorkspaceDashboardThemeSentimentColorsNeutral `field:"optional" json:"neutral" yaml:"neutral"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/account_setting_v2#positive AccountSettingV2#positive}.
	Positive *AccountSettingV2WorkspaceDashboardThemeSentimentColorsPositive `field:"optional" json:"positive" yaml:"positive"`
}

