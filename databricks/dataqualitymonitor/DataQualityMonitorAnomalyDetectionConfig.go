// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataqualitymonitor


type DataQualityMonitorAnomalyDetectionConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.142.0/docs/resources/data_quality_monitor#excluded_table_full_names DataQualityMonitor#excluded_table_full_names}.
	ExcludedTableFullNames *[]*string `field:"optional" json:"excludedTableFullNames" yaml:"excludedTableFullNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.142.0/docs/resources/data_quality_monitor#notification_settings DataQualityMonitor#notification_settings}.
	NotificationSettings *DataQualityMonitorAnomalyDetectionConfigNotificationSettings `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
}

