// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package app


type AppResourcesGenieSpace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/app#name App#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/app#permission App#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/app#space_id App#space_id}.
	SpaceId *string `field:"required" json:"spaceId" yaml:"spaceId"`
}

