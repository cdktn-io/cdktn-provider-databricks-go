// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskAiFunctionTaskVolume struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/job#volume_path Job#volume_path}.
	VolumePath *string `field:"required" json:"volumePath" yaml:"volumePath"`
}

