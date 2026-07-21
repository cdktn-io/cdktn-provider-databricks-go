// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package instanceprofile


type InstanceProfileProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/instance_profile#workspace_id InstanceProfile#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

