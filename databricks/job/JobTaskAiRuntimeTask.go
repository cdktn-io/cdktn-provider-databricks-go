// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskAiRuntimeTask struct {
	// deployments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/job#deployments Job#deployments}
	Deployments interface{} `field:"required" json:"deployments" yaml:"deployments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/job#experiment Job#experiment}.
	Experiment *string `field:"required" json:"experiment" yaml:"experiment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/job#code_source_path Job#code_source_path}.
	CodeSourcePath *string `field:"optional" json:"codeSourcePath" yaml:"codeSourcePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/job#mlflow_experiment_directory Job#mlflow_experiment_directory}.
	MlflowExperimentDirectory *string `field:"optional" json:"mlflowExperimentDirectory" yaml:"mlflowExperimentDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/job#mlflow_run Job#mlflow_run}.
	MlflowRun *string `field:"optional" json:"mlflowRun" yaml:"mlflowRun"`
}

