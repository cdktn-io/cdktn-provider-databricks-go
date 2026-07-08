// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresproject


type PostgresProjectInitialBranchSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/postgres_project#is_protected PostgresProject#is_protected}.
	IsProtected interface{} `field:"optional" json:"isProtected" yaml:"isProtected"`
}

