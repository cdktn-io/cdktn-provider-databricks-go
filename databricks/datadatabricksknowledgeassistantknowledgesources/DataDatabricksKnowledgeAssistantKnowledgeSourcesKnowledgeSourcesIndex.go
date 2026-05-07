// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksknowledgeassistantknowledgesources


type DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesIndex struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/knowledge_assistant_knowledge_sources#doc_uri_col DataDatabricksKnowledgeAssistantKnowledgeSources#doc_uri_col}.
	DocUriCol *string `field:"required" json:"docUriCol" yaml:"docUriCol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/knowledge_assistant_knowledge_sources#index_name DataDatabricksKnowledgeAssistantKnowledgeSources#index_name}.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/knowledge_assistant_knowledge_sources#text_col DataDatabricksKnowledgeAssistantKnowledgeSources#text_col}.
	TextCol *string `field:"required" json:"textCol" yaml:"textCol"`
}

