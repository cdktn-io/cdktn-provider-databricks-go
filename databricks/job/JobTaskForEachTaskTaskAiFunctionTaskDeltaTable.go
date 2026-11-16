// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskAiFunctionTaskDeltaTable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/job#input_column Job#input_column}.
	InputColumn *string `field:"required" json:"inputColumn" yaml:"inputColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/job#table_name Job#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/resources/job#merge_key_columns Job#merge_key_columns}.
	MergeKeyColumns *[]*string `field:"optional" json:"mergeKeyColumns" yaml:"mergeKeyColumns"`
}

