// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2


type WorkspaceSettingV2EffectiveWorkspaceLabel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/workspace_setting_v2#color WorkspaceSettingV2#color}.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/workspace_setting_v2#label WorkspaceSettingV2#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
}

