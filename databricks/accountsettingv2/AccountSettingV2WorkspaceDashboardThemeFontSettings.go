// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2


type AccountSettingV2WorkspaceDashboardThemeFontSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/account_setting_v2#base AccountSettingV2#base}.
	Base *AccountSettingV2WorkspaceDashboardThemeFontSettingsBase `field:"optional" json:"base" yaml:"base"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/account_setting_v2#field_title AccountSettingV2#field_title}.
	FieldTitle *AccountSettingV2WorkspaceDashboardThemeFontSettingsFieldTitle `field:"optional" json:"fieldTitle" yaml:"fieldTitle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/account_setting_v2#field_value AccountSettingV2#field_value}.
	FieldValue *AccountSettingV2WorkspaceDashboardThemeFontSettingsFieldValue `field:"optional" json:"fieldValue" yaml:"fieldValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/account_setting_v2#widget_description AccountSettingV2#widget_description}.
	WidgetDescription *AccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetDescription `field:"optional" json:"widgetDescription" yaml:"widgetDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/account_setting_v2#widget_title AccountSettingV2#widget_title}.
	WidgetTitle *AccountSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitle `field:"optional" json:"widgetTitle" yaml:"widgetTitle"`
}

