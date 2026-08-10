// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasesynceddatabasetable


type DataDatabricksDatabaseSyncedDatabaseTableSpecTypeOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/database_synced_database_table#column_name DataDatabricksDatabaseSyncedDatabaseTable#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/database_synced_database_table#pg_type DataDatabricksDatabaseSyncedDatabaseTable#pg_type}.
	PgType *string `field:"required" json:"pgType" yaml:"pgType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/database_synced_database_table#size DataDatabricksDatabaseSyncedDatabaseTable#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

