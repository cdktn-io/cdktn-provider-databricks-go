// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservice


type DataDatabricksAiGatewayModelProviderServiceConfigAzureOpenaiDirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/ai_gateway_model_provider_service#api_key DataDatabricksAiGatewayModelProviderService#api_key}.
	ApiKey *DataDatabricksAiGatewayModelProviderServiceConfigAzureOpenaiDirectApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/ai_gateway_model_provider_service#base_url DataDatabricksAiGatewayModelProviderService#base_url}.
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/ai_gateway_model_provider_service#entra_service_principal DataDatabricksAiGatewayModelProviderService#entra_service_principal}.
	EntraServicePrincipal *DataDatabricksAiGatewayModelProviderServiceConfigAzureOpenaiDirectEntraServicePrincipal `field:"optional" json:"entraServicePrincipal" yaml:"entraServicePrincipal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/ai_gateway_model_provider_service#service_credential DataDatabricksAiGatewayModelProviderService#service_credential}.
	ServiceCredential *DataDatabricksAiGatewayModelProviderServiceConfigAzureOpenaiDirectServiceCredential `field:"optional" json:"serviceCredential" yaml:"serviceCredential"`
}

