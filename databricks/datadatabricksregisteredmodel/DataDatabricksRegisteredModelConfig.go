// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksregisteredmodel

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksRegisteredModelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/registered_model#full_name DataDatabricksRegisteredModel#full_name}.
	FullName *string `field:"required" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/registered_model#include_aliases DataDatabricksRegisteredModel#include_aliases}.
	IncludeAliases interface{} `field:"optional" json:"includeAliases" yaml:"includeAliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/registered_model#include_browse DataDatabricksRegisteredModel#include_browse}.
	IncludeBrowse interface{} `field:"optional" json:"includeBrowse" yaml:"includeBrowse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/registered_model#model_info DataDatabricksRegisteredModel#model_info}.
	ModelInfo interface{} `field:"optional" json:"modelInfo" yaml:"modelInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/registered_model#provider_config DataDatabricksRegisteredModel#provider_config}.
	ProviderConfig *DataDatabricksRegisteredModelProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

