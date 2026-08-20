// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgressyncedtable


type DataDatabricksPostgresSyncedTableSpecExtraColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/postgres_synced_table#column_name DataDatabricksPostgresSyncedTable#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/postgres_synced_table#column_type DataDatabricksPostgresSyncedTable#column_type}.
	ColumnType *string `field:"required" json:"columnType" yaml:"columnType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/postgres_synced_table#compute DataDatabricksPostgresSyncedTable#compute}.
	Compute *string `field:"optional" json:"compute" yaml:"compute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/postgres_synced_table#maintenance DataDatabricksPostgresSyncedTable#maintenance}.
	Maintenance *string `field:"optional" json:"maintenance" yaml:"maintenance"`
}

