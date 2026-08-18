// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyIngressDryRun struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/account_network_policy#cross_workspace_access AccountNetworkPolicy#cross_workspace_access}.
	CrossWorkspaceAccess *AccountNetworkPolicyIngressDryRunCrossWorkspaceAccess `field:"optional" json:"crossWorkspaceAccess" yaml:"crossWorkspaceAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/account_network_policy#private_access AccountNetworkPolicy#private_access}.
	PrivateAccess *AccountNetworkPolicyIngressDryRunPrivateAccess `field:"optional" json:"privateAccess" yaml:"privateAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/account_network_policy#public_access AccountNetworkPolicy#public_access}.
	PublicAccess *AccountNetworkPolicyIngressDryRunPublicAccess `field:"optional" json:"publicAccess" yaml:"publicAccess"`
}

