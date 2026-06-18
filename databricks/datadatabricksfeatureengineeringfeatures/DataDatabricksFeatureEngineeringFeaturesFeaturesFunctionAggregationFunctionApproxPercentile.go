// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeatures


type DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionApproxPercentile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/feature_engineering_features#input DataDatabricksFeatureEngineeringFeatures#input}.
	Input *string `field:"required" json:"input" yaml:"input"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/feature_engineering_features#percentile DataDatabricksFeatureEngineeringFeatures#percentile}.
	Percentile *float64 `field:"required" json:"percentile" yaml:"percentile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/feature_engineering_features#accuracy DataDatabricksFeatureEngineeringFeatures#accuracy}.
	Accuracy *float64 `field:"optional" json:"accuracy" yaml:"accuracy"`
}

