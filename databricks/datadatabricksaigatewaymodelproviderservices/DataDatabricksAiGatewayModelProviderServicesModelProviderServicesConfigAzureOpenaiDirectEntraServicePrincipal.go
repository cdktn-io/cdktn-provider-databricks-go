// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservices


type DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAzureOpenaiDirectEntraServicePrincipal struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#client_id DataDatabricksAiGatewayModelProviderServices#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#client_secret DataDatabricksAiGatewayModelProviderServices#client_secret}.
	ClientSecret *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesConfigAzureOpenaiDirectEntraServicePrincipalClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/ai_gateway_model_provider_services#tenant_id DataDatabricksAiGatewayModelProviderServices#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

