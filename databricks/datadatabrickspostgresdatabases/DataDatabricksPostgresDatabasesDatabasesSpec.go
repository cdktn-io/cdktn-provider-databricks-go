// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresdatabases


type DataDatabricksPostgresDatabasesDatabasesSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/postgres_databases#role DataDatabricksPostgresDatabases#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/postgres_databases#postgres_database DataDatabricksPostgresDatabases#postgres_database}.
	PostgresDatabase *string `field:"optional" json:"postgresDatabase" yaml:"postgresDatabase"`
}

