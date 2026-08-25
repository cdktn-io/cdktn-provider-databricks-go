// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptionsCustomReportOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#action_attribution_windows Pipeline#action_attribution_windows}.
	ActionAttributionWindows *[]*string `field:"optional" json:"actionAttributionWindows" yaml:"actionAttributionWindows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#action_breakdowns Pipeline#action_breakdowns}.
	ActionBreakdowns *[]*string `field:"optional" json:"actionBreakdowns" yaml:"actionBreakdowns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#action_report_time Pipeline#action_report_time}.
	ActionReportTime *string `field:"optional" json:"actionReportTime" yaml:"actionReportTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#breakdowns Pipeline#breakdowns}.
	Breakdowns *[]*string `field:"optional" json:"breakdowns" yaml:"breakdowns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#level Pipeline#level}.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#time_increment Pipeline#time_increment}.
	TimeIncrement *string `field:"optional" json:"timeIncrement" yaml:"timeIncrement"`
}

