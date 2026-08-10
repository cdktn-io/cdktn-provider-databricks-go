// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksmlflowexperiment


type DataDatabricksMlflowExperimentTraceLocationUcTraceLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/mlflow_experiment#catalog DataDatabricksMlflowExperiment#catalog}.
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/mlflow_experiment#schema DataDatabricksMlflowExperiment#schema}.
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/mlflow_experiment#effective_table_prefix DataDatabricksMlflowExperiment#effective_table_prefix}.
	EffectiveTablePrefix *string `field:"optional" json:"effectiveTablePrefix" yaml:"effectiveTablePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/mlflow_experiment#table_prefix DataDatabricksMlflowExperiment#table_prefix}.
	TablePrefix *string `field:"optional" json:"tablePrefix" yaml:"tablePrefix"`
}

