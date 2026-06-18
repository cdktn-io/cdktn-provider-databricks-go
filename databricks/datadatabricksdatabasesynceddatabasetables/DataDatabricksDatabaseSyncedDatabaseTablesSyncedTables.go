// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasesynceddatabasetables


type DataDatabricksDatabaseSyncedDatabaseTablesSyncedTables struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/database_synced_database_tables#name DataDatabricksDatabaseSyncedDatabaseTables#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/database_synced_database_tables#provider_config DataDatabricksDatabaseSyncedDatabaseTables#provider_config}.
	ProviderConfig *DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

