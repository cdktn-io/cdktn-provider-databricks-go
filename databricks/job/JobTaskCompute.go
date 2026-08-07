// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskCompute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/job#hardware_accelerator Job#hardware_accelerator}.
	HardwareAccelerator *string `field:"optional" json:"hardwareAccelerator" yaml:"hardwareAccelerator"`
}

