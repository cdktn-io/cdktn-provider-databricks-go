// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package query


type QueryProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/query#workspace_id Query#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

