// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservices


type DataDatabricksAiGatewayModelProviderServicesModelProviderServices struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_services#name DataDatabricksAiGatewayModelProviderServices#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_model_provider_services#provider_config DataDatabricksAiGatewayModelProviderServices#provider_config}.
	ProviderConfig *DataDatabricksAiGatewayModelProviderServicesModelProviderServicesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

