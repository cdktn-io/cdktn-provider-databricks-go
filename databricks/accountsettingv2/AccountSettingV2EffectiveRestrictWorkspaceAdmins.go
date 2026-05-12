// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2


type AccountSettingV2EffectiveRestrictWorkspaceAdmins struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/account_setting_v2#status AccountSettingV2#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/account_setting_v2#disable_gov_tag_creation AccountSettingV2#disable_gov_tag_creation}.
	DisableGovTagCreation interface{} `field:"optional" json:"disableGovTagCreation" yaml:"disableGovTagCreation"`
}

