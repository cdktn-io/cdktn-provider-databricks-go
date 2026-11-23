// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksgenieagent


type DataDatabricksGenieAgentDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/genie_agent#benchmarks DataDatabricksGenieAgent#benchmarks}.
	Benchmarks *DataDatabricksGenieAgentDefinitionBenchmarks `field:"optional" json:"benchmarks" yaml:"benchmarks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/genie_agent#config DataDatabricksGenieAgent#config}.
	Config *DataDatabricksGenieAgentDefinitionConfig `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/genie_agent#data_sources DataDatabricksGenieAgent#data_sources}.
	DataSources *DataDatabricksGenieAgentDefinitionDataSources `field:"optional" json:"dataSources" yaml:"dataSources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/genie_agent#instructions DataDatabricksGenieAgent#instructions}.
	Instructions *DataDatabricksGenieAgentDefinitionInstructions `field:"optional" json:"instructions" yaml:"instructions"`
}

