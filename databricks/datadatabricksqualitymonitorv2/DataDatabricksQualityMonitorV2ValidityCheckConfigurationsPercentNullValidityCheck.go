// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksqualitymonitorv2


type DataDatabricksQualityMonitorV2ValidityCheckConfigurationsPercentNullValidityCheck struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/quality_monitor_v2#column_names DataDatabricksQualityMonitorV2#column_names}.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/quality_monitor_v2#upper_bound DataDatabricksQualityMonitorV2#upper_bound}.
	UpperBound *float64 `field:"optional" json:"upperBound" yaml:"upperBound"`
}

