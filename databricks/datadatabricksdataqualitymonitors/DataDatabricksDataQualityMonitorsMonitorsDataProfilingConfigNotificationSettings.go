// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitors


type DataDatabricksDataQualityMonitorsMonitorsDataProfilingConfigNotificationSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/data_quality_monitors#on_failure DataDatabricksDataQualityMonitors#on_failure}.
	OnFailure *DataDatabricksDataQualityMonitorsMonitorsDataProfilingConfigNotificationSettingsOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/data_quality_monitors#on_unhealthy_table DataDatabricksDataQualityMonitors#on_unhealthy_table}.
	OnUnhealthyTable *DataDatabricksDataQualityMonitorsMonitorsDataProfilingConfigNotificationSettingsOnUnhealthyTable `field:"optional" json:"onUnhealthyTable" yaml:"onUnhealthyTable"`
}

