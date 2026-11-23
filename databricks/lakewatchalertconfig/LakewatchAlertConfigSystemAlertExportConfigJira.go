// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchalertconfig


type LakewatchAlertConfigSystemAlertExportConfigJira struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_alert_config#api_token_uc_secret LakewatchAlertConfig#api_token_uc_secret}.
	ApiTokenUcSecret *string `field:"required" json:"apiTokenUcSecret" yaml:"apiTokenUcSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_alert_config#default_issue_type LakewatchAlertConfig#default_issue_type}.
	DefaultIssueType *string `field:"required" json:"defaultIssueType" yaml:"defaultIssueType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_alert_config#instance_url LakewatchAlertConfig#instance_url}.
	InstanceUrl *string `field:"required" json:"instanceUrl" yaml:"instanceUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_alert_config#project_key LakewatchAlertConfig#project_key}.
	ProjectKey *string `field:"required" json:"projectKey" yaml:"projectKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_alert_config#user_email LakewatchAlertConfig#user_email}.
	UserEmail *string `field:"required" json:"userEmail" yaml:"userEmail"`
}

