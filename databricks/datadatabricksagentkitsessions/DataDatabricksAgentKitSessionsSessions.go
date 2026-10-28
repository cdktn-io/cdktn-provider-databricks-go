// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksagentkitsessions


type DataDatabricksAgentKitSessionsSessions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/agent_kit_sessions#name DataDatabricksAgentKitSessions#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/agent_kit_sessions#provider_config DataDatabricksAgentKitSessions#provider_config}.
	ProviderConfig *DataDatabricksAgentKitSessionsSessionsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

