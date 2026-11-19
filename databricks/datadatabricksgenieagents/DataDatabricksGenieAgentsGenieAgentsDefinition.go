// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksgenieagents


type DataDatabricksGenieAgentsGenieAgentsDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/genie_agents#benchmarks DataDatabricksGenieAgents#benchmarks}.
	Benchmarks *DataDatabricksGenieAgentsGenieAgentsDefinitionBenchmarks `field:"optional" json:"benchmarks" yaml:"benchmarks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/genie_agents#config DataDatabricksGenieAgents#config}.
	Config *DataDatabricksGenieAgentsGenieAgentsDefinitionConfig `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/genie_agents#data_sources DataDatabricksGenieAgents#data_sources}.
	DataSources *DataDatabricksGenieAgentsGenieAgentsDefinitionDataSources `field:"optional" json:"dataSources" yaml:"dataSources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/genie_agents#instructions DataDatabricksGenieAgents#instructions}.
	Instructions *DataDatabricksGenieAgentsGenieAgentsDefinitionInstructions `field:"optional" json:"instructions" yaml:"instructions"`
}

