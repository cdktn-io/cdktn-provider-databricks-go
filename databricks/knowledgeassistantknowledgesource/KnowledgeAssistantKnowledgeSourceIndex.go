// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package knowledgeassistantknowledgesource


type KnowledgeAssistantKnowledgeSourceIndex struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/knowledge_assistant_knowledge_source#doc_uri_col KnowledgeAssistantKnowledgeSource#doc_uri_col}.
	DocUriCol *string `field:"required" json:"docUriCol" yaml:"docUriCol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/knowledge_assistant_knowledge_source#index_name KnowledgeAssistantKnowledgeSource#index_name}.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/knowledge_assistant_knowledge_source#text_col KnowledgeAssistantKnowledgeSource#text_col}.
	TextCol *string `field:"required" json:"textCol" yaml:"textCol"`
}

