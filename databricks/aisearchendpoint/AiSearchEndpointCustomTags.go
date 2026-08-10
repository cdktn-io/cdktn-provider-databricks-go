// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchendpoint


type AiSearchEndpointCustomTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/ai_search_endpoint#key AiSearchEndpoint#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/ai_search_endpoint#value AiSearchEndpoint#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

