// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchindex


type AiSearchIndexDeltaSyncIndexSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/ai_search_index#pipeline_type AiSearchIndex#pipeline_type}.
	PipelineType *string `field:"required" json:"pipelineType" yaml:"pipelineType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/ai_search_index#columns_to_sync AiSearchIndex#columns_to_sync}.
	ColumnsToSync *[]*string `field:"optional" json:"columnsToSync" yaml:"columnsToSync"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/ai_search_index#embedding_source_columns AiSearchIndex#embedding_source_columns}.
	EmbeddingSourceColumns interface{} `field:"optional" json:"embeddingSourceColumns" yaml:"embeddingSourceColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/ai_search_index#embedding_vector_columns AiSearchIndex#embedding_vector_columns}.
	EmbeddingVectorColumns interface{} `field:"optional" json:"embeddingVectorColumns" yaml:"embeddingVectorColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/ai_search_index#embedding_writeback_table AiSearchIndex#embedding_writeback_table}.
	EmbeddingWritebackTable *string `field:"optional" json:"embeddingWritebackTable" yaml:"embeddingWritebackTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/ai_search_index#source_table AiSearchIndex#source_table}.
	SourceTable *string `field:"optional" json:"sourceTable" yaml:"sourceTable"`
}

