// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappsspaces


type DataDatabricksAppsSpacesSpaces struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_spaces#name DataDatabricksAppsSpaces#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_spaces#provider_config DataDatabricksAppsSpaces#provider_config}.
	ProviderConfig *DataDatabricksAppsSpacesSpacesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

