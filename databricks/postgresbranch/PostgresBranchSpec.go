// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresbranch


type PostgresBranchSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/postgres_branch#expire_time PostgresBranch#expire_time}.
	ExpireTime *string `field:"optional" json:"expireTime" yaml:"expireTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/postgres_branch#is_protected PostgresBranch#is_protected}.
	IsProtected interface{} `field:"optional" json:"isProtected" yaml:"isProtected"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/postgres_branch#no_expiry PostgresBranch#no_expiry}.
	NoExpiry interface{} `field:"optional" json:"noExpiry" yaml:"noExpiry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/postgres_branch#source_branch PostgresBranch#source_branch}.
	SourceBranch *string `field:"optional" json:"sourceBranch" yaml:"sourceBranch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/postgres_branch#source_branch_lsn PostgresBranch#source_branch_lsn}.
	SourceBranchLsn *string `field:"optional" json:"sourceBranchLsn" yaml:"sourceBranchLsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/postgres_branch#source_branch_time PostgresBranch#source_branch_time}.
	SourceBranchTime *string `field:"optional" json:"sourceBranchTime" yaml:"sourceBranchTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/postgres_branch#ttl PostgresBranch#ttl}.
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

