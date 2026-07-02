// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchindex


type AiSearchIndexDirectAccessIndexSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/ai_search_index#embedding_source_columns AiSearchIndex#embedding_source_columns}.
	EmbeddingSourceColumns interface{} `field:"optional" json:"embeddingSourceColumns" yaml:"embeddingSourceColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/ai_search_index#embedding_vector_columns AiSearchIndex#embedding_vector_columns}.
	EmbeddingVectorColumns interface{} `field:"optional" json:"embeddingVectorColumns" yaml:"embeddingVectorColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/ai_search_index#schema_json AiSearchIndex#schema_json}.
	SchemaJson *string `field:"optional" json:"schemaJson" yaml:"schemaJson"`
}

