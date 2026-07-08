// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresroles


type DataDatabricksPostgresRolesRoles struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/postgres_roles#name DataDatabricksPostgresRoles#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/postgres_roles#provider_config DataDatabricksPostgresRoles#provider_config}.
	ProviderConfig *DataDatabricksPostgresRolesRolesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

