// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy


type DataDatabricksAccountNetworkPolicyEgressNetworkAccessAllowedInternetDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/account_network_policy#destination DataDatabricksAccountNetworkPolicy#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/account_network_policy#internet_destination_type DataDatabricksAccountNetworkPolicy#internet_destination_type}.
	InternetDestinationType *string `field:"optional" json:"internetDestinationType" yaml:"internetDestinationType"`
}

