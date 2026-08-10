// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesynceddatabasetable


type DatabaseSyncedDatabaseTableSpecTypeOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/database_synced_database_table#column_name DatabaseSyncedDatabaseTable#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/database_synced_database_table#pg_type DatabaseSyncedDatabaseTable#pg_type}.
	PgType *string `field:"required" json:"pgType" yaml:"pgType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/database_synced_database_table#size DatabaseSyncedDatabaseTable#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

