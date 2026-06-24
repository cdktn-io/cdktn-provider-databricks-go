// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsTableConnectorOptions struct {
	// confluence_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/pipeline#confluence_options Pipeline#confluence_options}
	ConfluenceOptions *PipelineIngestionDefinitionObjectsTableConnectorOptionsConfluenceOptions `field:"optional" json:"confluenceOptions" yaml:"confluenceOptions"`
	// gdrive_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/pipeline#gdrive_options Pipeline#gdrive_options}
	GdriveOptions *PipelineIngestionDefinitionObjectsTableConnectorOptionsGdriveOptions `field:"optional" json:"gdriveOptions" yaml:"gdriveOptions"`
	// google_ads_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/pipeline#google_ads_options Pipeline#google_ads_options}
	GoogleAdsOptions *PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptions `field:"optional" json:"googleAdsOptions" yaml:"googleAdsOptions"`
	// jira_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/pipeline#jira_options Pipeline#jira_options}
	JiraOptions *PipelineIngestionDefinitionObjectsTableConnectorOptionsJiraOptions `field:"optional" json:"jiraOptions" yaml:"jiraOptions"`
	// kafka_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/pipeline#kafka_options Pipeline#kafka_options}
	KafkaOptions *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptions `field:"optional" json:"kafkaOptions" yaml:"kafkaOptions"`
	// meta_ads_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/pipeline#meta_ads_options Pipeline#meta_ads_options}
	MetaAdsOptions *PipelineIngestionDefinitionObjectsTableConnectorOptionsMetaAdsOptions `field:"optional" json:"metaAdsOptions" yaml:"metaAdsOptions"`
	// outlook_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/pipeline#outlook_options Pipeline#outlook_options}
	OutlookOptions *PipelineIngestionDefinitionObjectsTableConnectorOptionsOutlookOptions `field:"optional" json:"outlookOptions" yaml:"outlookOptions"`
	// sharepoint_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/pipeline#sharepoint_options Pipeline#sharepoint_options}
	SharepointOptions *PipelineIngestionDefinitionObjectsTableConnectorOptionsSharepointOptions `field:"optional" json:"sharepointOptions" yaml:"sharepointOptions"`
	// smartsheet_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/pipeline#smartsheet_options Pipeline#smartsheet_options}
	SmartsheetOptions *PipelineIngestionDefinitionObjectsTableConnectorOptionsSmartsheetOptions `field:"optional" json:"smartsheetOptions" yaml:"smartsheetOptions"`
	// tiktok_ads_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/pipeline#tiktok_ads_options Pipeline#tiktok_ads_options}
	TiktokAdsOptions *PipelineIngestionDefinitionObjectsTableConnectorOptionsTiktokAdsOptions `field:"optional" json:"tiktokAdsOptions" yaml:"tiktokAdsOptions"`
	// zendesk_support_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/pipeline#zendesk_support_options Pipeline#zendesk_support_options}
	ZendeskSupportOptions *PipelineIngestionDefinitionObjectsTableConnectorOptionsZendeskSupportOptions `field:"optional" json:"zendeskSupportOptions" yaml:"zendeskSupportOptions"`
}

