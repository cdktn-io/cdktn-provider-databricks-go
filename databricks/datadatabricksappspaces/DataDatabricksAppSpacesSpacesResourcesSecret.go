// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappspaces


type DataDatabricksAppSpacesSpacesResourcesSecret struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/app_spaces#key DataDatabricksAppSpaces#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/app_spaces#permission DataDatabricksAppSpaces#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/app_spaces#scope DataDatabricksAppSpaces#scope}.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

