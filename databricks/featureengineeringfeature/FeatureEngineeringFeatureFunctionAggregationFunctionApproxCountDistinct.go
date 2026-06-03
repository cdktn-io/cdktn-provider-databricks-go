// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureFunctionAggregationFunctionApproxCountDistinct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/feature_engineering_feature#input FeatureEngineeringFeature#input}.
	Input *string `field:"required" json:"input" yaml:"input"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/feature_engineering_feature#relative_sd FeatureEngineeringFeature#relative_sd}.
	RelativeSd *float64 `field:"optional" json:"relativeSd" yaml:"relativeSd"`
}

