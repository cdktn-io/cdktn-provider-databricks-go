// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelproviderservice


type AiGatewayModelProviderServiceConfigTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/ai_gateway_model_provider_service#model AiGatewayModelProviderService#model}.
	Model *string `field:"required" json:"model" yaml:"model"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/ai_gateway_model_provider_service#native_api_types AiGatewayModelProviderService#native_api_types}.
	NativeApiTypes *[]*string `field:"optional" json:"nativeApiTypes" yaml:"nativeApiTypes"`
}

