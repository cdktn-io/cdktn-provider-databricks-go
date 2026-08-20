// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalog


type CatalogManagedEncryptionSettings struct {
	// azure_encryption_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/catalog#azure_encryption_settings Catalog#azure_encryption_settings}
	AzureEncryptionSettings *CatalogManagedEncryptionSettingsAzureEncryptionSettings `field:"optional" json:"azureEncryptionSettings" yaml:"azureEncryptionSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/catalog#azure_key_vault_key_id Catalog#azure_key_vault_key_id}.
	AzureKeyVaultKeyId *string `field:"optional" json:"azureKeyVaultKeyId" yaml:"azureKeyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/catalog#customer_managed_key_id Catalog#customer_managed_key_id}.
	CustomerManagedKeyId *string `field:"optional" json:"customerManagedKeyId" yaml:"customerManagedKeyId"`
}

