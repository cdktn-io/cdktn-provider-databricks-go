// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresdatabases


type DataDatabricksPostgresDatabasesDatabases struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/postgres_databases#name DataDatabricksPostgresDatabases#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/postgres_databases#provider_config DataDatabricksPostgresDatabases#provider_config}.
	ProviderConfig *DataDatabricksPostgresDatabasesDatabasesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

