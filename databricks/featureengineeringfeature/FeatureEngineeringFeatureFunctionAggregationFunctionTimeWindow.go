// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/feature_engineering_feature#continuous FeatureEngineeringFeature#continuous}.
	Continuous *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowContinuous `field:"optional" json:"continuous" yaml:"continuous"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/feature_engineering_feature#rolling FeatureEngineeringFeature#rolling}.
	Rolling *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowRolling `field:"optional" json:"rolling" yaml:"rolling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/feature_engineering_feature#sliding FeatureEngineeringFeature#sliding}.
	Sliding *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowSliding `field:"optional" json:"sliding" yaml:"sliding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/feature_engineering_feature#tumbling FeatureEngineeringFeature#tumbling}.
	Tumbling *FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowTumbling `field:"optional" json:"tumbling" yaml:"tumbling"`
}

