// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasedatabasecatalog

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabaseDatabaseCatalogConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/database_database_catalog#database_instance_name DatabaseDatabaseCatalog#database_instance_name}.
	DatabaseInstanceName *string `field:"required" json:"databaseInstanceName" yaml:"databaseInstanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/database_database_catalog#database_name DatabaseDatabaseCatalog#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/database_database_catalog#name DatabaseDatabaseCatalog#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/database_database_catalog#create_database_if_not_exists DatabaseDatabaseCatalog#create_database_if_not_exists}.
	CreateDatabaseIfNotExists interface{} `field:"optional" json:"createDatabaseIfNotExists" yaml:"createDatabaseIfNotExists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/database_database_catalog#provider_config DatabaseDatabaseCatalog#provider_config}.
	ProviderConfig *DatabaseDatabaseCatalogProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

