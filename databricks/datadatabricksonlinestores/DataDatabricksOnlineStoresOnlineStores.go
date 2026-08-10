// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksonlinestores


type DataDatabricksOnlineStoresOnlineStores struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/online_stores#name DataDatabricksOnlineStores#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/online_stores#provider_config DataDatabricksOnlineStores#provider_config}.
	ProviderConfig *DataDatabricksOnlineStoresOnlineStoresProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

