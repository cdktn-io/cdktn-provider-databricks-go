// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerProtobufOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#desc_file_path Pipeline#desc_file_path}.
	DescFilePath *string `field:"optional" json:"descFilePath" yaml:"descFilePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#message_name Pipeline#message_name}.
	MessageName *string `field:"optional" json:"messageName" yaml:"messageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#parse_mode Pipeline#parse_mode}.
	ParseMode *string `field:"optional" json:"parseMode" yaml:"parseMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#recursive_fields_max_depth Pipeline#recursive_fields_max_depth}.
	RecursiveFieldsMaxDepth *float64 `field:"optional" json:"recursiveFieldsMaxDepth" yaml:"recursiveFieldsMaxDepth"`
	// schema_registry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#schema_registry Pipeline#schema_registry}
	SchemaRegistry *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerProtobufOptionsSchemaRegistry `field:"optional" json:"schemaRegistry" yaml:"schemaRegistry"`
}

