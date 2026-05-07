// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mlflowexperiment


type MlflowExperimentTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/mlflow_experiment#key MlflowExperiment#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/mlflow_experiment#value MlflowExperiment#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

