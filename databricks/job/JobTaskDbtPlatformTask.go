// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskDbtPlatformTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/job#connection_resource_name Job#connection_resource_name}.
	ConnectionResourceName *string `field:"optional" json:"connectionResourceName" yaml:"connectionResourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/job#dbt_platform_job_id Job#dbt_platform_job_id}.
	DbtPlatformJobId *string `field:"optional" json:"dbtPlatformJobId" yaml:"dbtPlatformJobId"`
}

