// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package share


type ShareObjectPartition struct {
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/share#value Share#value}
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

