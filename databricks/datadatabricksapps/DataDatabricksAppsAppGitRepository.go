// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapps


type DataDatabricksAppsAppGitRepository struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/apps#provider DataDatabricksApps#provider}.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/apps#url DataDatabricksApps#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}

