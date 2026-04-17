// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2


type WorkspaceSettingV2EffectiveRestrictWorkspaceAdmins struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/workspace_setting_v2#status WorkspaceSettingV2#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/workspace_setting_v2#disable_gov_tag_creation WorkspaceSettingV2#disable_gov_tag_creation}.
	DisableGovTagCreation interface{} `field:"optional" json:"disableGovTagCreation" yaml:"disableGovTagCreation"`
}

