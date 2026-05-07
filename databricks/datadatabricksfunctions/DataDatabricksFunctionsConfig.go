// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfunctions

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksFunctionsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/functions#catalog_name DataDatabricksFunctions#catalog_name}.
	CatalogName *string `field:"required" json:"catalogName" yaml:"catalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/functions#schema_name DataDatabricksFunctions#schema_name}.
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/functions#functions DataDatabricksFunctions#functions}.
	Functions interface{} `field:"optional" json:"functions" yaml:"functions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/functions#include_browse DataDatabricksFunctions#include_browse}.
	IncludeBrowse interface{} `field:"optional" json:"includeBrowse" yaml:"includeBrowse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/functions#provider_config DataDatabricksFunctions#provider_config}.
	ProviderConfig *DataDatabricksFunctionsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

