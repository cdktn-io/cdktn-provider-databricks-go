// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsSchemaConnectorOptions struct {
	// gdrive_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/pipeline#gdrive_options Pipeline#gdrive_options}
	GdriveOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptions `field:"optional" json:"gdriveOptions" yaml:"gdriveOptions"`
	// google_ads_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/pipeline#google_ads_options Pipeline#google_ads_options}
	GoogleAdsOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptions `field:"optional" json:"googleAdsOptions" yaml:"googleAdsOptions"`
	// sharepoint_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/pipeline#sharepoint_options Pipeline#sharepoint_options}
	SharepointOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSharepointOptions `field:"optional" json:"sharepointOptions" yaml:"sharepointOptions"`
	// tiktok_ads_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/pipeline#tiktok_ads_options Pipeline#tiktok_ads_options}
	TiktokAdsOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsTiktokAdsOptions `field:"optional" json:"tiktokAdsOptions" yaml:"tiktokAdsOptions"`
}

