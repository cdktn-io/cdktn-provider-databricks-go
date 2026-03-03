// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappsspace


type DataDatabricksAppsSpaceResourcesUcSecurable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#permission DataDatabricksAppsSpace#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#securable_full_name DataDatabricksAppsSpace#securable_full_name}.
	SecurableFullName *string `field:"required" json:"securableFullName" yaml:"securableFullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/data-sources/apps_space#securable_type DataDatabricksAppsSpace#securable_type}.
	SecurableType *string `field:"required" json:"securableType" yaml:"securableType"`
}

