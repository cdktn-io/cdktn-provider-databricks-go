// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskSparkPythonTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/job#python_file Job#python_file}.
	PythonFile *string `field:"required" json:"pythonFile" yaml:"pythonFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/job#parameters Job#parameters}.
	Parameters *[]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/job#source Job#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

