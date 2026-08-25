// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspaceiamusersv2


type DataDatabricksWorkspaceIamUsersV2Users struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/workspace_iam_users_v2#user_id DataDatabricksWorkspaceIamUsersV2#user_id}.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/workspace_iam_users_v2#provider_config DataDatabricksWorkspaceIamUsersV2#provider_config}.
	ProviderConfig *DataDatabricksWorkspaceIamUsersV2UsersProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

