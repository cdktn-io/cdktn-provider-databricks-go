// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package endpoint


type EndpointGcpPscEndpointInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/endpoint#endpoint_region Endpoint#endpoint_region}.
	EndpointRegion *string `field:"required" json:"endpointRegion" yaml:"endpointRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/endpoint#project_id Endpoint#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/endpoint#psc_endpoint Endpoint#psc_endpoint}.
	PscEndpoint *string `field:"required" json:"pscEndpoint" yaml:"pscEndpoint"`
}

