// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTriggersJobCompletion struct {
	// jobs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/job#jobs Job#jobs}
	Jobs interface{} `field:"required" json:"jobs" yaml:"jobs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/job#if_status Job#if_status}.
	IfStatus *string `field:"optional" json:"ifStatus" yaml:"ifStatus"`
}

