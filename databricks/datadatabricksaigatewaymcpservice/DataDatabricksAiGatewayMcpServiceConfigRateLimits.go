// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaigatewaymcpservice


type DataDatabricksAiGatewayMcpServiceConfigRateLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_mcp_service#key DataDatabricksAiGatewayMcpService#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_mcp_service#renewal_period DataDatabricksAiGatewayMcpService#renewal_period}.
	RenewalPeriod *string `field:"required" json:"renewalPeriod" yaml:"renewalPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_mcp_service#principal DataDatabricksAiGatewayMcpService#principal}.
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_mcp_service#requests DataDatabricksAiGatewayMcpService#requests}.
	Requests *float64 `field:"optional" json:"requests" yaml:"requests"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_mcp_service#request_tag_key DataDatabricksAiGatewayMcpService#request_tag_key}.
	RequestTagKey *string `field:"optional" json:"requestTagKey" yaml:"requestTagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_mcp_service#request_tag_value DataDatabricksAiGatewayMcpService#request_tag_value}.
	RequestTagValue *string `field:"optional" json:"requestTagValue" yaml:"requestTagValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/ai_gateway_mcp_service#tokens DataDatabricksAiGatewayMcpService#tokens}.
	Tokens *float64 `field:"optional" json:"tokens" yaml:"tokens"`
}

