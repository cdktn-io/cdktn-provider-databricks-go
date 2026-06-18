// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappspace


type DataDatabricksAppSpaceResourcesUcSecurable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/app_space#permission DataDatabricksAppSpace#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/app_space#securable_full_name DataDatabricksAppSpace#securable_full_name}.
	SecurableFullName *string `field:"required" json:"securableFullName" yaml:"securableFullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/app_space#securable_type DataDatabricksAppSpace#securable_type}.
	SecurableType *string `field:"required" json:"securableType" yaml:"securableType"`
}

