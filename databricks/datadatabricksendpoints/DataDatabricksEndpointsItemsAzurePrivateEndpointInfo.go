// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksendpoints


type DataDatabricksEndpointsItemsAzurePrivateEndpointInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/endpoints#private_endpoint_name DataDatabricksEndpoints#private_endpoint_name}.
	PrivateEndpointName *string `field:"required" json:"privateEndpointName" yaml:"privateEndpointName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/endpoints#private_endpoint_resource_guid DataDatabricksEndpoints#private_endpoint_resource_guid}.
	PrivateEndpointResourceGuid *string `field:"required" json:"privateEndpointResourceGuid" yaml:"privateEndpointResourceGuid"`
}

