// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspaceiamdirectgroupmembersv2


type DataDatabricksWorkspaceIamDirectGroupMembersV2DirectGroupMembers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/workspace_iam_direct_group_members_v2#group_id DataDatabricksWorkspaceIamDirectGroupMembersV2#group_id}.
	GroupId *float64 `field:"required" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/workspace_iam_direct_group_members_v2#principal_id DataDatabricksWorkspaceIamDirectGroupMembersV2#principal_id}.
	PrincipalId *float64 `field:"required" json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/workspace_iam_direct_group_members_v2#provider_config DataDatabricksWorkspaceIamDirectGroupMembersV2#provider_config}.
	ProviderConfig *DataDatabricksWorkspaceIamDirectGroupMembersV2DirectGroupMembersProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

