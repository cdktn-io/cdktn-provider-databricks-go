// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservice


type DataDatabricksAiGatewayModelProviderServiceConfigMicrosoftFoundryDirectEntraServicePrincipal struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_provider_service#client_id DataDatabricksAiGatewayModelProviderService#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_provider_service#client_secret DataDatabricksAiGatewayModelProviderService#client_secret}.
	ClientSecret *DataDatabricksAiGatewayModelProviderServiceConfigMicrosoftFoundryDirectEntraServicePrincipalClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_provider_service#tenant_id DataDatabricksAiGatewayModelProviderService#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

