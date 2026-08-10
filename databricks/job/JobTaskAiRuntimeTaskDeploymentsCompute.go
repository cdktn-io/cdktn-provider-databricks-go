// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskAiRuntimeTaskDeploymentsCompute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/job#accelerator_count Job#accelerator_count}.
	AcceleratorCount *float64 `field:"required" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/job#accelerator_type Job#accelerator_type}.
	AcceleratorType *string `field:"required" json:"acceleratorType" yaml:"acceleratorType"`
}

