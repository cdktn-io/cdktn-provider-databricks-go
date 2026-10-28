// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskDashboardTaskSubscriptionAttachmentDataAttachment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/job#file_format Job#file_format}.
	FileFormat *string `field:"required" json:"fileFormat" yaml:"fileFormat"`
	// widgets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/job#widgets Job#widgets}
	Widgets interface{} `field:"required" json:"widgets" yaml:"widgets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/job#include_applied_filters Job#include_applied_filters}.
	IncludeAppliedFilters interface{} `field:"optional" json:"includeAppliedFilters" yaml:"includeAppliedFilters"`
}

