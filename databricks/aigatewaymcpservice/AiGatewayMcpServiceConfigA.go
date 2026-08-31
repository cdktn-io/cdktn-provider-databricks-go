// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymcpservice


type AiGatewayMcpServiceConfigA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/ai_gateway_mcp_service#include_tool_selectors AiGatewayMcpService#include_tool_selectors}.
	IncludeToolSelectors *[]*string `field:"optional" json:"includeToolSelectors" yaml:"includeToolSelectors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/ai_gateway_mcp_service#rate_limits AiGatewayMcpService#rate_limits}.
	RateLimits interface{} `field:"optional" json:"rateLimits" yaml:"rateLimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/ai_gateway_mcp_service#source_connection AiGatewayMcpService#source_connection}.
	SourceConnection *AiGatewayMcpServiceConfigSourceConnection `field:"optional" json:"sourceConnection" yaml:"sourceConnection"`
}

