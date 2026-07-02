// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/pipeline#connection_name Pipeline#connection_name}.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/pipeline#connector_type Pipeline#connector_type}.
	ConnectorType *string `field:"optional" json:"connectorType" yaml:"connectorType"`
	// data_staging_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/pipeline#data_staging_options Pipeline#data_staging_options}
	DataStagingOptions *PipelineIngestionDefinitionDataStagingOptions `field:"optional" json:"dataStagingOptions" yaml:"dataStagingOptions"`
	// full_refresh_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/pipeline#full_refresh_window Pipeline#full_refresh_window}
	FullRefreshWindow *PipelineIngestionDefinitionFullRefreshWindow `field:"optional" json:"fullRefreshWindow" yaml:"fullRefreshWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/pipeline#ingest_from_uc_foreign_catalog Pipeline#ingest_from_uc_foreign_catalog}.
	IngestFromUcForeignCatalog interface{} `field:"optional" json:"ingestFromUcForeignCatalog" yaml:"ingestFromUcForeignCatalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/pipeline#ingestion_gateway_id Pipeline#ingestion_gateway_id}.
	IngestionGatewayId *string `field:"optional" json:"ingestionGatewayId" yaml:"ingestionGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/pipeline#netsuite_jar_path Pipeline#netsuite_jar_path}.
	NetsuiteJarPath *string `field:"optional" json:"netsuiteJarPath" yaml:"netsuiteJarPath"`
	// objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/pipeline#objects Pipeline#objects}
	Objects interface{} `field:"optional" json:"objects" yaml:"objects"`
	// source_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/pipeline#source_configurations Pipeline#source_configurations}
	SourceConfigurations interface{} `field:"optional" json:"sourceConfigurations" yaml:"sourceConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/pipeline#source_type Pipeline#source_type}.
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// table_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/pipeline#table_configuration Pipeline#table_configuration}
	TableConfiguration *PipelineIngestionDefinitionTableConfiguration `field:"optional" json:"tableConfiguration" yaml:"tableConfiguration"`
}

