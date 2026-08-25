// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databaseinstance


type DatabaseInstanceProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/database_instance#workspace_id DatabaseInstance#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

