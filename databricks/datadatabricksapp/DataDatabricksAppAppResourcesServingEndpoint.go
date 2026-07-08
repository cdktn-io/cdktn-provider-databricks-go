// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapp


type DataDatabricksAppAppResourcesServingEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/app#name DataDatabricksApp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/app#permission DataDatabricksApp#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

