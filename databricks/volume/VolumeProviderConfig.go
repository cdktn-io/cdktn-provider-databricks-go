// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package volume


type VolumeProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/volume#workspace_id Volume#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

