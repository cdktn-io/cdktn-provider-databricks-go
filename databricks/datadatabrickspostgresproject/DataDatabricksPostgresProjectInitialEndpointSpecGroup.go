// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresproject


type DataDatabricksPostgresProjectInitialEndpointSpecGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/postgres_project#max DataDatabricksPostgresProject#max}.
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/postgres_project#min DataDatabricksPostgresProject#min}.
	Min *float64 `field:"required" json:"min" yaml:"min"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/postgres_project#enable_readable_secondaries DataDatabricksPostgresProject#enable_readable_secondaries}.
	EnableReadableSecondaries interface{} `field:"optional" json:"enableReadableSecondaries" yaml:"enableReadableSecondaries"`
}

