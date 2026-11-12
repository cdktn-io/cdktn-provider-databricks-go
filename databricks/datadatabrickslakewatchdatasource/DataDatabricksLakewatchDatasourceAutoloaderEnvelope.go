// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchdatasource


type DataDatabricksLakewatchDatasourceAutoloaderEnvelope struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.142.0/docs/data-sources/lakewatch_datasource#cloud_watch_logs DataDatabricksLakewatchDatasource#cloud_watch_logs}.
	CloudWatchLogs *DataDatabricksLakewatchDatasourceAutoloaderEnvelopeCloudWatchLogs `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.142.0/docs/data-sources/lakewatch_datasource#json_array DataDatabricksLakewatchDatasource#json_array}.
	JsonArray *DataDatabricksLakewatchDatasourceAutoloaderEnvelopeJsonArray `field:"optional" json:"jsonArray" yaml:"jsonArray"`
}

