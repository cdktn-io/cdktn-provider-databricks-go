// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapp


type DataDatabricksAppAppResourcesExperiment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app#experiment_id DataDatabricksApp#experiment_id}.
	ExperimentId *string `field:"required" json:"experimentId" yaml:"experimentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/data-sources/app#permission DataDatabricksApp#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

