// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringkafkaconfig


type FeatureEngineeringKafkaConfigValueSchemaProtoSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/feature_engineering_kafka_config#message_name FeatureEngineeringKafkaConfig#message_name}.
	MessageName *string `field:"required" json:"messageName" yaml:"messageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/feature_engineering_kafka_config#schema_text FeatureEngineeringKafkaConfig#schema_text}.
	SchemaText *string `field:"required" json:"schemaText" yaml:"schemaText"`
}

