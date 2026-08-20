// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservices


type DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#api_key DataDatabricksAiGatewayModelProviderServices#api_key}.
	ApiKey *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#base_url DataDatabricksAiGatewayModelProviderServices#base_url}.
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#entra_service_principal DataDatabricksAiGatewayModelProviderServices#entra_service_principal}.
	EntraServicePrincipal *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectEntraServicePrincipal `field:"optional" json:"entraServicePrincipal" yaml:"entraServicePrincipal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#service_credential DataDatabricksAiGatewayModelProviderServices#service_credential}.
	ServiceCredential *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigMicrosoftFoundryDirectServiceCredential `field:"optional" json:"serviceCredential" yaml:"serviceCredential"`
}

