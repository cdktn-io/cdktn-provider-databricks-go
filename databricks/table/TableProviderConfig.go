// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package table


type TableProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/table#workspace_id Table#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

