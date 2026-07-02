// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy


type DataDatabricksAccountNetworkPolicyIngress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/account_network_policy#cross_workspace_access DataDatabricksAccountNetworkPolicy#cross_workspace_access}.
	CrossWorkspaceAccess *DataDatabricksAccountNetworkPolicyIngressCrossWorkspaceAccess `field:"optional" json:"crossWorkspaceAccess" yaml:"crossWorkspaceAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/account_network_policy#private_access DataDatabricksAccountNetworkPolicy#private_access}.
	PrivateAccess *DataDatabricksAccountNetworkPolicyIngressPrivateAccess `field:"optional" json:"privateAccess" yaml:"privateAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/account_network_policy#public_access DataDatabricksAccountNetworkPolicy#public_access}.
	PublicAccess *DataDatabricksAccountNetworkPolicyIngressPublicAccess `field:"optional" json:"publicAccess" yaml:"publicAccess"`
}

