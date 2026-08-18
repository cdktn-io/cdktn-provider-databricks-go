// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaisearchindex


type DataDatabricksAiSearchIndexDirectAccessIndexSpecEmbeddingSourceColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_search_index#embedding_model_endpoint DataDatabricksAiSearchIndex#embedding_model_endpoint}.
	EmbeddingModelEndpoint *string `field:"optional" json:"embeddingModelEndpoint" yaml:"embeddingModelEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_search_index#model_endpoint_name_for_query DataDatabricksAiSearchIndex#model_endpoint_name_for_query}.
	ModelEndpointNameForQuery *string `field:"optional" json:"modelEndpointNameForQuery" yaml:"modelEndpointNameForQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_search_index#name DataDatabricksAiSearchIndex#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

