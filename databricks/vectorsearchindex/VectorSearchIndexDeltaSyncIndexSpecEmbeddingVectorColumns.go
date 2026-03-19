// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vectorsearchindex


type VectorSearchIndexDeltaSyncIndexSpecEmbeddingVectorColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/vector_search_index#embedding_dimension VectorSearchIndex#embedding_dimension}.
	EmbeddingDimension *float64 `field:"optional" json:"embeddingDimension" yaml:"embeddingDimension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/vector_search_index#name VectorSearchIndex#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

