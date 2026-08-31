// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package permissions


type PermissionsProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/permissions#workspace_id Permissions#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

