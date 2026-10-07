// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobEnvironmentVariablesSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/resources/job#files Job#files}.
	Files *[]*string `field:"optional" json:"files" yaml:"files"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/resources/job#variables Job#variables}.
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
}

