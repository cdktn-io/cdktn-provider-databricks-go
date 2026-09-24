// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsValueTransformerAvroOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#parse_mode Pipeline#parse_mode}.
	ParseMode *string `field:"optional" json:"parseMode" yaml:"parseMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#schema Pipeline#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#schema_file_path Pipeline#schema_file_path}.
	SchemaFilePath *string `field:"optional" json:"schemaFilePath" yaml:"schemaFilePath"`
	// schema_registry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#schema_registry Pipeline#schema_registry}
	SchemaRegistry *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistry `field:"optional" json:"schemaRegistry" yaml:"schemaRegistry"`
}

