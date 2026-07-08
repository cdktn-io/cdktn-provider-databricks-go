// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchindex


type AiSearchIndexProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/ai_search_index#workspace_id AiSearchIndex#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

