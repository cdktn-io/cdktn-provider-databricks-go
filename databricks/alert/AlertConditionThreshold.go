// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package alert


type AlertConditionThreshold struct {
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/alert#value Alert#value}
	Value *AlertConditionThresholdValue `field:"required" json:"value" yaml:"value"`
}

