// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssupervisoragenttools


type DataDatabricksSupervisorAgentToolsTools struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/supervisor_agent_tools#name DataDatabricksSupervisorAgentTools#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/supervisor_agent_tools#provider_config DataDatabricksSupervisorAgentTools#provider_config}.
	ProviderConfig *DataDatabricksSupervisorAgentToolsToolsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

