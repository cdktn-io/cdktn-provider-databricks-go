// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksserviceprincipals

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksServicePrincipalsConfig struct {
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
	// Specifies whether to use account-level or workspace-level API.
	//
	// Valid values are `account` and `workspace`. When not set, the API level is inferred from the provider host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/service_principals#api DataDatabricksServicePrincipals#api}
	Api *string `field:"optional" json:"api" yaml:"api"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/service_principals#application_ids DataDatabricksServicePrincipals#application_ids}.
	ApplicationIds *[]*string `field:"optional" json:"applicationIds" yaml:"applicationIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/service_principals#display_name_contains DataDatabricksServicePrincipals#display_name_contains}.
	DisplayNameContains *string `field:"optional" json:"displayNameContains" yaml:"displayNameContains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/service_principals#id DataDatabricksServicePrincipals#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/service_principals#provider_config DataDatabricksServicePrincipals#provider_config}
	ProviderConfig *DataDatabricksServicePrincipalsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// service_principals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/service_principals#service_principals DataDatabricksServicePrincipals#service_principals}
	ServicePrincipals interface{} `field:"optional" json:"servicePrincipals" yaml:"servicePrincipals"`
}

