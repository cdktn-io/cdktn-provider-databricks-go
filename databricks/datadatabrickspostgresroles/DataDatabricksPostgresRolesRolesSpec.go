// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresroles


type DataDatabricksPostgresRolesRolesSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/data-sources/postgres_roles#attributes DataDatabricksPostgresRoles#attributes}.
	Attributes *DataDatabricksPostgresRolesRolesSpecAttributes `field:"optional" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/data-sources/postgres_roles#auth_method DataDatabricksPostgresRoles#auth_method}.
	AuthMethod *string `field:"optional" json:"authMethod" yaml:"authMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/data-sources/postgres_roles#identity_type DataDatabricksPostgresRoles#identity_type}.
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/data-sources/postgres_roles#membership_roles DataDatabricksPostgresRoles#membership_roles}.
	MembershipRoles *[]*string `field:"optional" json:"membershipRoles" yaml:"membershipRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/data-sources/postgres_roles#postgres_role DataDatabricksPostgresRoles#postgres_role}.
	PostgresRole *string `field:"optional" json:"postgresRole" yaml:"postgresRole"`
}

