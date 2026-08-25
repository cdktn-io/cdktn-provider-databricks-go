// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksknowledgeassistantknowledgesources


type DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/knowledge_assistant_knowledge_sources#name DataDatabricksKnowledgeAssistantKnowledgeSources#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/knowledge_assistant_knowledge_sources#provider_config DataDatabricksKnowledgeAssistantKnowledgeSources#provider_config}.
	ProviderConfig *DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

