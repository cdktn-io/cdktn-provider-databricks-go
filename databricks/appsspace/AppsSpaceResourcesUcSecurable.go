// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsspace


type AppsSpaceResourcesUcSecurable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/resources/apps_space#permission AppsSpace#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/resources/apps_space#securable_full_name AppsSpace#securable_full_name}.
	SecurableFullName *string `field:"required" json:"securableFullName" yaml:"securableFullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/resources/apps_space#securable_type AppsSpace#securable_type}.
	SecurableType *string `field:"required" json:"securableType" yaml:"securableType"`
}

