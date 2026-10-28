// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelservice


type AiGatewayModelServiceConfigRoutingFallbackDestinationsExternalModelConfigTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/ai_gateway_model_service#model AiGatewayModelService#model}.
	Model *string `field:"required" json:"model" yaml:"model"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/ai_gateway_model_service#endpoint_route AiGatewayModelService#endpoint_route}.
	EndpointRoute *string `field:"optional" json:"endpointRoute" yaml:"endpointRoute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/ai_gateway_model_service#native_api_types AiGatewayModelService#native_api_types}.
	NativeApiTypes *[]*string `field:"optional" json:"nativeApiTypes" yaml:"nativeApiTypes"`
}

