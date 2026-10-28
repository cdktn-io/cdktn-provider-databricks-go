// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskDashboardTaskSubscriptionAttachmentDataAttachment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/job#file_format DataDatabricksJob#file_format}.
	FileFormat *string `field:"required" json:"fileFormat" yaml:"fileFormat"`
	// widgets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/job#widgets DataDatabricksJob#widgets}
	Widgets interface{} `field:"required" json:"widgets" yaml:"widgets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/job#include_applied_filters DataDatabricksJob#include_applied_filters}.
	IncludeAppliedFilters interface{} `field:"optional" json:"includeAppliedFilters" yaml:"includeAppliedFilters"`
}

