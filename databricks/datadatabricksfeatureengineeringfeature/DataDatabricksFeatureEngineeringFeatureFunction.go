// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeature


type DataDatabricksFeatureEngineeringFeatureFunction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/feature_engineering_feature#aggregation_function DataDatabricksFeatureEngineeringFeature#aggregation_function}.
	AggregationFunction *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunction `field:"optional" json:"aggregationFunction" yaml:"aggregationFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/feature_engineering_feature#column_selection DataDatabricksFeatureEngineeringFeature#column_selection}.
	ColumnSelection *DataDatabricksFeatureEngineeringFeatureFunctionColumnSelection `field:"optional" json:"columnSelection" yaml:"columnSelection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/feature_engineering_feature#extra_parameters DataDatabricksFeatureEngineeringFeature#extra_parameters}.
	ExtraParameters interface{} `field:"optional" json:"extraParameters" yaml:"extraParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/feature_engineering_feature#function_type DataDatabricksFeatureEngineeringFeature#function_type}.
	FunctionType *string `field:"optional" json:"functionType" yaml:"functionType"`
}

