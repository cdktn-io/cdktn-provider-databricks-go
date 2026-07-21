// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appspace


type AppSpaceResourcesExperiment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/app_space#experiment_id AppSpace#experiment_id}.
	ExperimentId *string `field:"required" json:"experimentId" yaml:"experimentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/app_space#permission AppSpace#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

