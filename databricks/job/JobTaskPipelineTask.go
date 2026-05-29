// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskPipelineTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/job#pipeline_id Job#pipeline_id}.
	PipelineId *string `field:"required" json:"pipelineId" yaml:"pipelineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/job#full_refresh Job#full_refresh}.
	FullRefresh interface{} `field:"optional" json:"fullRefresh" yaml:"fullRefresh"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/job#full_refresh_selection Job#full_refresh_selection}.
	FullRefreshSelection *[]*string `field:"optional" json:"fullRefreshSelection" yaml:"fullRefreshSelection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/job#parameters Job#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/job#refresh_flow_selection Job#refresh_flow_selection}.
	RefreshFlowSelection *[]*string `field:"optional" json:"refreshFlowSelection" yaml:"refreshFlowSelection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/job#refresh_selection Job#refresh_selection}.
	RefreshSelection *[]*string `field:"optional" json:"refreshSelection" yaml:"refreshSelection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/job#reset_checkpoint_selection Job#reset_checkpoint_selection}.
	ResetCheckpointSelection *[]*string `field:"optional" json:"resetCheckpointSelection" yaml:"resetCheckpointSelection"`
}

