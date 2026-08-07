// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymcpservices


type DataDatabricksAiGatewayMcpServicesMcpServicesConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_mcp_services#include_tool_selectors DataDatabricksAiGatewayMcpServices#include_tool_selectors}.
	IncludeToolSelectors *[]*string `field:"optional" json:"includeToolSelectors" yaml:"includeToolSelectors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_mcp_services#rate_limits DataDatabricksAiGatewayMcpServices#rate_limits}.
	RateLimits interface{} `field:"optional" json:"rateLimits" yaml:"rateLimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_mcp_services#source_connection DataDatabricksAiGatewayMcpServices#source_connection}.
	SourceConnection *DataDatabricksAiGatewayMcpServicesMcpServicesConfigSourceConnection `field:"optional" json:"sourceConnection" yaml:"sourceConnection"`
}

