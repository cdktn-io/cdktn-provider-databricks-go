// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy


type DataDatabricksAccountNetworkPolicyIngressPublicAccessAllowRulesAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/account_network_policy#identities DataDatabricksAccountNetworkPolicy#identities}.
	Identities interface{} `field:"optional" json:"identities" yaml:"identities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/account_network_policy#identity_type DataDatabricksAccountNetworkPolicy#identity_type}.
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
}

