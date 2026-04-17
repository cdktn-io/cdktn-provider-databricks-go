// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksexternalmetadatas


type DataDatabricksExternalMetadatasExternalMetadata struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/data-sources/external_metadatas#name DataDatabricksExternalMetadatas#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/data-sources/external_metadatas#provider_config DataDatabricksExternalMetadatas#provider_config}.
	ProviderConfig *DataDatabricksExternalMetadatasExternalMetadataProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

