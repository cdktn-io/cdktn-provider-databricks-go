// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsDeployment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/job#kind DataDatabricksJob#kind}.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/job#deployment_id DataDatabricksJob#deployment_id}.
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/job#metadata_file_path DataDatabricksJob#metadata_file_path}.
	MetadataFilePath *string `field:"optional" json:"metadataFilePath" yaml:"metadataFilePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/job#version_id DataDatabricksJob#version_id}.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

