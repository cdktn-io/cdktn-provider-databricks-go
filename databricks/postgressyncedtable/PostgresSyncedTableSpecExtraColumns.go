// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgressyncedtable


type PostgresSyncedTableSpecExtraColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/postgres_synced_table#column_name PostgresSyncedTable#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/postgres_synced_table#column_type PostgresSyncedTable#column_type}.
	ColumnType *string `field:"required" json:"columnType" yaml:"columnType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/postgres_synced_table#compute PostgresSyncedTable#compute}.
	Compute *string `field:"optional" json:"compute" yaml:"compute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/postgres_synced_table#maintenance PostgresSyncedTable#maintenance}.
	Maintenance *string `field:"optional" json:"maintenance" yaml:"maintenance"`
}

