// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureFunction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/feature_engineering_feature#aggregation_function FeatureEngineeringFeature#aggregation_function}.
	AggregationFunction *FeatureEngineeringFeatureFunctionAggregationFunction `field:"optional" json:"aggregationFunction" yaml:"aggregationFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/feature_engineering_feature#column_selection FeatureEngineeringFeature#column_selection}.
	ColumnSelection *FeatureEngineeringFeatureFunctionColumnSelection `field:"optional" json:"columnSelection" yaml:"columnSelection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/feature_engineering_feature#extra_parameters FeatureEngineeringFeature#extra_parameters}.
	ExtraParameters interface{} `field:"optional" json:"extraParameters" yaml:"extraParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/feature_engineering_feature#function_type FeatureEngineeringFeature#function_type}.
	FunctionType *string `field:"optional" json:"functionType" yaml:"functionType"`
}

