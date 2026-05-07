// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy


type DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessAllowRulesDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/account_network_policy#all_destinations DataDatabricksAccountNetworkPolicy#all_destinations}.
	AllDestinations interface{} `field:"optional" json:"allDestinations" yaml:"allDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/account_network_policy#workspace_api DataDatabricksAccountNetworkPolicy#workspace_api}.
	WorkspaceApi *DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessAllowRulesDestinationWorkspaceApi `field:"optional" json:"workspaceApi" yaml:"workspaceApi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/account_network_policy#workspace_ui DataDatabricksAccountNetworkPolicy#workspace_ui}.
	WorkspaceUi *DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessAllowRulesDestinationWorkspaceUi `field:"optional" json:"workspaceUi" yaml:"workspaceUi"`
}

