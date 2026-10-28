// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksagentkitsessionstores


type DataDatabricksAgentKitSessionStoresSessionStores struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/agent_kit_session_stores#name DataDatabricksAgentKitSessionStores#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/agent_kit_session_stores#provider_config DataDatabricksAgentKitSessionStores#provider_config}.
	ProviderConfig *DataDatabricksAgentKitSessionStoresSessionStoresProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

