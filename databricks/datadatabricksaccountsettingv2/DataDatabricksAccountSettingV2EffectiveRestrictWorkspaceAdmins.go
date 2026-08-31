// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountsettingv2


type DataDatabricksAccountSettingV2EffectiveRestrictWorkspaceAdmins struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/account_setting_v2#status DataDatabricksAccountSettingV2#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/account_setting_v2#disable_gov_tag_creation DataDatabricksAccountSettingV2#disable_gov_tag_creation}.
	DisableGovTagCreation interface{} `field:"optional" json:"disableGovTagCreation" yaml:"disableGovTagCreation"`
}

