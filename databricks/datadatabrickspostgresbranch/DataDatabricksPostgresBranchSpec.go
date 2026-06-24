// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresbranch


type DataDatabricksPostgresBranchSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/postgres_branch#expire_time DataDatabricksPostgresBranch#expire_time}.
	ExpireTime *string `field:"optional" json:"expireTime" yaml:"expireTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/postgres_branch#is_protected DataDatabricksPostgresBranch#is_protected}.
	IsProtected interface{} `field:"optional" json:"isProtected" yaml:"isProtected"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/postgres_branch#no_expiry DataDatabricksPostgresBranch#no_expiry}.
	NoExpiry interface{} `field:"optional" json:"noExpiry" yaml:"noExpiry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/postgres_branch#source_branch DataDatabricksPostgresBranch#source_branch}.
	SourceBranch *string `field:"optional" json:"sourceBranch" yaml:"sourceBranch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/postgres_branch#source_branch_lsn DataDatabricksPostgresBranch#source_branch_lsn}.
	SourceBranchLsn *string `field:"optional" json:"sourceBranchLsn" yaml:"sourceBranchLsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/postgres_branch#source_branch_time DataDatabricksPostgresBranch#source_branch_time}.
	SourceBranchTime *string `field:"optional" json:"sourceBranchTime" yaml:"sourceBranchTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/postgres_branch#ttl DataDatabricksPostgresBranch#ttl}.
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

