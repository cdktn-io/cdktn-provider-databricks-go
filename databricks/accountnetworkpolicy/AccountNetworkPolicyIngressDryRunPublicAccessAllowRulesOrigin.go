// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOrigin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/account_network_policy#all_ip_ranges AccountNetworkPolicy#all_ip_ranges}.
	AllIpRanges interface{} `field:"optional" json:"allIpRanges" yaml:"allIpRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/account_network_policy#excluded_ip_ranges AccountNetworkPolicy#excluded_ip_ranges}.
	ExcludedIpRanges *AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOriginExcludedIpRanges `field:"optional" json:"excludedIpRanges" yaml:"excludedIpRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/account_network_policy#included_ip_ranges AccountNetworkPolicy#included_ip_ranges}.
	IncludedIpRanges *AccountNetworkPolicyIngressDryRunPublicAccessAllowRulesOriginIncludedIpRanges `field:"optional" json:"includedIpRanges" yaml:"includedIpRanges"`
}

