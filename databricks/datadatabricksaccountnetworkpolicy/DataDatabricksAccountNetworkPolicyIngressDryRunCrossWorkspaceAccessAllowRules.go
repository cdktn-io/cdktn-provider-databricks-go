// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy


type DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/account_network_policy#authentication DataDatabricksAccountNetworkPolicy#authentication}.
	Authentication *DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/account_network_policy#destination DataDatabricksAccountNetworkPolicy#destination}.
	Destination *DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesDestination `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/account_network_policy#label DataDatabricksAccountNetworkPolicy#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/account_network_policy#origin DataDatabricksAccountNetworkPolicy#origin}.
	Origin *DataDatabricksAccountNetworkPolicyIngressDryRunCrossWorkspaceAccessAllowRulesOrigin `field:"optional" json:"origin" yaml:"origin"`
}

