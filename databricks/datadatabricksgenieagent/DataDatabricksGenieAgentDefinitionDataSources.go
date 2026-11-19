// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksgenieagent


type DataDatabricksGenieAgentDefinitionDataSources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/genie_agent#mcp_connections DataDatabricksGenieAgent#mcp_connections}.
	McpConnections interface{} `field:"optional" json:"mcpConnections" yaml:"mcpConnections"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/genie_agent#tables DataDatabricksGenieAgent#tables}.
	Tables interface{} `field:"optional" json:"tables" yaml:"tables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/genie_agent#volumes DataDatabricksGenieAgent#volumes}.
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

