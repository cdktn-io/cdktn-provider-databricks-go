// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package directory


type DirectoryProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.111.0/docs/resources/directory#workspace_id Directory#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

