// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaisearchendpoints


type DataDatabricksAiSearchEndpointsEndpoints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_search_endpoints#name DataDatabricksAiSearchEndpoints#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_search_endpoints#provider_config DataDatabricksAiSearchEndpoints#provider_config}.
	ProviderConfig *DataDatabricksAiSearchEndpointsEndpointsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

