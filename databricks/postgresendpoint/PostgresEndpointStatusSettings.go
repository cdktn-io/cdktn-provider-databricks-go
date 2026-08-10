// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresendpoint


type PostgresEndpointStatusSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/postgres_endpoint#pg_settings PostgresEndpoint#pg_settings}.
	PgSettings *map[string]*string `field:"optional" json:"pgSettings" yaml:"pgSettings"`
}

