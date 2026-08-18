// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy


type DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/account_network_policy#restriction_mode DataDatabricksAccountNetworkPolicy#restriction_mode}.
	RestrictionMode *string `field:"required" json:"restrictionMode" yaml:"restrictionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/account_network_policy#allow_rules DataDatabricksAccountNetworkPolicy#allow_rules}.
	AllowRules interface{} `field:"optional" json:"allowRules" yaml:"allowRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/account_network_policy#deny_rules DataDatabricksAccountNetworkPolicy#deny_rules}.
	DenyRules interface{} `field:"optional" json:"denyRules" yaml:"denyRules"`
}

