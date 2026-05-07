// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappspaces


type DataDatabricksAppSpacesSpacesResourcesUcSecurable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/app_spaces#permission DataDatabricksAppSpaces#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/app_spaces#securable_full_name DataDatabricksAppSpaces#securable_full_name}.
	SecurableFullName *string `field:"required" json:"securableFullName" yaml:"securableFullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/app_spaces#securable_type DataDatabricksAppSpaces#securable_type}.
	SecurableType *string `field:"required" json:"securableType" yaml:"securableType"`
}

