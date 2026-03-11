// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksendpoint


type DataDatabricksEndpointAzurePrivateEndpointInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.111.0/docs/data-sources/endpoint#private_endpoint_name DataDatabricksEndpoint#private_endpoint_name}.
	PrivateEndpointName *string `field:"required" json:"privateEndpointName" yaml:"privateEndpointName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.111.0/docs/data-sources/endpoint#private_endpoint_resource_guid DataDatabricksEndpoint#private_endpoint_resource_guid}.
	PrivateEndpointResourceGuid *string `field:"required" json:"privateEndpointResourceGuid" yaml:"privateEndpointResourceGuid"`
}

