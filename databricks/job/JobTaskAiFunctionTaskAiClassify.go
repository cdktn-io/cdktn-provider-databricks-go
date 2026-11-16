// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskAiFunctionTaskAiClassify struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/job#version Job#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/job#enable_confidence_scores Job#enable_confidence_scores}.
	EnableConfidenceScores interface{} `field:"optional" json:"enableConfidenceScores" yaml:"enableConfidenceScores"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/job#enable_rationales Job#enable_rationales}.
	EnableRationales interface{} `field:"optional" json:"enableRationales" yaml:"enableRationales"`
	// inline_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/job#inline_labels Job#inline_labels}
	InlineLabels *JobTaskAiFunctionTaskAiClassifyInlineLabels `field:"optional" json:"inlineLabels" yaml:"inlineLabels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/job#instructions Job#instructions}.
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/job#multilabel Job#multilabel}.
	Multilabel interface{} `field:"optional" json:"multilabel" yaml:"multilabel"`
}

