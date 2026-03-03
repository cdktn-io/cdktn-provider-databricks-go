// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresendpoint


type PostgresEndpointStatusGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/resources/postgres_endpoint#max PostgresEndpoint#max}.
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/resources/postgres_endpoint#min PostgresEndpoint#min}.
	Min *float64 `field:"required" json:"min" yaml:"min"`
}

