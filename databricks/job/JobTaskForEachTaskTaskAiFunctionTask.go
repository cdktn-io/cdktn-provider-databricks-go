// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskAiFunctionTask struct {
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/job#output Job#output}
	Output *JobTaskForEachTaskTaskAiFunctionTaskOutput `field:"required" json:"output" yaml:"output"`
	// ai_classify block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/job#ai_classify Job#ai_classify}
	AiClassify *JobTaskForEachTaskTaskAiFunctionTaskAiClassify `field:"optional" json:"aiClassify" yaml:"aiClassify"`
	// ai_extract block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/job#ai_extract Job#ai_extract}
	AiExtract *JobTaskForEachTaskTaskAiFunctionTaskAiExtract `field:"optional" json:"aiExtract" yaml:"aiExtract"`
	// ai_parse_document block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/job#ai_parse_document Job#ai_parse_document}
	AiParseDocument *JobTaskForEachTaskTaskAiFunctionTaskAiParseDocument `field:"optional" json:"aiParseDocument" yaml:"aiParseDocument"`
	// delta_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/job#delta_table Job#delta_table}
	DeltaTable *JobTaskForEachTaskTaskAiFunctionTaskDeltaTable `field:"optional" json:"deltaTable" yaml:"deltaTable"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.144.0/docs/resources/job#volume Job#volume}
	Volume *JobTaskForEachTaskTaskAiFunctionTaskVolume `field:"optional" json:"volume" yaml:"volume"`
}

