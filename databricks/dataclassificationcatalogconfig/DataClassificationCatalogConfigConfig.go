// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataclassificationcatalogconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataClassificationCatalogConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/data_classification_catalog_config#parent DataClassificationCatalogConfig#parent}.
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/data_classification_catalog_config#auto_tag_configs DataClassificationCatalogConfig#auto_tag_configs}.
	AutoTagConfigs interface{} `field:"optional" json:"autoTagConfigs" yaml:"autoTagConfigs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/data_classification_catalog_config#included_schemas DataClassificationCatalogConfig#included_schemas}.
	IncludedSchemas *DataClassificationCatalogConfigIncludedSchemas `field:"optional" json:"includedSchemas" yaml:"includedSchemas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/data_classification_catalog_config#provider_config DataClassificationCatalogConfig#provider_config}.
	ProviderConfig *DataClassificationCatalogConfigProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

