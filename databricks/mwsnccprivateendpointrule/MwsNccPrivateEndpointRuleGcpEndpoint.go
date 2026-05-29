// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mwsnccprivateendpointrule


type MwsNccPrivateEndpointRuleGcpEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/mws_ncc_private_endpoint_rule#psc_endpoint_uri MwsNccPrivateEndpointRule#psc_endpoint_uri}.
	PscEndpointUri *string `field:"optional" json:"pscEndpointUri" yaml:"pscEndpointUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/mws_ncc_private_endpoint_rule#service_attachment MwsNccPrivateEndpointRule#service_attachment}.
	ServiceAttachment *string `field:"optional" json:"serviceAttachment" yaml:"serviceAttachment"`
}

