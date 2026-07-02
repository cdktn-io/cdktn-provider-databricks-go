// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyIngressDryRunPublicAccessDenyRulesAuthenticationIdentities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/account_network_policy#principal_id AccountNetworkPolicy#principal_id}.
	PrincipalId *float64 `field:"optional" json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/account_network_policy#principal_type AccountNetworkPolicy#principal_type}.
	PrincipalType *string `field:"optional" json:"principalType" yaml:"principalType"`
}

