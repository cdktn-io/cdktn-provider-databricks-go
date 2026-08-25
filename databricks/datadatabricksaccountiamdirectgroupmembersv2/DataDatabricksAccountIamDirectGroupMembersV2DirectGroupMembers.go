// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountiamdirectgroupmembersv2


type DataDatabricksAccountIamDirectGroupMembersV2DirectGroupMembers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/account_iam_direct_group_members_v2#group_id DataDatabricksAccountIamDirectGroupMembersV2#group_id}.
	GroupId *float64 `field:"required" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/account_iam_direct_group_members_v2#principal_id DataDatabricksAccountIamDirectGroupMembersV2#principal_id}.
	PrincipalId *float64 `field:"required" json:"principalId" yaml:"principalId"`
}

