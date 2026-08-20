// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskPythonOperatorTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/job#main Job#main}.
	Main *string `field:"optional" json:"main" yaml:"main"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/job#parameters Job#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

