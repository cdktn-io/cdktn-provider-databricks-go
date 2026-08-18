// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresrole


type DataDatabricksPostgresRoleSpecAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/postgres_role#bypassrls DataDatabricksPostgresRole#bypassrls}.
	Bypassrls interface{} `field:"optional" json:"bypassrls" yaml:"bypassrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/postgres_role#createdb DataDatabricksPostgresRole#createdb}.
	Createdb interface{} `field:"optional" json:"createdb" yaml:"createdb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/postgres_role#createrole DataDatabricksPostgresRole#createrole}.
	Createrole interface{} `field:"optional" json:"createrole" yaml:"createrole"`
}

