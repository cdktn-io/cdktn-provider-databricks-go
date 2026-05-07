// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/account_network_policies#all_destinations DataDatabricksAccountNetworkPolicies#all_destinations}.
	AllDestinations interface{} `field:"optional" json:"allDestinations" yaml:"allDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/account_network_policies#workspace_api DataDatabricksAccountNetworkPolicies#workspace_api}.
	WorkspaceApi *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationWorkspaceApi `field:"optional" json:"workspaceApi" yaml:"workspaceApi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/account_network_policies#workspace_ui DataDatabricksAccountNetworkPolicies#workspace_ui}.
	WorkspaceUi *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationWorkspaceUi `field:"optional" json:"workspaceUi" yaml:"workspaceUi"`
}

