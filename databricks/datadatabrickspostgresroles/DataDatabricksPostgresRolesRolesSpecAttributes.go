// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresroles


type DataDatabricksPostgresRolesRolesSpecAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/postgres_roles#bypassrls DataDatabricksPostgresRoles#bypassrls}.
	Bypassrls interface{} `field:"optional" json:"bypassrls" yaml:"bypassrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/postgres_roles#createdb DataDatabricksPostgresRoles#createdb}.
	Createdb interface{} `field:"optional" json:"createdb" yaml:"createdb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/postgres_roles#createrole DataDatabricksPostgresRoles#createrole}.
	Createrole interface{} `field:"optional" json:"createrole" yaml:"createrole"`
}

