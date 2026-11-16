// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package genieagent


type GenieAgentDefinitionDataSources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/genie_agent#mcp_connections GenieAgent#mcp_connections}.
	McpConnections interface{} `field:"optional" json:"mcpConnections" yaml:"mcpConnections"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/genie_agent#tables GenieAgent#tables}.
	Tables interface{} `field:"optional" json:"tables" yaml:"tables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/genie_agent#volumes GenieAgent#volumes}.
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

