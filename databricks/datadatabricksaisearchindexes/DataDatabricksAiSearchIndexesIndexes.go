// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaisearchindexes


type DataDatabricksAiSearchIndexesIndexes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_search_indexes#name DataDatabricksAiSearchIndexes#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_search_indexes#provider_config DataDatabricksAiSearchIndexes#provider_config}.
	ProviderConfig *DataDatabricksAiSearchIndexesIndexesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

