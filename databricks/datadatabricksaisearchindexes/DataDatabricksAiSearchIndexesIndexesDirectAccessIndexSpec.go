// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaisearchindexes


type DataDatabricksAiSearchIndexesIndexesDirectAccessIndexSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/ai_search_indexes#embedding_source_columns DataDatabricksAiSearchIndexes#embedding_source_columns}.
	EmbeddingSourceColumns interface{} `field:"optional" json:"embeddingSourceColumns" yaml:"embeddingSourceColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/ai_search_indexes#embedding_vector_columns DataDatabricksAiSearchIndexes#embedding_vector_columns}.
	EmbeddingVectorColumns interface{} `field:"optional" json:"embeddingVectorColumns" yaml:"embeddingVectorColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/ai_search_indexes#schema_json DataDatabricksAiSearchIndexes#schema_json}.
	SchemaJson *string `field:"optional" json:"schemaJson" yaml:"schemaJson"`
}

