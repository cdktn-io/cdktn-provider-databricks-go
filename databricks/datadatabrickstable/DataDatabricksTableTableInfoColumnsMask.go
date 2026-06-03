// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickstable


type DataDatabricksTableTableInfoColumnsMask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/table#function_name DataDatabricksTable#function_name}.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// using_arguments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/table#using_arguments DataDatabricksTable#using_arguments}
	UsingArguments interface{} `field:"optional" json:"usingArguments" yaml:"usingArguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/table#using_column_names DataDatabricksTable#using_column_names}.
	UsingColumnNames *[]*string `field:"optional" json:"usingColumnNames" yaml:"usingColumnNames"`
}

