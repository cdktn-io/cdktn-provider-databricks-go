// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspaceiamgroupsv2


type DataDatabricksWorkspaceIamGroupsV2Groups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/workspace_iam_groups_v2#name DataDatabricksWorkspaceIamGroupsV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/workspace_iam_groups_v2#provider_config DataDatabricksWorkspaceIamGroupsV2#provider_config}.
	ProviderConfig *DataDatabricksWorkspaceIamGroupsV2GroupsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

