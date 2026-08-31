// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelproviderservice


type AiGatewayModelProviderServiceConfigAnthropic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/ai_gateway_model_provider_service#direct AiGatewayModelProviderService#direct}.
	Direct *AiGatewayModelProviderServiceConfigAnthropicDirect `field:"optional" json:"direct" yaml:"direct"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/ai_gateway_model_provider_service#relayed AiGatewayModelProviderService#relayed}.
	Relayed *AiGatewayModelProviderServiceConfigAnthropicRelayed `field:"optional" json:"relayed" yaml:"relayed"`
}

