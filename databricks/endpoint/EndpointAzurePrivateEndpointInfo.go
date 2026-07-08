// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package endpoint


type EndpointAzurePrivateEndpointInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/endpoint#private_endpoint_name Endpoint#private_endpoint_name}.
	PrivateEndpointName *string `field:"required" json:"privateEndpointName" yaml:"privateEndpointName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/endpoint#private_endpoint_resource_guid Endpoint#private_endpoint_resource_guid}.
	PrivateEndpointResourceGuid *string `field:"required" json:"privateEndpointResourceGuid" yaml:"privateEndpointResourceGuid"`
}

