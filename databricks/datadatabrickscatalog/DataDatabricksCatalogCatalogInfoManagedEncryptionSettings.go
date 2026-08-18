// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscatalog


type DataDatabricksCatalogCatalogInfoManagedEncryptionSettings struct {
	// azure_encryption_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/catalog#azure_encryption_settings DataDatabricksCatalog#azure_encryption_settings}
	AzureEncryptionSettings *DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettings `field:"optional" json:"azureEncryptionSettings" yaml:"azureEncryptionSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/catalog#azure_key_vault_key_id DataDatabricksCatalog#azure_key_vault_key_id}.
	AzureKeyVaultKeyId *string `field:"optional" json:"azureKeyVaultKeyId" yaml:"azureKeyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/data-sources/catalog#customer_managed_key_id DataDatabricksCatalog#customer_managed_key_id}.
	CustomerManagedKeyId *string `field:"optional" json:"customerManagedKeyId" yaml:"customerManagedKeyId"`
}

