// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connection


type ConnectionProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/connection#workspace_id Connection#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

