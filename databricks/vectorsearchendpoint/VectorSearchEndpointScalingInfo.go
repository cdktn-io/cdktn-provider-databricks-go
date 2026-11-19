// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vectorsearchendpoint


type VectorSearchEndpointScalingInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/vector_search_endpoint#delayed_reason VectorSearchEndpoint#delayed_reason}.
	DelayedReason *string `field:"optional" json:"delayedReason" yaml:"delayedReason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/vector_search_endpoint#estimated_completion_timestamp VectorSearchEndpoint#estimated_completion_timestamp}.
	EstimatedCompletionTimestamp *float64 `field:"optional" json:"estimatedCompletionTimestamp" yaml:"estimatedCompletionTimestamp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/vector_search_endpoint#index_initialization_percentage VectorSearchEndpoint#index_initialization_percentage}.
	IndexInitializationPercentage *float64 `field:"optional" json:"indexInitializationPercentage" yaml:"indexInitializationPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/vector_search_endpoint#phase VectorSearchEndpoint#phase}.
	Phase *string `field:"optional" json:"phase" yaml:"phase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/vector_search_endpoint#requested_target_qps VectorSearchEndpoint#requested_target_qps}.
	RequestedTargetQps *float64 `field:"optional" json:"requestedTargetQps" yaml:"requestedTargetQps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/vector_search_endpoint#state VectorSearchEndpoint#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/vector_search_endpoint#target_qps_change_allowed_after_timestamp VectorSearchEndpoint#target_qps_change_allowed_after_timestamp}.
	TargetQpsChangeAllowedAfterTimestamp *float64 `field:"optional" json:"targetQpsChangeAllowedAfterTimestamp" yaml:"targetQpsChangeAllowedAfterTimestamp"`
}

