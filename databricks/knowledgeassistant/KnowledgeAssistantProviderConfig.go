// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package knowledgeassistant


type KnowledgeAssistantProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/knowledge_assistant#workspace_id KnowledgeAssistant#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

