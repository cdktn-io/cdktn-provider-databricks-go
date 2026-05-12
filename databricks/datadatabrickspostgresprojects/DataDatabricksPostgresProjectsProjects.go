// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresprojects


type DataDatabricksPostgresProjectsProjects struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/data-sources/postgres_projects#name DataDatabricksPostgresProjects#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/data-sources/postgres_projects#provider_config DataDatabricksPostgresProjects#provider_config}.
	ProviderConfig *DataDatabricksPostgresProjectsProjectsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

