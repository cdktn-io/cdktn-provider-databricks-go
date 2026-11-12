// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchalertconfig


type LakewatchAlertConfigSystemAlertExportConfigTines struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.142.0/docs/resources/lakewatch_alert_config#webhook_secret_uc_secret LakewatchAlertConfig#webhook_secret_uc_secret}.
	WebhookSecretUcSecret *string `field:"required" json:"webhookSecretUcSecret" yaml:"webhookSecretUcSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.142.0/docs/resources/lakewatch_alert_config#webhook_uri LakewatchAlertConfig#webhook_uri}.
	WebhookUri *string `field:"required" json:"webhookUri" yaml:"webhookUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.142.0/docs/resources/lakewatch_alert_config#export_updates LakewatchAlertConfig#export_updates}.
	ExportUpdates interface{} `field:"optional" json:"exportUpdates" yaml:"exportUpdates"`
}

