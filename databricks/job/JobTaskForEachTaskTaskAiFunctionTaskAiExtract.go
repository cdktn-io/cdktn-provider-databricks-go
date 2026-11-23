// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskAiFunctionTaskAiExtract struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/job#version Job#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/job#enable_citations Job#enable_citations}.
	EnableCitations interface{} `field:"optional" json:"enableCitations" yaml:"enableCitations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/job#enable_confidence_scores Job#enable_confidence_scores}.
	EnableConfidenceScores interface{} `field:"optional" json:"enableConfidenceScores" yaml:"enableConfidenceScores"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/job#instructions Job#instructions}.
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/job#mode Job#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/job#schema Job#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

