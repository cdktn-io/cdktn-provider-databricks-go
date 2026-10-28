// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspolicyinfo


type DataDatabricksPolicyInfoGovernedTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/policy_info#tag_key DataDatabricksPolicyInfo#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/policy_info#tag_value DataDatabricksPolicyInfo#tag_value}.
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
}

