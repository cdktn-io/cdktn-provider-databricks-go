// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package masonmanagedmemorystore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MasonManagedMemoryStoreConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/resources/mason_managed_memory_store#managed_memory_store_id MasonManagedMemoryStore#managed_memory_store_id}.
	ManagedMemoryStoreId *string `field:"required" json:"managedMemoryStoreId" yaml:"managedMemoryStoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/resources/mason_managed_memory_store#description MasonManagedMemoryStore#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/resources/mason_managed_memory_store#display_name MasonManagedMemoryStore#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.137.0/docs/resources/mason_managed_memory_store#provider_config MasonManagedMemoryStore#provider_config}.
	ProviderConfig *MasonManagedMemoryStoreProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

