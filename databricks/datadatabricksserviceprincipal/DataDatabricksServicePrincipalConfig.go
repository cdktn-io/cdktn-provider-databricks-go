// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksserviceprincipal

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksServicePrincipalConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/service_principal#acl_principal_id DataDatabricksServicePrincipal#acl_principal_id}.
	AclPrincipalId *string `field:"optional" json:"aclPrincipalId" yaml:"aclPrincipalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/service_principal#active DataDatabricksServicePrincipal#active}.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Specifies whether to use account-level or workspace-level API.
	//
	// Valid values are `account` and `workspace`. When not set, the API level is inferred from the provider host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/service_principal#api DataDatabricksServicePrincipal#api}
	Api *string `field:"optional" json:"api" yaml:"api"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/service_principal#application_id DataDatabricksServicePrincipal#application_id}.
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/service_principal#display_name DataDatabricksServicePrincipal#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/service_principal#external_id DataDatabricksServicePrincipal#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/service_principal#home DataDatabricksServicePrincipal#home}.
	Home *string `field:"optional" json:"home" yaml:"home"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/service_principal#id DataDatabricksServicePrincipal#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/service_principal#provider_config DataDatabricksServicePrincipal#provider_config}
	ProviderConfig *DataDatabricksServicePrincipalProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/service_principal#repos DataDatabricksServicePrincipal#repos}.
	Repos *string `field:"optional" json:"repos" yaml:"repos"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/service_principal#scim_id DataDatabricksServicePrincipal#scim_id}.
	ScimId *string `field:"optional" json:"scimId" yaml:"scimId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/service_principal#sp_id DataDatabricksServicePrincipal#sp_id}.
	SpId *string `field:"optional" json:"spId" yaml:"spId"`
}

