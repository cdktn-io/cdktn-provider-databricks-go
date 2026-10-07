// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksprivatenetworkgateway


type DataDatabricksPrivateNetworkGatewayAzureCloudConnection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/data-sources/private_network_gateway#gateway_subnet DataDatabricksPrivateNetworkGateway#gateway_subnet}.
	GatewaySubnet *DataDatabricksPrivateNetworkGatewayAzureCloudConnectionGatewaySubnet `field:"required" json:"gatewaySubnet" yaml:"gatewaySubnet"`
}

