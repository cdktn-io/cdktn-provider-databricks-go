// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/pipeline#manager_account_id Pipeline#manager_account_id}.
	ManagerAccountId *string `field:"required" json:"managerAccountId" yaml:"managerAccountId"`
	// custom_report_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/pipeline#custom_report_options Pipeline#custom_report_options}
	CustomReportOptions *PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptions `field:"optional" json:"customReportOptions" yaml:"customReportOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/pipeline#lookback_window_days Pipeline#lookback_window_days}.
	LookbackWindowDays *float64 `field:"optional" json:"lookbackWindowDays" yaml:"lookbackWindowDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/pipeline#sync_start_date Pipeline#sync_start_date}.
	SyncStartDate *string `field:"optional" json:"syncStartDate" yaml:"syncStartDate"`
}

