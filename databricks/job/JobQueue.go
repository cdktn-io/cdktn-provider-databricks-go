// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobQueue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.111.0/docs/resources/job#enabled Job#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

