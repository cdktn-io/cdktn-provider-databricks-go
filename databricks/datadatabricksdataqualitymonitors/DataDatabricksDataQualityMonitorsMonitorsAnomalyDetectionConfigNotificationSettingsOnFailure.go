// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitors


type DataDatabricksDataQualityMonitorsMonitorsAnomalyDetectionConfigNotificationSettingsOnFailure struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.142.0/docs/data-sources/data_quality_monitors#destination_ids DataDatabricksDataQualityMonitors#destination_ids}.
	DestinationIds *[]*string `field:"optional" json:"destinationIds" yaml:"destinationIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.142.0/docs/data-sources/data_quality_monitors#destination_user_ids DataDatabricksDataQualityMonitors#destination_user_ids}.
	DestinationUserIds *[]*float64 `field:"optional" json:"destinationUserIds" yaml:"destinationUserIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.142.0/docs/data-sources/data_quality_monitors#email_addresses DataDatabricksDataQualityMonitors#email_addresses}.
	EmailAddresses *[]*string `field:"optional" json:"emailAddresses" yaml:"emailAddresses"`
}

