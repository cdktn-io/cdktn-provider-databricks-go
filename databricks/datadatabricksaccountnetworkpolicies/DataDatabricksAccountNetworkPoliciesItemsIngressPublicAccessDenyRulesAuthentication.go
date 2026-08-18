// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/account_network_policies#identities DataDatabricksAccountNetworkPolicies#identities}.
	Identities interface{} `field:"optional" json:"identities" yaml:"identities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/account_network_policies#identity_type DataDatabricksAccountNetworkPolicies#identity_type}.
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
}

