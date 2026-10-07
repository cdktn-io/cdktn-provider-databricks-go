// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobEnvironmentVariables struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/resources/job#environment_variables_key Job#environment_variables_key}.
	EnvironmentVariablesKey *string `field:"required" json:"environmentVariablesKey" yaml:"environmentVariablesKey"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/resources/job#spec Job#spec}
	Spec *JobEnvironmentVariablesSpec `field:"optional" json:"spec" yaml:"spec"`
}

