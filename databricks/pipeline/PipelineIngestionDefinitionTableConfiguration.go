// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionTableConfiguration struct {
	// auto_full_refresh_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/pipeline#auto_full_refresh_policy Pipeline#auto_full_refresh_policy}
	AutoFullRefreshPolicy *PipelineIngestionDefinitionTableConfigurationAutoFullRefreshPolicy `field:"optional" json:"autoFullRefreshPolicy" yaml:"autoFullRefreshPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/pipeline#exclude_columns Pipeline#exclude_columns}.
	ExcludeColumns *[]*string `field:"optional" json:"excludeColumns" yaml:"excludeColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/pipeline#include_columns Pipeline#include_columns}.
	IncludeColumns *[]*string `field:"optional" json:"includeColumns" yaml:"includeColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/pipeline#primary_keys Pipeline#primary_keys}.
	PrimaryKeys *[]*string `field:"optional" json:"primaryKeys" yaml:"primaryKeys"`
	// query_based_connector_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/pipeline#query_based_connector_config Pipeline#query_based_connector_config}
	QueryBasedConnectorConfig *PipelineIngestionDefinitionTableConfigurationQueryBasedConnectorConfig `field:"optional" json:"queryBasedConnectorConfig" yaml:"queryBasedConnectorConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/pipeline#row_filter Pipeline#row_filter}.
	RowFilter *string `field:"optional" json:"rowFilter" yaml:"rowFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/pipeline#salesforce_include_formula_fields Pipeline#salesforce_include_formula_fields}.
	SalesforceIncludeFormulaFields interface{} `field:"optional" json:"salesforceIncludeFormulaFields" yaml:"salesforceIncludeFormulaFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/pipeline#scd_type Pipeline#scd_type}.
	ScdType *string `field:"optional" json:"scdType" yaml:"scdType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/pipeline#sequence_by Pipeline#sequence_by}.
	SequenceBy *[]*string `field:"optional" json:"sequenceBy" yaml:"sequenceBy"`
	// workday_report_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/pipeline#workday_report_parameters Pipeline#workday_report_parameters}
	WorkdayReportParameters *PipelineIngestionDefinitionTableConfigurationWorkdayReportParameters `field:"optional" json:"workdayReportParameters" yaml:"workdayReportParameters"`
}

