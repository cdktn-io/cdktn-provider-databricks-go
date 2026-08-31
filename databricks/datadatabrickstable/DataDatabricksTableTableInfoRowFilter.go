// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickstable


type DataDatabricksTableTableInfoRowFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/table#function_name DataDatabricksTable#function_name}.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/table#input_column_names DataDatabricksTable#input_column_names}.
	InputColumnNames *[]*string `field:"required" json:"inputColumnNames" yaml:"inputColumnNames"`
	// input_arguments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/table#input_arguments DataDatabricksTable#input_arguments}
	InputArguments interface{} `field:"optional" json:"inputArguments" yaml:"inputArguments"`
}

