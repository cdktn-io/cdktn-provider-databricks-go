// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskAlertTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/job#alert_id Job#alert_id}.
	AlertId *string `field:"optional" json:"alertId" yaml:"alertId"`
	// subscribers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/job#subscribers Job#subscribers}
	Subscribers interface{} `field:"optional" json:"subscribers" yaml:"subscribers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/job#warehouse_id Job#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/job#workspace_path Job#workspace_path}.
	WorkspacePath *string `field:"optional" json:"workspacePath" yaml:"workspacePath"`
}

