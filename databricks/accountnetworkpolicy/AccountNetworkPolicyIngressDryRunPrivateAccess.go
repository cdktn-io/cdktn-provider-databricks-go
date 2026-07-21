// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyIngressDryRunPrivateAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/account_network_policy#restriction_mode AccountNetworkPolicy#restriction_mode}.
	RestrictionMode *string `field:"required" json:"restrictionMode" yaml:"restrictionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/account_network_policy#allow_rules AccountNetworkPolicy#allow_rules}.
	AllowRules interface{} `field:"optional" json:"allowRules" yaml:"allowRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/account_network_policy#deny_rules AccountNetworkPolicy#deny_rules}.
	DenyRules interface{} `field:"optional" json:"denyRules" yaml:"denyRules"`
}

