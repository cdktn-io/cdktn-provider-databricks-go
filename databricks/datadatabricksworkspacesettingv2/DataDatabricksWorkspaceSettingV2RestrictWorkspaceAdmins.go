// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspacesettingv2


type DataDatabricksWorkspaceSettingV2RestrictWorkspaceAdmins struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/workspace_setting_v2#status DataDatabricksWorkspaceSettingV2#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/workspace_setting_v2#disable_gov_tag_creation DataDatabricksWorkspaceSettingV2#disable_gov_tag_creation}.
	DisableGovTagCreation interface{} `field:"optional" json:"disableGovTagCreation" yaml:"disableGovTagCreation"`
}

