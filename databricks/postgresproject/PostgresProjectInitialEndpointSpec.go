// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresproject


type PostgresProjectInitialEndpointSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.111.0/docs/resources/postgres_project#group PostgresProject#group}.
	Group *PostgresProjectInitialEndpointSpecGroup `field:"optional" json:"group" yaml:"group"`
}

