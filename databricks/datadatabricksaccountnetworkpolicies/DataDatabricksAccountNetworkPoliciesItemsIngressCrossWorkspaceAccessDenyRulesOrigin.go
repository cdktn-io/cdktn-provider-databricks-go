// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesOrigin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/account_network_policies#all_source_workspaces DataDatabricksAccountNetworkPolicies#all_source_workspaces}.
	AllSourceWorkspaces interface{} `field:"optional" json:"allSourceWorkspaces" yaml:"allSourceWorkspaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/account_network_policies#selected_workspaces DataDatabricksAccountNetworkPolicies#selected_workspaces}.
	SelectedWorkspaces *DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesOriginSelectedWorkspaces `field:"optional" json:"selectedWorkspaces" yaml:"selectedWorkspaces"`
}

