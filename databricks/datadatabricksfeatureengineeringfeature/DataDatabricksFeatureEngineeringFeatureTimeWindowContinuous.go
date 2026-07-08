// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeature


type DataDatabricksFeatureEngineeringFeatureTimeWindowContinuous struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/feature_engineering_feature#window_duration DataDatabricksFeatureEngineeringFeature#window_duration}.
	WindowDuration *string `field:"required" json:"windowDuration" yaml:"windowDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/feature_engineering_feature#offset DataDatabricksFeatureEngineeringFeature#offset}.
	Offset *string `field:"optional" json:"offset" yaml:"offset"`
}

