// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTriggersSqlCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/job#sql_query_id Job#sql_query_id}.
	SqlQueryId *string `field:"required" json:"sqlQueryId" yaml:"sqlQueryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/job#warehouse_id Job#warehouse_id}.
	WarehouseId *string `field:"required" json:"warehouseId" yaml:"warehouseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/job#trigger_mode Job#trigger_mode}.
	TriggerMode *string `field:"optional" json:"triggerMode" yaml:"triggerMode"`
}

