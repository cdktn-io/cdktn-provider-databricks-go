// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyEgressNetworkAccessBlockedInternetDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/account_network_policy#destination AccountNetworkPolicy#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/account_network_policy#internet_destination_type AccountNetworkPolicy#internet_destination_type}.
	InternetDestinationType *string `field:"optional" json:"internetDestinationType" yaml:"internetDestinationType"`
}

