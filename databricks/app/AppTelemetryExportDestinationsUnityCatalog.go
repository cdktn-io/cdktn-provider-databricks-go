// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package app


type AppTelemetryExportDestinationsUnityCatalog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/app#logs_table App#logs_table}.
	LogsTable *string `field:"required" json:"logsTable" yaml:"logsTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/app#metrics_table App#metrics_table}.
	MetricsTable *string `field:"required" json:"metricsTable" yaml:"metricsTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/app#traces_table App#traces_table}.
	TracesTable *string `field:"required" json:"tracesTable" yaml:"tracesTable"`
}

