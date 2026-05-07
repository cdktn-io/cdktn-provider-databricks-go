// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyIngressPublicAccessDenyRulesDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/account_network_policy#all_destinations AccountNetworkPolicy#all_destinations}.
	AllDestinations interface{} `field:"optional" json:"allDestinations" yaml:"allDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/account_network_policy#workspace_api AccountNetworkPolicy#workspace_api}.
	WorkspaceApi *AccountNetworkPolicyIngressPublicAccessDenyRulesDestinationWorkspaceApi `field:"optional" json:"workspaceApi" yaml:"workspaceApi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/account_network_policy#workspace_ui AccountNetworkPolicy#workspace_ui}.
	WorkspaceUi *AccountNetworkPolicyIngressPublicAccessDenyRulesDestinationWorkspaceUi `field:"optional" json:"workspaceUi" yaml:"workspaceUi"`
}

