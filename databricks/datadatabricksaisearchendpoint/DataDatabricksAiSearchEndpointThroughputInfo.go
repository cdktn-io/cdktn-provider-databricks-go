// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaisearchendpoint


type DataDatabricksAiSearchEndpointThroughputInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/ai_search_endpoint#maximum_concurrency_allowed DataDatabricksAiSearchEndpoint#maximum_concurrency_allowed}.
	MaximumConcurrencyAllowed *float64 `field:"optional" json:"maximumConcurrencyAllowed" yaml:"maximumConcurrencyAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/ai_search_endpoint#minimal_concurrency_allowed DataDatabricksAiSearchEndpoint#minimal_concurrency_allowed}.
	MinimalConcurrencyAllowed *float64 `field:"optional" json:"minimalConcurrencyAllowed" yaml:"minimalConcurrencyAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/ai_search_endpoint#requested_concurrency DataDatabricksAiSearchEndpoint#requested_concurrency}.
	RequestedConcurrency *float64 `field:"optional" json:"requestedConcurrency" yaml:"requestedConcurrency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/ai_search_endpoint#requested_num_replicas DataDatabricksAiSearchEndpoint#requested_num_replicas}.
	RequestedNumReplicas *float64 `field:"optional" json:"requestedNumReplicas" yaml:"requestedNumReplicas"`
}

