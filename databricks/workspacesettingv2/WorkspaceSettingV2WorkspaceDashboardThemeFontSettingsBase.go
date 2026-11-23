// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2


type WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsBase struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#font_color WorkspaceSettingV2#font_color}.
	FontColor *WorkspaceSettingV2WorkspaceDashboardThemeFontSettingsBaseFontColor `field:"optional" json:"fontColor" yaml:"fontColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#font_family WorkspaceSettingV2#font_family}.
	FontFamily *string `field:"optional" json:"fontFamily" yaml:"fontFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#font_size WorkspaceSettingV2#font_size}.
	FontSize *float64 `field:"optional" json:"fontSize" yaml:"fontSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#font_style WorkspaceSettingV2#font_style}.
	FontStyle *string `field:"optional" json:"fontStyle" yaml:"fontStyle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#font_weight WorkspaceSettingV2#font_weight}.
	FontWeight *float64 `field:"optional" json:"fontWeight" yaml:"fontWeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#text_case WorkspaceSettingV2#text_case}.
	TextCase *string `field:"optional" json:"textCase" yaml:"textCase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#text_decoration WorkspaceSettingV2#text_decoration}.
	TextDecoration *string `field:"optional" json:"textDecoration" yaml:"textDecoration"`
}

