// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapp


type DataDatabricksAppAppGitRepository struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/app#provider DataDatabricksApp#provider}.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/app#url DataDatabricksApp#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}

