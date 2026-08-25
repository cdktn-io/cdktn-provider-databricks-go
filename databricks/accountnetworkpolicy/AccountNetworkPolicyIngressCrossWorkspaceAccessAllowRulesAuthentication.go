// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyIngressCrossWorkspaceAccessAllowRulesAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/account_network_policy#identities AccountNetworkPolicy#identities}.
	Identities interface{} `field:"optional" json:"identities" yaml:"identities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/account_network_policy#identity_type AccountNetworkPolicy#identity_type}.
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
}

