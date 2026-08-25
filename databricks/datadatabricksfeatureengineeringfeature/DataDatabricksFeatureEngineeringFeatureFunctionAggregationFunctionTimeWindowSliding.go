// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeature


type DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSliding struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#slide_duration DataDatabricksFeatureEngineeringFeature#slide_duration}.
	SlideDuration *string `field:"required" json:"slideDuration" yaml:"slideDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_feature#window_duration DataDatabricksFeatureEngineeringFeature#window_duration}.
	WindowDuration *string `field:"optional" json:"windowDuration" yaml:"windowDuration"`
}

