// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsSchemaFanoutOptionsTransformsJsonOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#as_variant Pipeline#as_variant}.
	AsVariant interface{} `field:"optional" json:"asVariant" yaml:"asVariant"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#schema Pipeline#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#schema_evolution_mode Pipeline#schema_evolution_mode}.
	SchemaEvolutionMode *string `field:"optional" json:"schemaEvolutionMode" yaml:"schemaEvolutionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#schema_file_path Pipeline#schema_file_path}.
	SchemaFilePath *string `field:"optional" json:"schemaFilePath" yaml:"schemaFilePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#schema_hints Pipeline#schema_hints}.
	SchemaHints *string `field:"optional" json:"schemaHints" yaml:"schemaHints"`
}

