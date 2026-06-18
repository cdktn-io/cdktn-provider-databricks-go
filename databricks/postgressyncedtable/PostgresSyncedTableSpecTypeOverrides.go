// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgressyncedtable


type PostgresSyncedTableSpecTypeOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/postgres_synced_table#column_name PostgresSyncedTable#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/postgres_synced_table#pg_type PostgresSyncedTable#pg_type}.
	PgType *string `field:"required" json:"pgType" yaml:"pgType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/postgres_synced_table#size PostgresSyncedTable#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

