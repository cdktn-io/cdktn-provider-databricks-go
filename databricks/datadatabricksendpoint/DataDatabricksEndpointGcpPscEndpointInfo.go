// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksendpoint


type DataDatabricksEndpointGcpPscEndpointInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/endpoint#endpoint_region DataDatabricksEndpoint#endpoint_region}.
	EndpointRegion *string `field:"required" json:"endpointRegion" yaml:"endpointRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/endpoint#project_id DataDatabricksEndpoint#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/endpoint#psc_endpoint DataDatabricksEndpoint#psc_endpoint}.
	PscEndpoint *string `field:"required" json:"pscEndpoint" yaml:"pscEndpoint"`
}

