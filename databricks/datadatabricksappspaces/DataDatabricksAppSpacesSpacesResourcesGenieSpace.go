// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappspaces


type DataDatabricksAppSpacesSpacesResourcesGenieSpace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/app_spaces#name DataDatabricksAppSpaces#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/app_spaces#permission DataDatabricksAppSpaces#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/app_spaces#space_id DataDatabricksAppSpaces#space_id}.
	SpaceId *string `field:"required" json:"spaceId" yaml:"spaceId"`
}

