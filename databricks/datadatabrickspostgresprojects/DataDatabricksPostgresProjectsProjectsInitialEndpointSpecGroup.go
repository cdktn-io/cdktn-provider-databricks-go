// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresprojects


type DataDatabricksPostgresProjectsProjectsInitialEndpointSpecGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/postgres_projects#max DataDatabricksPostgresProjects#max}.
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/postgres_projects#min DataDatabricksPostgresProjects#min}.
	Min *float64 `field:"required" json:"min" yaml:"min"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/postgres_projects#enable_readable_secondaries DataDatabricksPostgresProjects#enable_readable_secondaries}.
	EnableReadableSecondaries interface{} `field:"optional" json:"enableReadableSecondaries" yaml:"enableReadableSecondaries"`
}

