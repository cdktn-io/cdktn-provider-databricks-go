// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budget


type BudgetFilterTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/budget#key Budget#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/budget#value Budget#value}
	Value *BudgetFilterTagsValue `field:"optional" json:"value" yaml:"value"`
}

