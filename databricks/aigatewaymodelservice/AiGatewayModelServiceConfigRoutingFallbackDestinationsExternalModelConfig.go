// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelservice


type AiGatewayModelServiceConfigRoutingFallbackDestinationsExternalModelConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/ai_gateway_model_service#model_provider_service AiGatewayModelService#model_provider_service}.
	ModelProviderService *string `field:"required" json:"modelProviderService" yaml:"modelProviderService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/ai_gateway_model_service#target AiGatewayModelService#target}.
	Target *AiGatewayModelServiceConfigRoutingFallbackDestinationsExternalModelConfigTarget `field:"required" json:"target" yaml:"target"`
}

