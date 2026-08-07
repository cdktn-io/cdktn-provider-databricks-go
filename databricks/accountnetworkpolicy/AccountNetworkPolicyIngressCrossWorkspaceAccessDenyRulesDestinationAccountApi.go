// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestinationAccountApi struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/account_network_policy#scope_qualifier AccountNetworkPolicy#scope_qualifier}.
	ScopeQualifier *string `field:"optional" json:"scopeQualifier" yaml:"scopeQualifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/account_network_policy#scopes AccountNetworkPolicy#scopes}.
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

