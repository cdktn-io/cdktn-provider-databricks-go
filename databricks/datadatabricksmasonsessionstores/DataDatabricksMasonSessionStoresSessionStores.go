// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksmasonsessionstores


type DataDatabricksMasonSessionStoresSessionStores struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/data-sources/mason_session_stores#name DataDatabricksMasonSessionStores#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/data-sources/mason_session_stores#provider_config DataDatabricksMasonSessionStores#provider_config}.
	ProviderConfig *DataDatabricksMasonSessionStoresSessionStoresProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

