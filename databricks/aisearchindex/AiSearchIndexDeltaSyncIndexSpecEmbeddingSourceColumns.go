// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchindex


type AiSearchIndexDeltaSyncIndexSpecEmbeddingSourceColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/ai_search_index#embedding_model_endpoint AiSearchIndex#embedding_model_endpoint}.
	EmbeddingModelEndpoint *string `field:"optional" json:"embeddingModelEndpoint" yaml:"embeddingModelEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/ai_search_index#model_endpoint_name_for_query AiSearchIndex#model_endpoint_name_for_query}.
	ModelEndpointNameForQuery *string `field:"optional" json:"modelEndpointNameForQuery" yaml:"modelEndpointNameForQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/ai_search_index#name AiSearchIndex#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

