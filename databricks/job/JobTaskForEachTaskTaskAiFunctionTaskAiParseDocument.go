// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskAiFunctionTaskAiParseDocument struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/job#version Job#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/job#description_element_types Job#description_element_types}.
	DescriptionElementTypes *[]*string `field:"optional" json:"descriptionElementTypes" yaml:"descriptionElementTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/job#image_output_path Job#image_output_path}.
	ImageOutputPath *string `field:"optional" json:"imageOutputPath" yaml:"imageOutputPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/job#page_range Job#page_range}.
	PageRange *string `field:"optional" json:"pageRange" yaml:"pageRange"`
}

