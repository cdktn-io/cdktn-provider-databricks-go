// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package permissionassignment


type PermissionAssignmentProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/permission_assignment#workspace_id PermissionAssignment#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

