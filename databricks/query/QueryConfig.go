// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package query

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QueryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/query#display_name Query#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/query#query_text Query#query_text}.
	QueryText *string `field:"required" json:"queryText" yaml:"queryText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/query#warehouse_id Query#warehouse_id}.
	WarehouseId *string `field:"required" json:"warehouseId" yaml:"warehouseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/query#apply_auto_limit Query#apply_auto_limit}.
	ApplyAutoLimit interface{} `field:"optional" json:"applyAutoLimit" yaml:"applyAutoLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/query#catalog Query#catalog}.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/query#description Query#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/query#owner_user_name Query#owner_user_name}.
	OwnerUserName *string `field:"optional" json:"ownerUserName" yaml:"ownerUserName"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/query#parameter Query#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/query#parent_path Query#parent_path}.
	ParentPath *string `field:"optional" json:"parentPath" yaml:"parentPath"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/query#provider_config Query#provider_config}
	ProviderConfig *QueryProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/query#run_as_mode Query#run_as_mode}.
	RunAsMode *string `field:"optional" json:"runAsMode" yaml:"runAsMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/query#schema Query#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/query#tags Query#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

