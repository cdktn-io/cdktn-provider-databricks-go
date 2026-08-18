// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeatures


type DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/feature_engineering_features#continuous DataDatabricksFeatureEngineeringFeatures#continuous}.
	Continuous *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowContinuous `field:"optional" json:"continuous" yaml:"continuous"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/feature_engineering_features#rolling DataDatabricksFeatureEngineeringFeatures#rolling}.
	Rolling *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowRolling `field:"optional" json:"rolling" yaml:"rolling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/feature_engineering_features#sawtooth DataDatabricksFeatureEngineeringFeatures#sawtooth}.
	Sawtooth *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowSawtooth `field:"optional" json:"sawtooth" yaml:"sawtooth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/feature_engineering_features#sliding DataDatabricksFeatureEngineeringFeatures#sliding}.
	Sliding *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowSliding `field:"optional" json:"sliding" yaml:"sliding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/feature_engineering_features#tumbling DataDatabricksFeatureEngineeringFeatures#tumbling}.
	Tumbling *DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionAggregationFunctionTimeWindowTumbling `field:"optional" json:"tumbling" yaml:"tumbling"`
}

