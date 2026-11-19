// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchalertconfig


type LakewatchAlertConfigSecurityAlertExportConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/lakewatch_alert_config#enabled LakewatchAlertConfig#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/lakewatch_alert_config#jira LakewatchAlertConfig#jira}.
	Jira *LakewatchAlertConfigSecurityAlertExportConfigJira `field:"optional" json:"jira" yaml:"jira"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/lakewatch_alert_config#service_now LakewatchAlertConfig#service_now}.
	ServiceNow *LakewatchAlertConfigSecurityAlertExportConfigServiceNow `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/lakewatch_alert_config#slack LakewatchAlertConfig#slack}.
	Slack *LakewatchAlertConfigSecurityAlertExportConfigSlack `field:"optional" json:"slack" yaml:"slack"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/lakewatch_alert_config#tines LakewatchAlertConfig#tines}.
	Tines *LakewatchAlertConfigSecurityAlertExportConfigTines `field:"optional" json:"tines" yaml:"tines"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/lakewatch_alert_config#webhook LakewatchAlertConfig#webhook}.
	Webhook *LakewatchAlertConfigSecurityAlertExportConfigWebhook `field:"optional" json:"webhook" yaml:"webhook"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/lakewatch_alert_config#workspace_id LakewatchAlertConfig#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

