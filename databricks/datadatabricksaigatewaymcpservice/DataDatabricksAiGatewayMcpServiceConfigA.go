// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymcpservice


type DataDatabricksAiGatewayMcpServiceConfigA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/ai_gateway_mcp_service#include_tool_selectors DataDatabricksAiGatewayMcpService#include_tool_selectors}.
	IncludeToolSelectors *[]*string `field:"optional" json:"includeToolSelectors" yaml:"includeToolSelectors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/ai_gateway_mcp_service#rate_limits DataDatabricksAiGatewayMcpService#rate_limits}.
	RateLimits interface{} `field:"optional" json:"rateLimits" yaml:"rateLimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/ai_gateway_mcp_service#source_connection DataDatabricksAiGatewayMcpService#source_connection}.
	SourceConnection *DataDatabricksAiGatewayMcpServiceConfigSourceConnection `field:"optional" json:"sourceConnection" yaml:"sourceConnection"`
}

