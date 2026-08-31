// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeatures


type DataDatabricksFeatureEngineeringFeaturesFeaturesFunction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/feature_engineering_features#aggregation_function DataDatabricksFeatureEngineeringFeatures#aggregation_function}.
	AggregationFunction *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunction `field:"optional" json:"aggregationFunction" yaml:"aggregationFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/feature_engineering_features#column_selection DataDatabricksFeatureEngineeringFeatures#column_selection}.
	ColumnSelection *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionColumnSelection `field:"optional" json:"columnSelection" yaml:"columnSelection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/feature_engineering_features#custom_udf DataDatabricksFeatureEngineeringFeatures#custom_udf}.
	CustomUdf *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionCustomUdf `field:"optional" json:"customUdf" yaml:"customUdf"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/feature_engineering_features#extra_parameters DataDatabricksFeatureEngineeringFeatures#extra_parameters}.
	ExtraParameters interface{} `field:"optional" json:"extraParameters" yaml:"extraParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/feature_engineering_features#function_type DataDatabricksFeatureEngineeringFeatures#function_type}.
	FunctionType *string `field:"optional" json:"functionType" yaml:"functionType"`
}

