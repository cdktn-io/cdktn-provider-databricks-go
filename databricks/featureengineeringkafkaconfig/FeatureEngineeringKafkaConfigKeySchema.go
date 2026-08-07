// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringkafkaconfig


type FeatureEngineeringKafkaConfigKeySchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/feature_engineering_kafka_config#avro_schema FeatureEngineeringKafkaConfig#avro_schema}.
	AvroSchema *string `field:"optional" json:"avroSchema" yaml:"avroSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/feature_engineering_kafka_config#json_schema FeatureEngineeringKafkaConfig#json_schema}.
	JsonSchema *string `field:"optional" json:"jsonSchema" yaml:"jsonSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/feature_engineering_kafka_config#proto_schema FeatureEngineeringKafkaConfig#proto_schema}.
	ProtoSchema *FeatureEngineeringKafkaConfigKeySchemaProtoSchema `field:"optional" json:"protoSchema" yaml:"protoSchema"`
}

