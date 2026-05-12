// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grouprole


type GroupRoleProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/group_role#workspace_id GroupRole#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

