// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapps


type DataDatabricksAppsAppResourcesExperiment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/apps#experiment_id DataDatabricksApps#experiment_id}.
	ExperimentId *string `field:"required" json:"experimentId" yaml:"experimentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/data-sources/apps#permission DataDatabricksApps#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

