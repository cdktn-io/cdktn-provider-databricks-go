// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservices


type DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_services#api_key DataDatabricksAiGatewayModelProviderServices#api_key}.
	ApiKey *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_services#base_url DataDatabricksAiGatewayModelProviderServices#base_url}.
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_services#client_id DataDatabricksAiGatewayModelProviderServices#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_services#client_secret DataDatabricksAiGatewayModelProviderServices#client_secret}.
	ClientSecret *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_services#service_credential DataDatabricksAiGatewayModelProviderServices#service_credential}.
	ServiceCredential *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectServiceCredential `field:"optional" json:"serviceCredential" yaml:"serviceCredential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_services#tenant_id DataDatabricksAiGatewayModelProviderServices#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

