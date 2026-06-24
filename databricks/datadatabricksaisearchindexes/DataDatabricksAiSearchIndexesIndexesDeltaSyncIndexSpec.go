// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaisearchindexes


type DataDatabricksAiSearchIndexesIndexesDeltaSyncIndexSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/ai_search_indexes#pipeline_type DataDatabricksAiSearchIndexes#pipeline_type}.
	PipelineType *string `field:"required" json:"pipelineType" yaml:"pipelineType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/ai_search_indexes#columns_to_sync DataDatabricksAiSearchIndexes#columns_to_sync}.
	ColumnsToSync *[]*string `field:"optional" json:"columnsToSync" yaml:"columnsToSync"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/ai_search_indexes#embedding_source_columns DataDatabricksAiSearchIndexes#embedding_source_columns}.
	EmbeddingSourceColumns interface{} `field:"optional" json:"embeddingSourceColumns" yaml:"embeddingSourceColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/ai_search_indexes#embedding_vector_columns DataDatabricksAiSearchIndexes#embedding_vector_columns}.
	EmbeddingVectorColumns interface{} `field:"optional" json:"embeddingVectorColumns" yaml:"embeddingVectorColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/ai_search_indexes#embedding_writeback_table DataDatabricksAiSearchIndexes#embedding_writeback_table}.
	EmbeddingWritebackTable *string `field:"optional" json:"embeddingWritebackTable" yaml:"embeddingWritebackTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/ai_search_indexes#source_table DataDatabricksAiSearchIndexes#source_table}.
	SourceTable *string `field:"optional" json:"sourceTable" yaml:"sourceTable"`
}

