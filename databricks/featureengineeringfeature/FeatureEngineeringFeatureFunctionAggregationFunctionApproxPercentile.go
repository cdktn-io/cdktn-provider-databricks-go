// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureFunctionAggregationFunctionApproxPercentile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#input FeatureEngineeringFeature#input}.
	Input *string `field:"required" json:"input" yaml:"input"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#percentile FeatureEngineeringFeature#percentile}.
	Percentile *float64 `field:"required" json:"percentile" yaml:"percentile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_feature#accuracy FeatureEngineeringFeature#accuracy}.
	Accuracy *float64 `field:"optional" json:"accuracy" yaml:"accuracy"`
}

