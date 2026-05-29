// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresdatabase


type DataDatabricksPostgresDatabaseSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/postgres_database#postgres_database DataDatabricksPostgresDatabase#postgres_database}.
	PostgresDatabase *string `field:"optional" json:"postgresDatabase" yaml:"postgresDatabase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/postgres_database#role DataDatabricksPostgresDatabase#role}.
	Role *string `field:"optional" json:"role" yaml:"role"`
}

