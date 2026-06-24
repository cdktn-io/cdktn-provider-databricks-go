// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyIngressPublicAccessDenyRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/account_network_policy#authentication AccountNetworkPolicy#authentication}.
	Authentication *AccountNetworkPolicyIngressPublicAccessDenyRulesAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/account_network_policy#destination AccountNetworkPolicy#destination}.
	Destination *AccountNetworkPolicyIngressPublicAccessDenyRulesDestination `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/account_network_policy#label AccountNetworkPolicy#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/account_network_policy#origin AccountNetworkPolicy#origin}.
	Origin *AccountNetworkPolicyIngressPublicAccessDenyRulesOrigin `field:"optional" json:"origin" yaml:"origin"`
}

