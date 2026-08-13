// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDashboardTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/job#dashboard_id DataDatabricksJob#dashboard_id}.
	DashboardId *string `field:"optional" json:"dashboardId" yaml:"dashboardId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/job#filters DataDatabricksJob#filters}.
	Filters *map[string]*string `field:"optional" json:"filters" yaml:"filters"`
	// subscription block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/job#subscription DataDatabricksJob#subscription}
	Subscription *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDashboardTaskSubscription `field:"optional" json:"subscription" yaml:"subscription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/job#warehouse_id DataDatabricksJob#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

