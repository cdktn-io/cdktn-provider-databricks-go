// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy


type DataDatabricksAccountNetworkPolicyIngressCrossWorkspaceAccessAllowRulesOrigin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/account_network_policy#all_source_workspaces DataDatabricksAccountNetworkPolicy#all_source_workspaces}.
	AllSourceWorkspaces interface{} `field:"optional" json:"allSourceWorkspaces" yaml:"allSourceWorkspaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/account_network_policy#selected_workspaces DataDatabricksAccountNetworkPolicy#selected_workspaces}.
	SelectedWorkspaces *DataDatabricksAccountNetworkPolicyIngressCrossWorkspaceAccessAllowRulesOriginSelectedWorkspaces `field:"optional" json:"selectedWorkspaces" yaml:"selectedWorkspaces"`
}

