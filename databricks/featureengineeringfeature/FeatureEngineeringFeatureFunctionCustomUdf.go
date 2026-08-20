// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureFunctionCustomUdf struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/feature_engineering_feature#function_path FeatureEngineeringFeature#function_path}.
	FunctionPath *string `field:"required" json:"functionPath" yaml:"functionPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/feature_engineering_feature#input_bindings FeatureEngineeringFeature#input_bindings}.
	InputBindings interface{} `field:"optional" json:"inputBindings" yaml:"inputBindings"`
}

