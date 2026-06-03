// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyEgressNetworkAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/account_network_policy#restriction_mode AccountNetworkPolicy#restriction_mode}.
	RestrictionMode *string `field:"required" json:"restrictionMode" yaml:"restrictionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/account_network_policy#allowed_internet_destinations AccountNetworkPolicy#allowed_internet_destinations}.
	AllowedInternetDestinations interface{} `field:"optional" json:"allowedInternetDestinations" yaml:"allowedInternetDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/account_network_policy#allowed_storage_destinations AccountNetworkPolicy#allowed_storage_destinations}.
	AllowedStorageDestinations interface{} `field:"optional" json:"allowedStorageDestinations" yaml:"allowedStorageDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/account_network_policy#blocked_internet_destinations AccountNetworkPolicy#blocked_internet_destinations}.
	BlockedInternetDestinations interface{} `field:"optional" json:"blockedInternetDestinations" yaml:"blockedInternetDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/account_network_policy#policy_enforcement AccountNetworkPolicy#policy_enforcement}.
	PolicyEnforcement *AccountNetworkPolicyEgressNetworkAccessPolicyEnforcement `field:"optional" json:"policyEnforcement" yaml:"policyEnforcement"`
}

