// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresdatabase


type PostgresDatabaseSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/postgres_database#postgres_database PostgresDatabase#postgres_database}.
	PostgresDatabase *string `field:"optional" json:"postgresDatabase" yaml:"postgresDatabase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/postgres_database#role PostgresDatabase#role}.
	Role *string `field:"optional" json:"role" yaml:"role"`
}

