// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptions struct {
	// custom_report_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/pipeline#custom_report_options Pipeline#custom_report_options}
	CustomReportOptions *PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptionsCustomReportOptions `field:"optional" json:"customReportOptions" yaml:"customReportOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/pipeline#data_level Pipeline#data_level}.
	DataLevel *string `field:"optional" json:"dataLevel" yaml:"dataLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/pipeline#dimensions Pipeline#dimensions}.
	Dimensions *[]*string `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/pipeline#lookback_window_days Pipeline#lookback_window_days}.
	LookbackWindowDays *float64 `field:"optional" json:"lookbackWindowDays" yaml:"lookbackWindowDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/pipeline#metrics Pipeline#metrics}.
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/pipeline#query_lifetime Pipeline#query_lifetime}.
	QueryLifetime interface{} `field:"optional" json:"queryLifetime" yaml:"queryLifetime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/pipeline#report_type Pipeline#report_type}.
	ReportType *string `field:"optional" json:"reportType" yaml:"reportType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/pipeline#sync_start_date Pipeline#sync_start_date}.
	SyncStartDate *string `field:"optional" json:"syncStartDate" yaml:"syncStartDate"`
}

