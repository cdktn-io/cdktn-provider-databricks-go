// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package restrictworkspaceadminssetting


type RestrictWorkspaceAdminsSettingRestrictWorkspaceAdmins struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/restrict_workspace_admins_setting#status RestrictWorkspaceAdminsSetting#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/restrict_workspace_admins_setting#disable_gov_tag_creation RestrictWorkspaceAdminsSetting#disable_gov_tag_creation}.
	DisableGovTagCreation interface{} `field:"optional" json:"disableGovTagCreation" yaml:"disableGovTagCreation"`
}

