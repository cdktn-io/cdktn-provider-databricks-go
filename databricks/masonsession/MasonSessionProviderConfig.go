// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package masonsession


type MasonSessionProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/resources/mason_session#workspace_id MasonSession#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

