// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy


type DataDatabricksAccountNetworkPolicyIngressPublicAccessDenyRulesOrigin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/data-sources/account_network_policy#all_ip_ranges DataDatabricksAccountNetworkPolicy#all_ip_ranges}.
	AllIpRanges interface{} `field:"optional" json:"allIpRanges" yaml:"allIpRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/data-sources/account_network_policy#excluded_ip_ranges DataDatabricksAccountNetworkPolicy#excluded_ip_ranges}.
	ExcludedIpRanges *DataDatabricksAccountNetworkPolicyIngressPublicAccessDenyRulesOriginExcludedIpRanges `field:"optional" json:"excludedIpRanges" yaml:"excludedIpRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/data-sources/account_network_policy#included_ip_ranges DataDatabricksAccountNetworkPolicy#included_ip_ranges}.
	IncludedIpRanges *DataDatabricksAccountNetworkPolicyIngressPublicAccessDenyRulesOriginIncludedIpRanges `field:"optional" json:"includedIpRanges" yaml:"includedIpRanges"`
}

