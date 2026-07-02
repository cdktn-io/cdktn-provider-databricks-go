// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasesynceddatabasetables


type DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/database_synced_database_tables#column_name DataDatabricksDatabaseSyncedDatabaseTables#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/database_synced_database_tables#pg_type DataDatabricksDatabaseSyncedDatabaseTables#pg_type}.
	PgType *string `field:"required" json:"pgType" yaml:"pgType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/database_synced_database_tables#size DataDatabricksDatabaseSyncedDatabaseTables#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

