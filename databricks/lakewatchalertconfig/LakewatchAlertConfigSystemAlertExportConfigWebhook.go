// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchalertconfig


type LakewatchAlertConfigSystemAlertExportConfigWebhook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/lakewatch_alert_config#export_updates LakewatchAlertConfig#export_updates}.
	ExportUpdates interface{} `field:"optional" json:"exportUpdates" yaml:"exportUpdates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/lakewatch_alert_config#url LakewatchAlertConfig#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

