// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#format Pipeline#format}.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// json_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#json_options Pipeline#json_options}
	JsonOptions *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsValueTransformerJsonOptions `field:"optional" json:"jsonOptions" yaml:"jsonOptions"`
}

