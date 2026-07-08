// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSharepointOptionsFileIngestionOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#corrupt_record_column Pipeline#corrupt_record_column}.
	CorruptRecordColumn *string `field:"optional" json:"corruptRecordColumn" yaml:"corruptRecordColumn"`
	// file_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#file_filters Pipeline#file_filters}
	FileFilters interface{} `field:"optional" json:"fileFilters" yaml:"fileFilters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#format Pipeline#format}.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#format_options Pipeline#format_options}.
	FormatOptions *map[string]*string `field:"optional" json:"formatOptions" yaml:"formatOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#ignore_corrupt_files Pipeline#ignore_corrupt_files}.
	IgnoreCorruptFiles interface{} `field:"optional" json:"ignoreCorruptFiles" yaml:"ignoreCorruptFiles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#infer_column_types Pipeline#infer_column_types}.
	InferColumnTypes interface{} `field:"optional" json:"inferColumnTypes" yaml:"inferColumnTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#reader_case_sensitive Pipeline#reader_case_sensitive}.
	ReaderCaseSensitive interface{} `field:"optional" json:"readerCaseSensitive" yaml:"readerCaseSensitive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#rescued_data_column Pipeline#rescued_data_column}.
	RescuedDataColumn *string `field:"optional" json:"rescuedDataColumn" yaml:"rescuedDataColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#schema_evolution_mode Pipeline#schema_evolution_mode}.
	SchemaEvolutionMode *string `field:"optional" json:"schemaEvolutionMode" yaml:"schemaEvolutionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#schema_hints Pipeline#schema_hints}.
	SchemaHints *string `field:"optional" json:"schemaHints" yaml:"schemaHints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/pipeline#single_variant_column Pipeline#single_variant_column}.
	SingleVariantColumn *string `field:"optional" json:"singleVariantColumn" yaml:"singleVariantColumn"`
}

