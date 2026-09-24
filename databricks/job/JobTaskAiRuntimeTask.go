// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskAiRuntimeTask struct {
	// deployments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/job#deployments Job#deployments}
	Deployments interface{} `field:"required" json:"deployments" yaml:"deployments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/job#experiment Job#experiment}.
	Experiment *string `field:"required" json:"experiment" yaml:"experiment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/job#code_source_path Job#code_source_path}.
	CodeSourcePath *string `field:"optional" json:"codeSourcePath" yaml:"codeSourcePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/job#docker_image_url Job#docker_image_url}.
	DockerImageUrl *string `field:"optional" json:"dockerImageUrl" yaml:"dockerImageUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/job#mlflow_artifact_location Job#mlflow_artifact_location}.
	MlflowArtifactLocation *string `field:"optional" json:"mlflowArtifactLocation" yaml:"mlflowArtifactLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/job#mlflow_experiment_directory Job#mlflow_experiment_directory}.
	MlflowExperimentDirectory *string `field:"optional" json:"mlflowExperimentDirectory" yaml:"mlflowExperimentDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/job#mlflow_run Job#mlflow_run}.
	MlflowRun *string `field:"optional" json:"mlflowRun" yaml:"mlflowRun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/job#priority_class Job#priority_class}.
	PriorityClass *string `field:"optional" json:"priorityClass" yaml:"priorityClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/job#unity_catalog_image_path Job#unity_catalog_image_path}.
	UnityCatalogImagePath *string `field:"optional" json:"unityCatalogImagePath" yaml:"unityCatalogImagePath"`
}

