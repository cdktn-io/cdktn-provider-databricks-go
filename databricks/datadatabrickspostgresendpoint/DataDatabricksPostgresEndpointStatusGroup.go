// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresendpoint


type DataDatabricksPostgresEndpointStatusGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/data-sources/postgres_endpoint#max DataDatabricksPostgresEndpoint#max}.
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/data-sources/postgres_endpoint#min DataDatabricksPostgresEndpoint#min}.
	Min *float64 `field:"required" json:"min" yaml:"min"`
}

