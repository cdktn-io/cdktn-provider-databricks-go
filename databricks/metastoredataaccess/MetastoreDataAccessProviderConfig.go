// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package metastoredataaccess


type MetastoreDataAccessProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/metastore_data_access#workspace_id MetastoreDataAccess#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

