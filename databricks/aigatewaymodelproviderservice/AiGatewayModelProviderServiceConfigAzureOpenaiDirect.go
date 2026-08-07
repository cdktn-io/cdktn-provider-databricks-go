// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelproviderservice


type AiGatewayModelProviderServiceConfigAzureOpenaiDirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/ai_gateway_model_provider_service#api_key AiGatewayModelProviderService#api_key}.
	ApiKey *AiGatewayModelProviderServiceConfigAzureOpenaiDirectApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/ai_gateway_model_provider_service#base_url AiGatewayModelProviderService#base_url}.
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/ai_gateway_model_provider_service#client_id AiGatewayModelProviderService#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/ai_gateway_model_provider_service#client_secret AiGatewayModelProviderService#client_secret}.
	ClientSecret *AiGatewayModelProviderServiceConfigAzureOpenaiDirectClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/ai_gateway_model_provider_service#service_credential AiGatewayModelProviderService#service_credential}.
	ServiceCredential *AiGatewayModelProviderServiceConfigAzureOpenaiDirectServiceCredential `field:"optional" json:"serviceCredential" yaml:"serviceCredential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/ai_gateway_model_provider_service#tenant_id AiGatewayModelProviderService#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

