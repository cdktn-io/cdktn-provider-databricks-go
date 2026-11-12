// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchdatasource


type LakewatchDatasourceAutoloaderEnvelope struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.142.0/docs/resources/lakewatch_datasource#cloud_watch_logs LakewatchDatasource#cloud_watch_logs}.
	CloudWatchLogs *LakewatchDatasourceAutoloaderEnvelopeCloudWatchLogs `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.142.0/docs/resources/lakewatch_datasource#json_array LakewatchDatasource#json_array}.
	JsonArray *LakewatchDatasourceAutoloaderEnvelopeJsonArray `field:"optional" json:"jsonArray" yaml:"jsonArray"`
}

