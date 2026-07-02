// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresrole


type PostgresRoleStatusAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/postgres_role#bypassrls PostgresRole#bypassrls}.
	Bypassrls interface{} `field:"optional" json:"bypassrls" yaml:"bypassrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/postgres_role#createdb PostgresRole#createdb}.
	Createdb interface{} `field:"optional" json:"createdb" yaml:"createdb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/postgres_role#createrole PostgresRole#createrole}.
	Createrole interface{} `field:"optional" json:"createrole" yaml:"createrole"`
}

