// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package userinstanceprofile


type UserInstanceProfileProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/user_instance_profile#workspace_id UserInstanceProfile#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

