// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapp


type DataDatabricksAppAppTelemetryExportDestinationsUnityCatalog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app#logs_table DataDatabricksApp#logs_table}.
	LogsTable *string `field:"required" json:"logsTable" yaml:"logsTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app#metrics_table DataDatabricksApp#metrics_table}.
	MetricsTable *string `field:"required" json:"metricsTable" yaml:"metricsTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app#traces_table DataDatabricksApp#traces_table}.
	TracesTable *string `field:"required" json:"tracesTable" yaml:"tracesTable"`
}

