// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksendpoints


type DataDatabricksEndpointsItemsGcpPscEndpointInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/endpoints#endpoint_region DataDatabricksEndpoints#endpoint_region}.
	EndpointRegion *string `field:"required" json:"endpointRegion" yaml:"endpointRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/endpoints#project_id DataDatabricksEndpoints#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/endpoints#psc_endpoint DataDatabricksEndpoints#psc_endpoint}.
	PscEndpoint *string `field:"required" json:"pscEndpoint" yaml:"pscEndpoint"`
}

