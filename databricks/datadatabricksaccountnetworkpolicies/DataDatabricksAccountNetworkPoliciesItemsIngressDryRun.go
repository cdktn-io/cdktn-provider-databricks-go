// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItemsIngressDryRun struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/data-sources/account_network_policies#private_access DataDatabricksAccountNetworkPolicies#private_access}.
	PrivateAccess *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccess `field:"optional" json:"privateAccess" yaml:"privateAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/data-sources/account_network_policies#public_access DataDatabricksAccountNetworkPolicies#public_access}.
	PublicAccess *DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccess `field:"optional" json:"publicAccess" yaml:"publicAccess"`
}

