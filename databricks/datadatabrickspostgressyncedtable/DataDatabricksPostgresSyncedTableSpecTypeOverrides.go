// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgressyncedtable


type DataDatabricksPostgresSyncedTableSpecTypeOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/postgres_synced_table#column_name DataDatabricksPostgresSyncedTable#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/postgres_synced_table#pg_type DataDatabricksPostgresSyncedTable#pg_type}.
	PgType *string `field:"required" json:"pgType" yaml:"pgType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/postgres_synced_table#size DataDatabricksPostgresSyncedTable#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

