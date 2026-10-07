// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksprivatenetworkgateway


type DataDatabricksPrivateNetworkGatewayPrivateDnsResolvers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/data-sources/private_network_gateway#resolver_type DataDatabricksPrivateNetworkGateway#resolver_type}.
	ResolverType *string `field:"required" json:"resolverType" yaml:"resolverType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/data-sources/private_network_gateway#value DataDatabricksPrivateNetworkGateway#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

