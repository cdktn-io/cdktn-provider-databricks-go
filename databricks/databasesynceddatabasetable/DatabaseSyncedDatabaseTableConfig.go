// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesynceddatabasetable

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabaseSyncedDatabaseTableConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/database_synced_database_table#name DatabaseSyncedDatabaseTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/database_synced_database_table#database_instance_name DatabaseSyncedDatabaseTable#database_instance_name}.
	DatabaseInstanceName *string `field:"optional" json:"databaseInstanceName" yaml:"databaseInstanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/database_synced_database_table#logical_database_name DatabaseSyncedDatabaseTable#logical_database_name}.
	LogicalDatabaseName *string `field:"optional" json:"logicalDatabaseName" yaml:"logicalDatabaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/database_synced_database_table#provider_config DatabaseSyncedDatabaseTable#provider_config}.
	ProviderConfig *DatabaseSyncedDatabaseTableProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/database_synced_database_table#spec DatabaseSyncedDatabaseTable#spec}.
	Spec *DatabaseSyncedDatabaseTableSpec `field:"optional" json:"spec" yaml:"spec"`
}

