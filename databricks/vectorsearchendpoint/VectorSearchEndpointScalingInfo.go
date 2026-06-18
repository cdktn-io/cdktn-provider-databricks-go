// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vectorsearchendpoint


type VectorSearchEndpointScalingInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/vector_search_endpoint#requested_target_qps VectorSearchEndpoint#requested_target_qps}.
	RequestedTargetQps *float64 `field:"optional" json:"requestedTargetQps" yaml:"requestedTargetQps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/vector_search_endpoint#state VectorSearchEndpoint#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
}

