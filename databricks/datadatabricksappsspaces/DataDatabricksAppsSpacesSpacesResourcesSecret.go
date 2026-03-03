// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappsspaces


type DataDatabricksAppsSpacesSpacesResourcesSecret struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_spaces#key DataDatabricksAppsSpaces#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_spaces#permission DataDatabricksAppsSpaces#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_spaces#scope DataDatabricksAppsSpaces#scope}.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

