// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package privatenetworkgateway


type PrivateNetworkGatewayAzureCloudConnection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/resources/private_network_gateway#gateway_subnet PrivateNetworkGateway#gateway_subnet}.
	GatewaySubnet *PrivateNetworkGatewayAzureCloudConnectionGatewaySubnet `field:"required" json:"gatewaySubnet" yaml:"gatewaySubnet"`
}

