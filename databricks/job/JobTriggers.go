// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTriggers struct {
	// continuous block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/job#continuous Job#continuous}
	Continuous *JobTriggersContinuous `field:"optional" json:"continuous" yaml:"continuous"`
	// file_arrival block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/job#file_arrival Job#file_arrival}
	FileArrival *JobTriggersFileArrival `field:"optional" json:"fileArrival" yaml:"fileArrival"`
	// model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/job#model Job#model}
	Model *JobTriggersModel `field:"optional" json:"model" yaml:"model"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/job#pause_status Job#pause_status}.
	PauseStatus *string `field:"optional" json:"pauseStatus" yaml:"pauseStatus"`
	// periodic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/job#periodic Job#periodic}
	Periodic *JobTriggersPeriodic `field:"optional" json:"periodic" yaml:"periodic"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/job#schedule Job#schedule}
	Schedule *JobTriggersSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// sql_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/job#sql_condition Job#sql_condition}
	SqlCondition *JobTriggersSqlCondition `field:"optional" json:"sqlCondition" yaml:"sqlCondition"`
	// table_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/job#table_update Job#table_update}
	TableUpdate *JobTriggersTableUpdate `field:"optional" json:"tableUpdate" yaml:"tableUpdate"`
}

