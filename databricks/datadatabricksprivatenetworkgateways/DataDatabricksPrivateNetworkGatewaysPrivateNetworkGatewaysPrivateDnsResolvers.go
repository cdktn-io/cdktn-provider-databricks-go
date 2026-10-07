// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksprivatenetworkgateways


type DataDatabricksPrivateNetworkGatewaysPrivateNetworkGatewaysPrivateDnsResolvers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/data-sources/private_network_gateways#resolver_type DataDatabricksPrivateNetworkGateways#resolver_type}.
	ResolverType *string `field:"required" json:"resolverType" yaml:"resolverType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/data-sources/private_network_gateways#value DataDatabricksPrivateNetworkGateways#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

