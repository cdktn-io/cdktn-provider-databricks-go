// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mlflowexperiment


type MlflowExperimentTraceLocationUcTraceLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/mlflow_experiment#catalog MlflowExperiment#catalog}.
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/mlflow_experiment#schema MlflowExperiment#schema}.
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/mlflow_experiment#effective_table_prefix MlflowExperiment#effective_table_prefix}.
	EffectiveTablePrefix *string `field:"optional" json:"effectiveTablePrefix" yaml:"effectiveTablePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/mlflow_experiment#table_prefix MlflowExperiment#table_prefix}.
	TablePrefix *string `field:"optional" json:"tablePrefix" yaml:"tablePrefix"`
}

