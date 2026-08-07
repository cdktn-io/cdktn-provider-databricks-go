// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeature


type DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/feature_engineering_feature#continuous DataDatabricksFeatureEngineeringFeature#continuous}.
	Continuous *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowContinuous `field:"optional" json:"continuous" yaml:"continuous"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/feature_engineering_feature#rolling DataDatabricksFeatureEngineeringFeature#rolling}.
	Rolling *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowRolling `field:"optional" json:"rolling" yaml:"rolling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/feature_engineering_feature#sawtooth DataDatabricksFeatureEngineeringFeature#sawtooth}.
	Sawtooth *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSawtooth `field:"optional" json:"sawtooth" yaml:"sawtooth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/feature_engineering_feature#sliding DataDatabricksFeatureEngineeringFeature#sliding}.
	Sliding *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSliding `field:"optional" json:"sliding" yaml:"sliding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/feature_engineering_feature#tumbling DataDatabricksFeatureEngineeringFeature#tumbling}.
	Tumbling *DataDatabricksFeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowTumbling `field:"optional" json:"tumbling" yaml:"tumbling"`
}

