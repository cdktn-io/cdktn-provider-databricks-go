// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsSchemaConnectorOptions struct {
	// confluence_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#confluence_options Pipeline#confluence_options}
	ConfluenceOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsConfluenceOptions `field:"optional" json:"confluenceOptions" yaml:"confluenceOptions"`
	// gdrive_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#gdrive_options Pipeline#gdrive_options}
	GdriveOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptions `field:"optional" json:"gdriveOptions" yaml:"gdriveOptions"`
	// google_ads_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#google_ads_options Pipeline#google_ads_options}
	GoogleAdsOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptions `field:"optional" json:"googleAdsOptions" yaml:"googleAdsOptions"`
	// jira_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#jira_options Pipeline#jira_options}
	JiraOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsJiraOptions `field:"optional" json:"jiraOptions" yaml:"jiraOptions"`
	// kafka_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#kafka_options Pipeline#kafka_options}
	KafkaOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptions `field:"optional" json:"kafkaOptions" yaml:"kafkaOptions"`
	// meta_ads_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#meta_ads_options Pipeline#meta_ads_options}
	MetaAdsOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsMetaAdsOptions `field:"optional" json:"metaAdsOptions" yaml:"metaAdsOptions"`
	// outlook_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#outlook_options Pipeline#outlook_options}
	OutlookOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptions `field:"optional" json:"outlookOptions" yaml:"outlookOptions"`
	// reddit_ads_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#reddit_ads_options Pipeline#reddit_ads_options}
	RedditAdsOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsRedditAdsOptions `field:"optional" json:"redditAdsOptions" yaml:"redditAdsOptions"`
	// sharepoint_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#sharepoint_options Pipeline#sharepoint_options}
	SharepointOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSharepointOptions `field:"optional" json:"sharepointOptions" yaml:"sharepointOptions"`
	// smartsheet_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#smartsheet_options Pipeline#smartsheet_options}
	SmartsheetOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSmartsheetOptions `field:"optional" json:"smartsheetOptions" yaml:"smartsheetOptions"`
	// tiktok_ads_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#tiktok_ads_options Pipeline#tiktok_ads_options}
	TiktokAdsOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsTiktokAdsOptions `field:"optional" json:"tiktokAdsOptions" yaml:"tiktokAdsOptions"`
	// zendesk_support_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#zendesk_support_options Pipeline#zendesk_support_options}
	ZendeskSupportOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsZendeskSupportOptions `field:"optional" json:"zendeskSupportOptions" yaml:"zendeskSupportOptions"`
}

