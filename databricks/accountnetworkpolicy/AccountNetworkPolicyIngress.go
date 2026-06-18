// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyIngress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/account_network_policy#cross_workspace_access AccountNetworkPolicy#cross_workspace_access}.
	CrossWorkspaceAccess *AccountNetworkPolicyIngressCrossWorkspaceAccess `field:"optional" json:"crossWorkspaceAccess" yaml:"crossWorkspaceAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/account_network_policy#private_access AccountNetworkPolicy#private_access}.
	PrivateAccess *AccountNetworkPolicyIngressPrivateAccess `field:"optional" json:"privateAccess" yaml:"privateAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/account_network_policy#public_access AccountNetworkPolicy#public_access}.
	PublicAccess *AccountNetworkPolicyIngressPublicAccess `field:"optional" json:"publicAccess" yaml:"publicAccess"`
}

