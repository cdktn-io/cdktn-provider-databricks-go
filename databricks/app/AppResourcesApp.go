// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package app


type AppResourcesApp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/app#name App#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/app#permission App#permission}.
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

