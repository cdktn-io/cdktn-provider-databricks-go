// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataclassificationcatalogconfig


type DataClassificationCatalogConfigAutoTagConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/data_classification_catalog_config#auto_tagging_mode DataClassificationCatalogConfig#auto_tagging_mode}.
	AutoTaggingMode *string `field:"required" json:"autoTaggingMode" yaml:"autoTaggingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/data_classification_catalog_config#classification_tag DataClassificationCatalogConfig#classification_tag}.
	ClassificationTag *string `field:"required" json:"classificationTag" yaml:"classificationTag"`
}

