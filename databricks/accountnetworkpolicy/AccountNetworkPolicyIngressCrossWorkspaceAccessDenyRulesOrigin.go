// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesOrigin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/account_network_policy#all_source_workspaces AccountNetworkPolicy#all_source_workspaces}.
	AllSourceWorkspaces interface{} `field:"optional" json:"allSourceWorkspaces" yaml:"allSourceWorkspaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/account_network_policy#selected_workspaces AccountNetworkPolicy#selected_workspaces}.
	SelectedWorkspaces *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesOriginSelectedWorkspaces `field:"optional" json:"selectedWorkspaces" yaml:"selectedWorkspaces"`
}

