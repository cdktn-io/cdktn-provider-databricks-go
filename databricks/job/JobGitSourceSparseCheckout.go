// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobGitSourceSparseCheckout struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/job#patterns Job#patterns}.
	Patterns *[]*string `field:"optional" json:"patterns" yaml:"patterns"`
}

