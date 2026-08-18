// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints


type DataDatabricksServingEndpointsEndpointsTelemetryConfigTableNames struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/serving_endpoints#annotations_table DataDatabricksServingEndpoints#annotations_table}.
	AnnotationsTable *string `field:"optional" json:"annotationsTable" yaml:"annotationsTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/serving_endpoints#logs_table DataDatabricksServingEndpoints#logs_table}.
	LogsTable *string `field:"optional" json:"logsTable" yaml:"logsTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/serving_endpoints#metrics_table DataDatabricksServingEndpoints#metrics_table}.
	MetricsTable *string `field:"optional" json:"metricsTable" yaml:"metricsTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/serving_endpoints#traces_table DataDatabricksServingEndpoints#traces_table}.
	TracesTable *string `field:"optional" json:"tracesTable" yaml:"tracesTable"`
}

