// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksmasonmanagedmemorystores


type DataDatabricksMasonManagedMemoryStoresManagedMemoryStores struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/data-sources/mason_managed_memory_stores#name DataDatabricksMasonManagedMemoryStores#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/data-sources/mason_managed_memory_stores#provider_config DataDatabricksMasonManagedMemoryStores#provider_config}.
	ProviderConfig *DataDatabricksMasonManagedMemoryStoresManagedMemoryStoresProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

