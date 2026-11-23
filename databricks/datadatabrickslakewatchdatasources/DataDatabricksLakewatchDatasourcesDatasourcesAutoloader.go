// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchdatasources


type DataDatabricksLakewatchDatasourcesDatasourcesAutoloader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_datasources#format DataDatabricksLakewatchDatasources#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_datasources#location DataDatabricksLakewatchDatasources#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_datasources#raw_table DataDatabricksLakewatchDatasources#raw_table}.
	RawTable *DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderRawTable `field:"required" json:"rawTable" yaml:"rawTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_datasources#aws_secret_scope DataDatabricksLakewatchDatasources#aws_secret_scope}.
	AwsSecretScope *string `field:"optional" json:"awsSecretScope" yaml:"awsSecretScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_datasources#envelope DataDatabricksLakewatchDatasources#envelope}.
	Envelope *DataDatabricksLakewatchDatasourcesDatasourcesAutoloaderEnvelope `field:"optional" json:"envelope" yaml:"envelope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_datasources#format_options DataDatabricksLakewatchDatasources#format_options}.
	FormatOptions *map[string]*string `field:"optional" json:"formatOptions" yaml:"formatOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_datasources#include_existing_files DataDatabricksLakewatchDatasources#include_existing_files}.
	IncludeExistingFiles interface{} `field:"optional" json:"includeExistingFiles" yaml:"includeExistingFiles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_datasources#override_partition_columns DataDatabricksLakewatchDatasources#override_partition_columns}.
	OverridePartitionColumns *[]*string `field:"optional" json:"overridePartitionColumns" yaml:"overridePartitionColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_datasources#queue_url DataDatabricksLakewatchDatasources#queue_url}.
	QueueUrl *string `field:"optional" json:"queueUrl" yaml:"queueUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/data-sources/lakewatch_datasources#service_credentials DataDatabricksLakewatchDatasources#service_credentials}.
	ServiceCredentials *string `field:"optional" json:"serviceCredentials" yaml:"serviceCredentials"`
}

