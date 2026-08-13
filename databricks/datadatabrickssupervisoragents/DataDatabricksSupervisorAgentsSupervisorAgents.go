// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssupervisoragents


type DataDatabricksSupervisorAgentsSupervisorAgents struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/supervisor_agents#name DataDatabricksSupervisorAgents#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/supervisor_agents#provider_config DataDatabricksSupervisorAgents#provider_config}.
	ProviderConfig *DataDatabricksSupervisorAgentsSupervisorAgentsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

