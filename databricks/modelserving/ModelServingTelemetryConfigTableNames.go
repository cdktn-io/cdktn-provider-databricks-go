// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingTelemetryConfigTableNames struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/model_serving#annotations_table ModelServing#annotations_table}.
	AnnotationsTable *string `field:"optional" json:"annotationsTable" yaml:"annotationsTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/model_serving#logs_table ModelServing#logs_table}.
	LogsTable *string `field:"optional" json:"logsTable" yaml:"logsTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/model_serving#metrics_table ModelServing#metrics_table}.
	MetricsTable *string `field:"optional" json:"metricsTable" yaml:"metricsTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/model_serving#traces_table ModelServing#traces_table}.
	TracesTable *string `field:"optional" json:"tracesTable" yaml:"tracesTable"`
}

