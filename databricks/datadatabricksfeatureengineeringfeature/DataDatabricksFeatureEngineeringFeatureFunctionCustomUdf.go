// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeature


type DataDatabricksFeatureEngineeringFeatureFunctionCustomUdf struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/feature_engineering_feature#function_path DataDatabricksFeatureEngineeringFeature#function_path}.
	FunctionPath *string `field:"required" json:"functionPath" yaml:"functionPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/feature_engineering_feature#input_bindings DataDatabricksFeatureEngineeringFeature#input_bindings}.
	InputBindings interface{} `field:"optional" json:"inputBindings" yaml:"inputBindings"`
}

