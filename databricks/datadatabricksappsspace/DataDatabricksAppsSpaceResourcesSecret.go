// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappsspace


type DataDatabricksAppsSpaceResourcesSecret struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#key DataDatabricksAppsSpace#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#permission DataDatabricksAppsSpace#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#scope DataDatabricksAppsSpace#scope}.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

