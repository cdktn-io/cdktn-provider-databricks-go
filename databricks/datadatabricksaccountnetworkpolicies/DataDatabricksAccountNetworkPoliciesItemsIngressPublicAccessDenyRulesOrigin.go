// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOrigin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/account_network_policies#all_ip_ranges DataDatabricksAccountNetworkPolicies#all_ip_ranges}.
	AllIpRanges interface{} `field:"optional" json:"allIpRanges" yaml:"allIpRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/account_network_policies#excluded_ip_ranges DataDatabricksAccountNetworkPolicies#excluded_ip_ranges}.
	ExcludedIpRanges *DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOriginExcludedIpRanges `field:"optional" json:"excludedIpRanges" yaml:"excludedIpRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/account_network_policies#included_ip_ranges DataDatabricksAccountNetworkPolicies#included_ip_ranges}.
	IncludedIpRanges *DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOriginIncludedIpRanges `field:"optional" json:"includedIpRanges" yaml:"includedIpRanges"`
}

