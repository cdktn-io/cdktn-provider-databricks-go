// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitor


type DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/data_quality_monitor#on_failure DataDatabricksDataQualityMonitor#on_failure}.
	OnFailure *DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/data_quality_monitor#on_unhealthy_table DataDatabricksDataQualityMonitor#on_unhealthy_table}.
	OnUnhealthyTable *DataDatabricksDataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTable `field:"optional" json:"onUnhealthyTable" yaml:"onUnhealthyTable"`
}

