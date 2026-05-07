// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTriggerPeriodic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/job#interval DataDatabricksJob#interval}.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/data-sources/job#unit DataDatabricksJob#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
}

