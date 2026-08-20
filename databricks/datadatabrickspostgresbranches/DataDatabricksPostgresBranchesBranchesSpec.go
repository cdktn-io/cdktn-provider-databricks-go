// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresbranches


type DataDatabricksPostgresBranchesBranchesSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/postgres_branches#expire_time DataDatabricksPostgresBranches#expire_time}.
	ExpireTime *string `field:"optional" json:"expireTime" yaml:"expireTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/postgres_branches#is_protected DataDatabricksPostgresBranches#is_protected}.
	IsProtected interface{} `field:"optional" json:"isProtected" yaml:"isProtected"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/postgres_branches#no_expiry DataDatabricksPostgresBranches#no_expiry}.
	NoExpiry interface{} `field:"optional" json:"noExpiry" yaml:"noExpiry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/postgres_branches#source_branch DataDatabricksPostgresBranches#source_branch}.
	SourceBranch *string `field:"optional" json:"sourceBranch" yaml:"sourceBranch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/postgres_branches#source_branch_lsn DataDatabricksPostgresBranches#source_branch_lsn}.
	SourceBranchLsn *string `field:"optional" json:"sourceBranchLsn" yaml:"sourceBranchLsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/postgres_branches#source_branch_time DataDatabricksPostgresBranches#source_branch_time}.
	SourceBranchTime *string `field:"optional" json:"sourceBranchTime" yaml:"sourceBranchTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/postgres_branches#ttl DataDatabricksPostgresBranches#ttl}.
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

