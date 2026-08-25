// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelservices


type DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingDestinationsExternalModelConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/ai_gateway_model_services#model_provider_service DataDatabricksAiGatewayModelServices#model_provider_service}.
	ModelProviderService *string `field:"required" json:"modelProviderService" yaml:"modelProviderService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/ai_gateway_model_services#target DataDatabricksAiGatewayModelServices#target}.
	Target *DataDatabricksAiGatewayModelServicesModelServicesConfigRoutingDestinationsExternalModelConfigTarget `field:"required" json:"target" yaml:"target"`
}

