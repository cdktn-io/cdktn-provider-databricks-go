// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickstagpolicies


type DataDatabricksTagPoliciesTagPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/tag_policies#tag_key DataDatabricksTagPolicies#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/tag_policies#provider_config DataDatabricksTagPolicies#provider_config}.
	ProviderConfig *DataDatabricksTagPoliciesTagPoliciesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

