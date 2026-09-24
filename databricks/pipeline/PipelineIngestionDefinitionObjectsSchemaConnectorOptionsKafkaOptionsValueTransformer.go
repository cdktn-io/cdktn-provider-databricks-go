// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformer struct {
	// avro_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#avro_options Pipeline#avro_options}
	AvroOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerAvroOptions `field:"optional" json:"avroOptions" yaml:"avroOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#format Pipeline#format}.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#input_column Pipeline#input_column}.
	InputColumn *string `field:"optional" json:"inputColumn" yaml:"inputColumn"`
	// json_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#json_options Pipeline#json_options}
	JsonOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerJsonOptions `field:"optional" json:"jsonOptions" yaml:"jsonOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#output_column Pipeline#output_column}.
	OutputColumn *string `field:"optional" json:"outputColumn" yaml:"outputColumn"`
	// protobuf_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/resources/pipeline#protobuf_options Pipeline#protobuf_options}
	ProtobufOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerProtobufOptions `field:"optional" json:"protobufOptions" yaml:"protobufOptions"`
}

