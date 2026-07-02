// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsReportTableConfigurationAutoFullRefreshPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/pipeline#enabled Pipeline#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/pipeline#min_interval_hours Pipeline#min_interval_hours}.
	MinIntervalHours *float64 `field:"optional" json:"minIntervalHours" yaml:"minIntervalHours"`
}

