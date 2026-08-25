// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspolicyinfo


type DataDatabricksPolicyInfoGrant struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/policy_info#privileges DataDatabricksPolicyInfo#privileges}.
	Privileges *[]*string `field:"required" json:"privileges" yaml:"privileges"`
}

