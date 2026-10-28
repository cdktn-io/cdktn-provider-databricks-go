// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDashboardTaskSubscriptionAttachment struct {
	// data_attachment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/job#data_attachment DataDatabricksJob#data_attachment}
	DataAttachment *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentDataAttachment `field:"optional" json:"dataAttachment" yaml:"dataAttachment"`
	// page_attachment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/data-sources/job#page_attachment DataDatabricksJob#page_attachment}
	PageAttachment *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentPageAttachment `field:"optional" json:"pageAttachment" yaml:"pageAttachment"`
}

