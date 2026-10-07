// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksmasonmanagedmemoryentries


type DataDatabricksMasonManagedMemoryEntriesManagedMemoryEntries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/data-sources/mason_managed_memory_entries#name DataDatabricksMasonManagedMemoryEntries#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/data-sources/mason_managed_memory_entries#provider_config DataDatabricksMasonManagedMemoryEntries#provider_config}.
	ProviderConfig *DataDatabricksMasonManagedMemoryEntriesManagedMemoryEntriesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

