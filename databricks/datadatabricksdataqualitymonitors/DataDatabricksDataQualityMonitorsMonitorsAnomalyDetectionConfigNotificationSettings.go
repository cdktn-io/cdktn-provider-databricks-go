// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitors


type DataDatabricksDataQualityMonitorsMonitorsAnomalyDetectionConfigNotificationSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/data_quality_monitors#on_failure DataDatabricksDataQualityMonitors#on_failure}.
	OnFailure *DataDatabricksDataQualityMonitorsMonitorsAnomalyDetectionConfigNotificationSettingsOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/data-sources/data_quality_monitors#on_unhealthy_table DataDatabricksDataQualityMonitors#on_unhealthy_table}.
	OnUnhealthyTable *DataDatabricksDataQualityMonitorsMonitorsAnomalyDetectionConfigNotificationSettingsOnUnhealthyTable `field:"optional" json:"onUnhealthyTable" yaml:"onUnhealthyTable"`
}

