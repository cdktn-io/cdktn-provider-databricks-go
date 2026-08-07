// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservice


type DataDatabricksAiGatewayModelProviderServiceConfigAzureOpenaiDirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_service#api_key DataDatabricksAiGatewayModelProviderService#api_key}.
	ApiKey *DataDatabricksAiGatewayModelProviderServiceConfigAzureOpenaiDirectApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_service#base_url DataDatabricksAiGatewayModelProviderService#base_url}.
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_service#client_id DataDatabricksAiGatewayModelProviderService#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_service#client_secret DataDatabricksAiGatewayModelProviderService#client_secret}.
	ClientSecret *DataDatabricksAiGatewayModelProviderServiceConfigAzureOpenaiDirectClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_service#service_credential DataDatabricksAiGatewayModelProviderService#service_credential}.
	ServiceCredential *DataDatabricksAiGatewayModelProviderServiceConfigAzureOpenaiDirectServiceCredential `field:"optional" json:"serviceCredential" yaml:"serviceCredential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_service#tenant_id DataDatabricksAiGatewayModelProviderService#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

