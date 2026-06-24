// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapps


type DataDatabricksAppsAppTelemetryExportDestinationsUnityCatalog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/apps#logs_table DataDatabricksApps#logs_table}.
	LogsTable *string `field:"required" json:"logsTable" yaml:"logsTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/apps#metrics_table DataDatabricksApps#metrics_table}.
	MetricsTable *string `field:"required" json:"metricsTable" yaml:"metricsTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/data-sources/apps#traces_table DataDatabricksApps#traces_table}.
	TracesTable *string `field:"required" json:"tracesTable" yaml:"tracesTable"`
}

