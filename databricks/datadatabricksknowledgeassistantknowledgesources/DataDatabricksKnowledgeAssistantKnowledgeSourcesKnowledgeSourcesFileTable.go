// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksknowledgeassistantknowledgesources


type DataDatabricksKnowledgeAssistantKnowledgeSourcesKnowledgeSourcesFileTable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/knowledge_assistant_knowledge_sources#file_col DataDatabricksKnowledgeAssistantKnowledgeSources#file_col}.
	FileCol *string `field:"required" json:"fileCol" yaml:"fileCol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/knowledge_assistant_knowledge_sources#table_name DataDatabricksKnowledgeAssistantKnowledgeSources#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

