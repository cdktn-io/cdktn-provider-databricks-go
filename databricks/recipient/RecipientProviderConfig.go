// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package recipient


type RecipientProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/recipient#workspace_id Recipient#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

