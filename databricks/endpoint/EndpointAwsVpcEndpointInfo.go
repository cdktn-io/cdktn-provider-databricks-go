// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package endpoint


type EndpointAwsVpcEndpointInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/endpoint#aws_vpc_endpoint_id Endpoint#aws_vpc_endpoint_id}.
	AwsVpcEndpointId *string `field:"required" json:"awsVpcEndpointId" yaml:"awsVpcEndpointId"`
}

