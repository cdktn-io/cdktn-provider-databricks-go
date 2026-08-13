// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelproviderservice


type AiGatewayModelProviderServiceConfigMicrosoftFoundryDirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/ai_gateway_model_provider_service#api_key AiGatewayModelProviderService#api_key}.
	ApiKey *AiGatewayModelProviderServiceConfigMicrosoftFoundryDirectApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/ai_gateway_model_provider_service#base_url AiGatewayModelProviderService#base_url}.
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/ai_gateway_model_provider_service#entra_service_principal AiGatewayModelProviderService#entra_service_principal}.
	EntraServicePrincipal *AiGatewayModelProviderServiceConfigMicrosoftFoundryDirectEntraServicePrincipal `field:"optional" json:"entraServicePrincipal" yaml:"entraServicePrincipal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/ai_gateway_model_provider_service#service_credential AiGatewayModelProviderService#service_credential}.
	ServiceCredential *AiGatewayModelProviderServiceConfigMicrosoftFoundryDirectServiceCredential `field:"optional" json:"serviceCredential" yaml:"serviceCredential"`
}

