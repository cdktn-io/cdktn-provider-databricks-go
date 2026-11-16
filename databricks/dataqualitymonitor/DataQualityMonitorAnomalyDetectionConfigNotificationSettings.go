// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataqualitymonitor


type DataQualityMonitorAnomalyDetectionConfigNotificationSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/data_quality_monitor#on_failure DataQualityMonitor#on_failure}.
	OnFailure *DataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/data_quality_monitor#on_unhealthy_table DataQualityMonitor#on_unhealthy_table}.
	OnUnhealthyTable *DataQualityMonitorAnomalyDetectionConfigNotificationSettingsOnUnhealthyTable `field:"optional" json:"onUnhealthyTable" yaml:"onUnhealthyTable"`
}

