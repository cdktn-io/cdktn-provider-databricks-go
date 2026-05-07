// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeature


type DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionApproxPercentile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/feature_engineering_feature#input DataDatabricksFeatureEngineeringFeature#input}.
	Input *string `field:"required" json:"input" yaml:"input"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/feature_engineering_feature#percentile DataDatabricksFeatureEngineeringFeature#percentile}.
	Percentile *float64 `field:"required" json:"percentile" yaml:"percentile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/feature_engineering_feature#accuracy DataDatabricksFeatureEngineeringFeature#accuracy}.
	Accuracy *float64 `field:"optional" json:"accuracy" yaml:"accuracy"`
}

