// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitor


type DataDatabricksDataQualityMonitorDataProfilingConfigNotificationSettingsOnUnhealthyTable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/data_quality_monitor#destination_ids DataDatabricksDataQualityMonitor#destination_ids}.
	DestinationIds *[]*string `field:"optional" json:"destinationIds" yaml:"destinationIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/data_quality_monitor#destination_user_ids DataDatabricksDataQualityMonitor#destination_user_ids}.
	DestinationUserIds *[]*float64 `field:"optional" json:"destinationUserIds" yaml:"destinationUserIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/data_quality_monitor#email_addresses DataDatabricksDataQualityMonitor#email_addresses}.
	EmailAddresses *[]*string `field:"optional" json:"emailAddresses" yaml:"emailAddresses"`
}

