// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskAiRuntimeTaskDeployments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/job#command_path Job#command_path}.
	CommandPath *string `field:"required" json:"commandPath" yaml:"commandPath"`
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/job#compute Job#compute}
	Compute *JobTaskAiRuntimeTaskDeploymentsCompute `field:"required" json:"compute" yaml:"compute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/job#name Job#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

