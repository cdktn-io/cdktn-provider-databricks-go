// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2


type WorkspaceSettingV2WorkspaceDashboardThemeSentimentColors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#negative WorkspaceSettingV2#negative}.
	Negative *WorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsNegative `field:"optional" json:"negative" yaml:"negative"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#neutral WorkspaceSettingV2#neutral}.
	Neutral *WorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsNeutral `field:"optional" json:"neutral" yaml:"neutral"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/workspace_setting_v2#positive WorkspaceSettingV2#positive}.
	Positive *WorkspaceSettingV2WorkspaceDashboardThemeSentimentColorsPositive `field:"optional" json:"positive" yaml:"positive"`
}

