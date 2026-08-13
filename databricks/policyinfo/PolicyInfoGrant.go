// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package policyinfo


type PolicyInfoGrant struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/policy_info#privileges PolicyInfo#privileges}.
	Privileges *[]*string `field:"required" json:"privileges" yaml:"privileges"`
}

