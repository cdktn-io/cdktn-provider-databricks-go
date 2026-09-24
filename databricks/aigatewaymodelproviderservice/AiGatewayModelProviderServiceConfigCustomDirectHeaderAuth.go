// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelproviderservice


type AiGatewayModelProviderServiceConfigCustomDirectHeaderAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/ai_gateway_model_provider_service#api_key_name AiGatewayModelProviderService#api_key_name}.
	ApiKeyName *string `field:"optional" json:"apiKeyName" yaml:"apiKeyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/ai_gateway_model_provider_service#api_key_value AiGatewayModelProviderService#api_key_value}.
	ApiKeyValue *AiGatewayModelProviderServiceConfigCustomDirectHeaderAuthApiKeyValue `field:"optional" json:"apiKeyValue" yaml:"apiKeyValue"`
}

