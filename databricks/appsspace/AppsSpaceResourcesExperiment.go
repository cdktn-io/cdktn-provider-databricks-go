// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsspace


type AppsSpaceResourcesExperiment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/resources/apps_space#experiment_id AppsSpace#experiment_id}.
	ExperimentId *string `field:"required" json:"experimentId" yaml:"experimentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/resources/apps_space#permission AppsSpace#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

