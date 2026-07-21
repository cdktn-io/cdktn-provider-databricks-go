// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mount


type MountProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/mount#workspace_id Mount#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

