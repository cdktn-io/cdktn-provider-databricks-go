// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchalertconfig


type DataDatabricksLakewatchAlertConfigSystemAlertExportConfigTines struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_alert_config#webhook_secret_uc_secret DataDatabricksLakewatchAlertConfig#webhook_secret_uc_secret}.
	WebhookSecretUcSecret *string `field:"required" json:"webhookSecretUcSecret" yaml:"webhookSecretUcSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_alert_config#webhook_uri DataDatabricksLakewatchAlertConfig#webhook_uri}.
	WebhookUri *string `field:"required" json:"webhookUri" yaml:"webhookUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_alert_config#export_updates DataDatabricksLakewatchAlertConfig#export_updates}.
	ExportUpdates interface{} `field:"optional" json:"exportUpdates" yaml:"exportUpdates"`
}

