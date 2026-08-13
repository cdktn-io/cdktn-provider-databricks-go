// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/account_network_policies#authentication DataDatabricksAccountNetworkPolicies#authentication}.
	Authentication *DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/account_network_policies#destination DataDatabricksAccountNetworkPolicies#destination}.
	Destination *DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestination `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/account_network_policies#label DataDatabricksAccountNetworkPolicies#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/account_network_policies#origin DataDatabricksAccountNetworkPolicies#origin}.
	Origin *DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOrigin `field:"optional" json:"origin" yaml:"origin"`
}

