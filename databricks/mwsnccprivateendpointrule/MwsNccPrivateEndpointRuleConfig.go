// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mwsnccprivateendpointrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MwsNccPrivateEndpointRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/mws_ncc_private_endpoint_rule#network_connectivity_config_id MwsNccPrivateEndpointRule#network_connectivity_config_id}.
	NetworkConnectivityConfigId *string `field:"required" json:"networkConnectivityConfigId" yaml:"networkConnectivityConfigId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/mws_ncc_private_endpoint_rule#domain_names MwsNccPrivateEndpointRule#domain_names}.
	DomainNames *[]*string `field:"optional" json:"domainNames" yaml:"domainNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/mws_ncc_private_endpoint_rule#enabled MwsNccPrivateEndpointRule#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/mws_ncc_private_endpoint_rule#endpoint_service MwsNccPrivateEndpointRule#endpoint_service}.
	EndpointService *string `field:"optional" json:"endpointService" yaml:"endpointService"`
	// gcp_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/mws_ncc_private_endpoint_rule#gcp_endpoint MwsNccPrivateEndpointRule#gcp_endpoint}
	GcpEndpoint *MwsNccPrivateEndpointRuleGcpEndpoint `field:"optional" json:"gcpEndpoint" yaml:"gcpEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/mws_ncc_private_endpoint_rule#group_id MwsNccPrivateEndpointRule#group_id}.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/mws_ncc_private_endpoint_rule#id MwsNccPrivateEndpointRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/mws_ncc_private_endpoint_rule#resource_id MwsNccPrivateEndpointRule#resource_id}.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/mws_ncc_private_endpoint_rule#resource_names MwsNccPrivateEndpointRule#resource_names}.
	ResourceNames *[]*string `field:"optional" json:"resourceNames" yaml:"resourceNames"`
}

