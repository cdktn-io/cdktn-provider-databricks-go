// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/data-sources/account_network_policies#all_destinations DataDatabricksAccountNetworkPolicies#all_destinations}.
	AllDestinations interface{} `field:"optional" json:"allDestinations" yaml:"allDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/data-sources/account_network_policies#workspace_api DataDatabricksAccountNetworkPolicies#workspace_api}.
	WorkspaceApi *DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationWorkspaceApi `field:"optional" json:"workspaceApi" yaml:"workspaceApi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/data-sources/account_network_policies#workspace_ui DataDatabricksAccountNetworkPolicies#workspace_ui}.
	WorkspaceUi *DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationWorkspaceUi `field:"optional" json:"workspaceUi" yaml:"workspaceUi"`
}

