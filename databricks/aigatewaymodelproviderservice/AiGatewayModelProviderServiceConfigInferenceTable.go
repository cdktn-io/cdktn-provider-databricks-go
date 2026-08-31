// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelproviderservice


type AiGatewayModelProviderServiceConfigInferenceTable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/ai_gateway_model_provider_service#parent AiGatewayModelProviderService#parent}.
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/ai_gateway_model_provider_service#disabled AiGatewayModelProviderService#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/ai_gateway_model_provider_service#table_name_prefix AiGatewayModelProviderService#table_name_prefix}.
	TableNamePrefix *string `field:"optional" json:"tableNamePrefix" yaml:"tableNamePrefix"`
}

