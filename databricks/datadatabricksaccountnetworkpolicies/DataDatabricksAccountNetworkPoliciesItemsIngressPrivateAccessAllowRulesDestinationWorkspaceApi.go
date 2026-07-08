// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationWorkspaceApi struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/account_network_policies#scope_qualifier DataDatabricksAccountNetworkPolicies#scope_qualifier}.
	ScopeQualifier *string `field:"optional" json:"scopeQualifier" yaml:"scopeQualifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/account_network_policies#scopes DataDatabricksAccountNetworkPolicies#scopes}.
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

