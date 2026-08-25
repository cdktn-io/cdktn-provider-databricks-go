// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeatures


type DataDatabricksFeatureEngineeringFeaturesFeaturesFunctionCustomUdf struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_features#function_path DataDatabricksFeatureEngineeringFeatures#function_path}.
	FunctionPath *string `field:"required" json:"functionPath" yaml:"functionPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/feature_engineering_features#input_bindings DataDatabricksFeatureEngineeringFeatures#input_bindings}.
	InputBindings interface{} `field:"optional" json:"inputBindings" yaml:"inputBindings"`
}

