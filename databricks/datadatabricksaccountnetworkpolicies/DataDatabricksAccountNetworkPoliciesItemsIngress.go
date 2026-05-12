// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItemsIngress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/data-sources/account_network_policies#private_access DataDatabricksAccountNetworkPolicies#private_access}.
	PrivateAccess *DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccess `field:"optional" json:"privateAccess" yaml:"privateAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/data-sources/account_network_policies#public_access DataDatabricksAccountNetworkPolicies#public_access}.
	PublicAccess *DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccess `field:"optional" json:"publicAccess" yaml:"publicAccess"`
}

