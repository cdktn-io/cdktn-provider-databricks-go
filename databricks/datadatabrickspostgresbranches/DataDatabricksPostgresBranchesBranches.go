// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresbranches


type DataDatabricksPostgresBranchesBranches struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/postgres_branches#name DataDatabricksPostgresBranches#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/postgres_branches#provider_config DataDatabricksPostgresBranches#provider_config}.
	ProviderConfig *DataDatabricksPostgresBranchesBranchesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

