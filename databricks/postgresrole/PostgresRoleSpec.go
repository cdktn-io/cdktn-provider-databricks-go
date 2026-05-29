// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresrole


type PostgresRoleSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/postgres_role#attributes PostgresRole#attributes}.
	Attributes *PostgresRoleSpecAttributes `field:"optional" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/postgres_role#auth_method PostgresRole#auth_method}.
	AuthMethod *string `field:"optional" json:"authMethod" yaml:"authMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/postgres_role#identity_type PostgresRole#identity_type}.
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/postgres_role#membership_roles PostgresRole#membership_roles}.
	MembershipRoles *[]*string `field:"optional" json:"membershipRoles" yaml:"membershipRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/postgres_role#postgres_role PostgresRole#postgres_role}.
	PostgresRole *string `field:"optional" json:"postgresRole" yaml:"postgresRole"`
}

