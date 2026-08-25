// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#client_config Pipeline#client_config}.
	ClientConfig *map[string]*string `field:"optional" json:"clientConfig" yaml:"clientConfig"`
	// key_transformer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#key_transformer Pipeline#key_transformer}
	KeyTransformer *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformer `field:"optional" json:"keyTransformer" yaml:"keyTransformer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#max_offsets_per_trigger Pipeline#max_offsets_per_trigger}.
	MaxOffsetsPerTrigger *float64 `field:"optional" json:"maxOffsetsPerTrigger" yaml:"maxOffsetsPerTrigger"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#starting_offset Pipeline#starting_offset}.
	StartingOffset *string `field:"optional" json:"startingOffset" yaml:"startingOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#topic_pattern Pipeline#topic_pattern}.
	TopicPattern *string `field:"optional" json:"topicPattern" yaml:"topicPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#topics Pipeline#topics}.
	Topics *[]*string `field:"optional" json:"topics" yaml:"topics"`
	// value_transformer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#value_transformer Pipeline#value_transformer}
	ValueTransformer *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsValueTransformer `field:"optional" json:"valueTransformer" yaml:"valueTransformer"`
}

