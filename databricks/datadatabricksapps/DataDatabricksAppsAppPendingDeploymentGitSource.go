// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapps


type DataDatabricksAppsAppPendingDeploymentGitSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/apps#branch DataDatabricksApps#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/apps#commit DataDatabricksApps#commit}.
	Commit *string `field:"optional" json:"commit" yaml:"commit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/apps#source_code_path DataDatabricksApps#source_code_path}.
	SourceCodePath *string `field:"optional" json:"sourceCodePath" yaml:"sourceCodePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/data-sources/apps#tag DataDatabricksApps#tag}.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

