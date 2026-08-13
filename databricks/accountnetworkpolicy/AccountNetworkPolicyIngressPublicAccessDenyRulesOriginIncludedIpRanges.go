// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyIngressPublicAccessDenyRulesOriginIncludedIpRanges struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/account_network_policy#ip_ranges AccountNetworkPolicy#ip_ranges}.
	IpRanges *[]*string `field:"optional" json:"ipRanges" yaml:"ipRanges"`
}

