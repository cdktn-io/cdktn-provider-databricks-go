// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/pipeline#format Pipeline#format}.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/pipeline#input_column Pipeline#input_column}.
	InputColumn *string `field:"optional" json:"inputColumn" yaml:"inputColumn"`
	// json_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/pipeline#json_options Pipeline#json_options}
	JsonOptions *PipelineIngestionDefinitionObjectsTableConnectorOptionsKafkaOptionsKeyTransformerJsonOptions `field:"optional" json:"jsonOptions" yaml:"jsonOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/pipeline#output_column Pipeline#output_column}.
	OutputColumn *string `field:"optional" json:"outputColumn" yaml:"outputColumn"`
}

