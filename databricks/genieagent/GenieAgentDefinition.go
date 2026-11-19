// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package genieagent


type GenieAgentDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/genie_agent#benchmarks GenieAgent#benchmarks}.
	Benchmarks *GenieAgentDefinitionBenchmarks `field:"optional" json:"benchmarks" yaml:"benchmarks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/genie_agent#config GenieAgent#config}.
	Config *GenieAgentDefinitionConfig `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/genie_agent#data_sources GenieAgent#data_sources}.
	DataSources *GenieAgentDefinitionDataSources `field:"optional" json:"dataSources" yaml:"dataSources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/genie_agent#instructions GenieAgent#instructions}.
	Instructions *GenieAgentDefinitionInstructions `field:"optional" json:"instructions" yaml:"instructions"`
}

