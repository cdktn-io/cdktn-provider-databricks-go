// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package app


type AppActiveDeploymentGitSourceGitRepository struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/app#provider App#provider}.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/app#url App#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}

