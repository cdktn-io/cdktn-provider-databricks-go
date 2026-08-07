// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineTrigger struct {
	// cron block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/pipeline#cron Pipeline#cron}
	Cron *PipelineTriggerCron `field:"optional" json:"cron" yaml:"cron"`
	// manual block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/pipeline#manual Pipeline#manual}
	Manual *PipelineTriggerManual `field:"optional" json:"manual" yaml:"manual"`
}

