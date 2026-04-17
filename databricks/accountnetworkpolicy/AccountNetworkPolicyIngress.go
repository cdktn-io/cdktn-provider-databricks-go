// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyIngress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/account_network_policy#public_access AccountNetworkPolicy#public_access}.
	PublicAccess *AccountNetworkPolicyIngressPublicAccess `field:"optional" json:"publicAccess" yaml:"publicAccess"`
}

