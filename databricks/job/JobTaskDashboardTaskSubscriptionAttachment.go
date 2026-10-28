// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskDashboardTaskSubscriptionAttachment struct {
	// data_attachment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/job#data_attachment Job#data_attachment}
	DataAttachment *JobTaskDashboardTaskSubscriptionAttachmentDataAttachment `field:"optional" json:"dataAttachment" yaml:"dataAttachment"`
	// page_attachment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/job#page_attachment Job#page_attachment}
	PageAttachment *JobTaskDashboardTaskSubscriptionAttachmentPageAttachment `field:"optional" json:"pageAttachment" yaml:"pageAttachment"`
}

