// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2


type WorkspaceSettingV2WorkspaceDashboardThemeFontSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#base WorkspaceSettingV2#base}.
	Base *WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsBase `field:"optional" json:"base" yaml:"base"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#field_title WorkspaceSettingV2#field_title}.
	FieldTitle *WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldTitle `field:"optional" json:"fieldTitle" yaml:"fieldTitle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#field_value WorkspaceSettingV2#field_value}.
	FieldValue *WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsFieldValue `field:"optional" json:"fieldValue" yaml:"fieldValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#widget_description WorkspaceSettingV2#widget_description}.
	WidgetDescription *WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsWidgetDescription `field:"optional" json:"widgetDescription" yaml:"widgetDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#widget_title WorkspaceSettingV2#widget_title}.
	WidgetTitle *WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsWidgetTitle `field:"optional" json:"widgetTitle" yaml:"widgetTitle"`
}

