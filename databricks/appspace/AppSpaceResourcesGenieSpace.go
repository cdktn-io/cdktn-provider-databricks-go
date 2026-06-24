// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appspace


type AppSpaceResourcesGenieSpace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/app_space#name AppSpace#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/app_space#permission AppSpace#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/app_space#space_id AppSpace#space_id}.
	SpaceId *string `field:"required" json:"spaceId" yaml:"spaceId"`
}

