// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTriggersContinuous struct {
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.132.0/docs/resources/job#maintenance_window Job#maintenance_window}
	MaintenanceWindow *JobTriggersContinuousMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.132.0/docs/resources/job#task_retry_mode Job#task_retry_mode}.
	TaskRetryMode *string `field:"optional" json:"taskRetryMode" yaml:"taskRetryMode"`
}

