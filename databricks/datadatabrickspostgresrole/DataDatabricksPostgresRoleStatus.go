// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresrole


type DataDatabricksPostgresRoleStatus struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/postgres_role#attributes DataDatabricksPostgresRole#attributes}.
	Attributes *DataDatabricksPostgresRoleStatusAttributes `field:"optional" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/postgres_role#auth_method DataDatabricksPostgresRole#auth_method}.
	AuthMethod *string `field:"optional" json:"authMethod" yaml:"authMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/postgres_role#identity_type DataDatabricksPostgresRole#identity_type}.
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/postgres_role#membership_roles DataDatabricksPostgresRole#membership_roles}.
	MembershipRoles *[]*string `field:"optional" json:"membershipRoles" yaml:"membershipRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/postgres_role#postgres_role DataDatabricksPostgresRole#postgres_role}.
	PostgresRole *string `field:"optional" json:"postgresRole" yaml:"postgresRole"`
}

