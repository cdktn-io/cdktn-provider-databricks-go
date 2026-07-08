// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineDeployment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#kind Pipeline#kind}.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#deployment_id Pipeline#deployment_id}.
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#metadata_file_path Pipeline#metadata_file_path}.
	MetadataFilePath *string `field:"optional" json:"metadataFilePath" yaml:"metadataFilePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#version_id Pipeline#version_id}.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

