// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchindex


type AiSearchIndexDirectAccessIndexSpecEmbeddingVectorColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/ai_search_index#embedding_dimension AiSearchIndex#embedding_dimension}.
	EmbeddingDimension *float64 `field:"optional" json:"embeddingDimension" yaml:"embeddingDimension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/ai_search_index#name AiSearchIndex#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

