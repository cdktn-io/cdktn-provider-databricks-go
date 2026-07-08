// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aisearchendpoint


type AiSearchEndpointScalingInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/ai_search_endpoint#requested_target_qps AiSearchEndpoint#requested_target_qps}.
	RequestedTargetQps *float64 `field:"optional" json:"requestedTargetQps" yaml:"requestedTargetQps"`
}

