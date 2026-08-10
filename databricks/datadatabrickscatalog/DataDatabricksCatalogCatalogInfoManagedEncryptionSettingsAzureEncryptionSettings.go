// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscatalog


type DataDatabricksCatalogCatalogInfoManagedEncryptionSettingsAzureEncryptionSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/catalog#azure_tenant_id DataDatabricksCatalog#azure_tenant_id}.
	AzureTenantId *string `field:"required" json:"azureTenantId" yaml:"azureTenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/catalog#azure_cmk_access_connector_id DataDatabricksCatalog#azure_cmk_access_connector_id}.
	AzureCmkAccessConnectorId *string `field:"optional" json:"azureCmkAccessConnectorId" yaml:"azureCmkAccessConnectorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/catalog#azure_cmk_managed_identity_id DataDatabricksCatalog#azure_cmk_managed_identity_id}.
	AzureCmkManagedIdentityId *string `field:"optional" json:"azureCmkManagedIdentityId" yaml:"azureCmkManagedIdentityId"`
}

