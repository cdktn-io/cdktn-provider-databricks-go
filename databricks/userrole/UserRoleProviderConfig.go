// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package userrole


type UserRoleProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/user_role#workspace_id UserRole#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

