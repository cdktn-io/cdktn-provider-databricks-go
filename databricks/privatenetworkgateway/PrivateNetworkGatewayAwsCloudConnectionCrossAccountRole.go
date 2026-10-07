// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package privatenetworkgateway


type PrivateNetworkGatewayAwsCloudConnectionCrossAccountRole struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/resources/private_network_gateway#role_arn PrivateNetworkGateway#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

