// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksqualitymonitorsv2


type DataDatabricksQualityMonitorsV2QualityMonitorsValidityCheckConfigurationsRangeValidityCheck struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/quality_monitors_v2#column_names DataDatabricksQualityMonitorsV2#column_names}.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/quality_monitors_v2#lower_bound DataDatabricksQualityMonitorsV2#lower_bound}.
	LowerBound *float64 `field:"optional" json:"lowerBound" yaml:"lowerBound"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/quality_monitors_v2#upper_bound DataDatabricksQualityMonitorsV2#upper_bound}.
	UpperBound *float64 `field:"optional" json:"upperBound" yaml:"upperBound"`
}

