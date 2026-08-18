// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionSourceConfigurations struct {
	// catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/pipeline#catalog Pipeline#catalog}
	Catalog *PipelineIngestionDefinitionSourceConfigurationsCatalog `field:"optional" json:"catalog" yaml:"catalog"`
	// google_ads_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/pipeline#google_ads_config Pipeline#google_ads_config}
	GoogleAdsConfig *PipelineIngestionDefinitionSourceConfigurationsGoogleAdsConfig `field:"optional" json:"googleAdsConfig" yaml:"googleAdsConfig"`
}

