// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappspace


type DataDatabricksAppSpaceResourcesServingEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/app_space#name DataDatabricksAppSpace#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/data-sources/app_space#permission DataDatabricksAppSpace#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

