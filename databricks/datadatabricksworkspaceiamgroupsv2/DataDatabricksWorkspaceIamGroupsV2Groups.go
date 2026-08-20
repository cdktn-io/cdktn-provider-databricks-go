// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspaceiamgroupsv2


type DataDatabricksWorkspaceIamGroupsV2Groups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/workspace_iam_groups_v2#group_id DataDatabricksWorkspaceIamGroupsV2#group_id}.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/workspace_iam_groups_v2#provider_config DataDatabricksWorkspaceIamGroupsV2#provider_config}.
	ProviderConfig *DataDatabricksWorkspaceIamGroupsV2GroupsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

