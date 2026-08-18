// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymodelproviderservice


type DataDatabricksAiGatewayModelProviderServiceConfigAnthropic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_provider_service#direct DataDatabricksAiGatewayModelProviderService#direct}.
	Direct *DataDatabricksAiGatewayModelProviderServiceConfigAnthropicDirect `field:"optional" json:"direct" yaml:"direct"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/ai_gateway_model_provider_service#relayed DataDatabricksAiGatewayModelProviderService#relayed}.
	Relayed *DataDatabricksAiGatewayModelProviderServiceConfigAnthropicRelayed `field:"optional" json:"relayed" yaml:"relayed"`
}

