// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksknowledgeassistants


type DataDatabricksKnowledgeAssistantsKnowledgeAssistants struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/knowledge_assistants#name DataDatabricksKnowledgeAssistants#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/knowledge_assistants#provider_config DataDatabricksKnowledgeAssistants#provider_config}.
	ProviderConfig *DataDatabricksKnowledgeAssistantsKnowledgeAssistantsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

