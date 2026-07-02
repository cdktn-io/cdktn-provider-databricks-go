// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresendpoint


type DataDatabricksPostgresEndpointStatusSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/postgres_endpoint#pg_settings DataDatabricksPostgresEndpoint#pg_settings}.
	PgSettings *map[string]*string `field:"optional" json:"pgSettings" yaml:"pgSettings"`
}

