// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PostgresRoleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/postgres_role#parent PostgresRole#parent}.
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/postgres_role#provider_config PostgresRole#provider_config}.
	ProviderConfig *PostgresRoleProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/postgres_role#replace_existing PostgresRole#replace_existing}.
	ReplaceExisting interface{} `field:"optional" json:"replaceExisting" yaml:"replaceExisting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/postgres_role#role_id PostgresRole#role_id}.
	RoleId *string `field:"optional" json:"roleId" yaml:"roleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/postgres_role#spec PostgresRole#spec}.
	Spec *PostgresRoleSpec `field:"optional" json:"spec" yaml:"spec"`
}

