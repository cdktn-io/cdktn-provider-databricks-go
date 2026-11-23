// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksgenieagent


type DataDatabricksGenieAgentDefinitionDataSourcesTables struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/genie_agent#identifier DataDatabricksGenieAgent#identifier}.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/genie_agent#column_configs DataDatabricksGenieAgent#column_configs}.
	ColumnConfigs interface{} `field:"optional" json:"columnConfigs" yaml:"columnConfigs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/genie_agent#description DataDatabricksGenieAgent#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

