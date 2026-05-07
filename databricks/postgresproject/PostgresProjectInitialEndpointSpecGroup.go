// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresproject


type PostgresProjectInitialEndpointSpecGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/postgres_project#max PostgresProject#max}.
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/postgres_project#min PostgresProject#min}.
	Min *float64 `field:"required" json:"min" yaml:"min"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/postgres_project#enable_readable_secondaries PostgresProject#enable_readable_secondaries}.
	EnableReadableSecondaries interface{} `field:"optional" json:"enableReadableSecondaries" yaml:"enableReadableSecondaries"`
}

