// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalogworkspacebinding

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CatalogWorkspaceBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/catalog_workspace_binding#workspace_id CatalogWorkspaceBinding#workspace_id}.
	WorkspaceId *float64 `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/catalog_workspace_binding#binding_type CatalogWorkspaceBinding#binding_type}.
	BindingType *string `field:"optional" json:"bindingType" yaml:"bindingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/catalog_workspace_binding#catalog_name CatalogWorkspaceBinding#catalog_name}.
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/catalog_workspace_binding#id CatalogWorkspaceBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/catalog_workspace_binding#provider_config CatalogWorkspaceBinding#provider_config}
	ProviderConfig *CatalogWorkspaceBindingProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/catalog_workspace_binding#securable_name CatalogWorkspaceBinding#securable_name}.
	SecurableName *string `field:"optional" json:"securableName" yaml:"securableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/catalog_workspace_binding#securable_type CatalogWorkspaceBinding#securable_type}.
	SecurableType *string `field:"optional" json:"securableType" yaml:"securableType"`
}

