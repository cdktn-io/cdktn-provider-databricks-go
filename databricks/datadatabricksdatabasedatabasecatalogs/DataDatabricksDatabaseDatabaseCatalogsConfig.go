// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasedatabasecatalogs

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksDatabaseDatabaseCatalogsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/database_database_catalogs#instance_name DataDatabricksDatabaseDatabaseCatalogs#instance_name}.
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/database_database_catalogs#page_size DataDatabricksDatabaseDatabaseCatalogs#page_size}.
	PageSize *float64 `field:"optional" json:"pageSize" yaml:"pageSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/database_database_catalogs#provider_config DataDatabricksDatabaseDatabaseCatalogs#provider_config}.
	ProviderConfig *DataDatabricksDatabaseDatabaseCatalogsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

