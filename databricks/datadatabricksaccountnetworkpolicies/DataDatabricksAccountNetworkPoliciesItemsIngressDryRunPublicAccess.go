// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/account_network_policies#restriction_mode DataDatabricksAccountNetworkPolicies#restriction_mode}.
	RestrictionMode *string `field:"required" json:"restrictionMode" yaml:"restrictionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/account_network_policies#allow_rules DataDatabricksAccountNetworkPolicies#allow_rules}.
	AllowRules interface{} `field:"optional" json:"allowRules" yaml:"allowRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/account_network_policies#deny_rules DataDatabricksAccountNetworkPolicies#deny_rules}.
	DenyRules interface{} `field:"optional" json:"denyRules" yaml:"denyRules"`
}

