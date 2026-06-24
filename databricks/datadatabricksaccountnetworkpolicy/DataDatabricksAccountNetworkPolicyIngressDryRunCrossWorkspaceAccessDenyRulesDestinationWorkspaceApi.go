// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy


type DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessDenyRulesDestinationWorkspaceApi struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/account_network_policy#scope_qualifier DataDatabricksAccountNetworkPolicy#scope_qualifier}.
	ScopeQualifier *string `field:"optional" json:"scopeQualifier" yaml:"scopeQualifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/account_network_policy#scopes DataDatabricksAccountNetworkPolicy#scopes}.
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

