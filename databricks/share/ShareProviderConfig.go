// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package share


type ShareProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/share#workspace_id Share#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

