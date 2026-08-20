// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelservice


type DataDatabricksAiGatewayModelServiceConfigRoutingFallbackDestinationsExternalModelConfigTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_service#model DataDatabricksAiGatewayModelService#model}.
	Model *string `field:"required" json:"model" yaml:"model"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_service#native_api_types DataDatabricksAiGatewayModelService#native_api_types}.
	NativeApiTypes *[]*string `field:"optional" json:"nativeApiTypes" yaml:"nativeApiTypes"`
}

