// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package app


type AppActiveDeploymentGitSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/app#branch App#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/app#commit App#commit}.
	Commit *string `field:"optional" json:"commit" yaml:"commit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/app#source_code_path App#source_code_path}.
	SourceCodePath *string `field:"optional" json:"sourceCodePath" yaml:"sourceCodePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/app#tag App#tag}.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

