// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalog


type CatalogManagedEncryptionSettingsAzureEncryptionSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/catalog#azure_tenant_id Catalog#azure_tenant_id}.
	AzureTenantId *string `field:"required" json:"azureTenantId" yaml:"azureTenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/catalog#azure_cmk_access_connector_id Catalog#azure_cmk_access_connector_id}.
	AzureCmkAccessConnectorId *string `field:"optional" json:"azureCmkAccessConnectorId" yaml:"azureCmkAccessConnectorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/catalog#azure_cmk_managed_identity_id Catalog#azure_cmk_managed_identity_id}.
	AzureCmkManagedIdentityId *string `field:"optional" json:"azureCmkManagedIdentityId" yaml:"azureCmkManagedIdentityId"`
}

