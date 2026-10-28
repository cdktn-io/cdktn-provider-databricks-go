// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskDashboardTaskSubscriptionAttachmentPageAttachment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/job#page_ids Job#page_ids}.
	PageIds *[]*string `field:"optional" json:"pageIds" yaml:"pageIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/job#page_selection Job#page_selection}.
	PageSelection *string `field:"optional" json:"pageSelection" yaml:"pageSelection"`
}

