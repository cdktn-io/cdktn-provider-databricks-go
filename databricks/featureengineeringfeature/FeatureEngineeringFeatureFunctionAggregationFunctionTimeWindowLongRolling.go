// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureFunctionAggregationFunctionTimeWindowLongRolling struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/feature_engineering_feature#window_duration FeatureEngineeringFeature#window_duration}.
	WindowDuration *string `field:"required" json:"windowDuration" yaml:"windowDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/feature_engineering_feature#delay FeatureEngineeringFeature#delay}.
	Delay *string `field:"optional" json:"delay" yaml:"delay"`
}

