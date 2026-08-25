// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelproviderservice


type AiGatewayModelProviderServiceConfigCustomDirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/ai_gateway_model_provider_service#api_key AiGatewayModelProviderService#api_key}.
	ApiKey *AiGatewayModelProviderServiceConfigCustomDirectApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/ai_gateway_model_provider_service#base_url AiGatewayModelProviderService#base_url}.
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
}

