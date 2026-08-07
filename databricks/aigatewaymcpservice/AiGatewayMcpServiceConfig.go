// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymcpservice

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiGatewayMcpServiceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/ai_gateway_mcp_service#mcp_service_id AiGatewayMcpService#mcp_service_id}.
	McpServiceId *string `field:"required" json:"mcpServiceId" yaml:"mcpServiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/ai_gateway_mcp_service#parent AiGatewayMcpService#parent}.
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/ai_gateway_mcp_service#comment AiGatewayMcpService#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/ai_gateway_mcp_service#config AiGatewayMcpService#config}.
	Config *AiGatewayMcpServiceConfigA `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/ai_gateway_mcp_service#owner AiGatewayMcpService#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/ai_gateway_mcp_service#provider_config AiGatewayMcpService#provider_config}.
	ProviderConfig *AiGatewayMcpServiceProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

