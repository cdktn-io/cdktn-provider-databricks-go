// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringkafkaconfigs


type DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsKeySchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/feature_engineering_kafka_configs#avro_schema DataDatabricksFeatureEngineeringKafkaConfigs#avro_schema}.
	AvroSchema *string `field:"optional" json:"avroSchema" yaml:"avroSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/feature_engineering_kafka_configs#json_schema DataDatabricksFeatureEngineeringKafkaConfigs#json_schema}.
	JsonSchema *string `field:"optional" json:"jsonSchema" yaml:"jsonSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/data-sources/feature_engineering_kafka_configs#proto_schema DataDatabricksFeatureEngineeringKafkaConfigs#proto_schema}.
	ProtoSchema *DataDatabricksFeatureEngineeringKafkaConfigsKafkaConfigsKeySchemaProtoSchema `field:"optional" json:"protoSchema" yaml:"protoSchema"`
}

