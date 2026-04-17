// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groupmember


type GroupMemberProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/group_member#workspace_id GroupMember#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

