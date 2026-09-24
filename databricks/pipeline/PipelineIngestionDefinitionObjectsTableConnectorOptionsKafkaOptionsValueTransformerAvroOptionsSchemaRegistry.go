// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistry struct {
	// confluent_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#confluent_options Pipeline#confluent_options}
	ConfluentOptions *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsValueTransformerAvroOptionsSchemaRegistryConfluentOptions `field:"optional" json:"confluentOptions" yaml:"confluentOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#connection_name Pipeline#connection_name}.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#protobuf_message_name Pipeline#protobuf_message_name}.
	ProtobufMessageName *string `field:"optional" json:"protobufMessageName" yaml:"protobufMessageName"`
}

