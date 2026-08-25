// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureFunctionCustomUdfInputBindings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/feature_engineering_feature#column FeatureEngineeringFeature#column}.
	Column *string `field:"required" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/feature_engineering_feature#parameter FeatureEngineeringFeature#parameter}.
	Parameter *string `field:"required" json:"parameter" yaml:"parameter"`
}

