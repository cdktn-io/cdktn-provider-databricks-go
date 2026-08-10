// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfig


type DataDatabricksFeatureEngineeringKafkaConfigValueSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/feature_engineering_kafka_config#avro_schema DataDatabricksFeatureEngineeringKafkaConfig#avro_schema}.
	AvroSchema *string `field:"optional" json:"avroSchema" yaml:"avroSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/feature_engineering_kafka_config#json_schema DataDatabricksFeatureEngineeringKafkaConfig#json_schema}.
	JsonSchema *string `field:"optional" json:"jsonSchema" yaml:"jsonSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/feature_engineering_kafka_config#proto_schema DataDatabricksFeatureEngineeringKafkaConfig#proto_schema}.
	ProtoSchema *DataDatabricksFeatureEngineeringKafkaConfigValueSchemaProtoSchema `field:"optional" json:"protoSchema" yaml:"protoSchema"`
}

