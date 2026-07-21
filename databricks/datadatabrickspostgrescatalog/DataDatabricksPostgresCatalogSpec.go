// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgrescatalog


type DataDatabricksPostgresCatalogSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/postgres_catalog#postgres_database DataDatabricksPostgresCatalog#postgres_database}.
	PostgresDatabase *string `field:"required" json:"postgresDatabase" yaml:"postgresDatabase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/postgres_catalog#branch DataDatabricksPostgresCatalog#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/postgres_catalog#create_database_if_missing DataDatabricksPostgresCatalog#create_database_if_missing}.
	CreateDatabaseIfMissing interface{} `field:"optional" json:"createDatabaseIfMissing" yaml:"createDatabaseIfMissing"`
}

