// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dbfsfile


type DbfsFileProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/dbfs_file#workspace_id DbfsFile#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

