// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/account_network_policy#authentication AccountNetworkPolicy#authentication}.
	Authentication *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/account_network_policy#destination AccountNetworkPolicy#destination}.
	Destination *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesDestination `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/account_network_policy#label AccountNetworkPolicy#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/account_network_policy#origin AccountNetworkPolicy#origin}.
	Origin *AccountNetworkPolicyIngressCrossWorkspaceAccessDenyRulesOrigin `field:"optional" json:"origin" yaml:"origin"`
}

