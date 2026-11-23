// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchalertconfig


type DataDatabricksLakewatchAlertConfigSystemAlertExportConfigSlack struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_alert_config#bot_token_uc_secret DataDatabricksLakewatchAlertConfig#bot_token_uc_secret}.
	BotTokenUcSecret *string `field:"required" json:"botTokenUcSecret" yaml:"botTokenUcSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_alert_config#channel DataDatabricksLakewatchAlertConfig#channel}.
	Channel *string `field:"required" json:"channel" yaml:"channel"`
}

