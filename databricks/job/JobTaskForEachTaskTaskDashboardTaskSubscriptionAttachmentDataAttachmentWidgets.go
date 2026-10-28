// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachmentDataAttachmentWidgets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/job#page_id Job#page_id}.
	PageId *string `field:"required" json:"pageId" yaml:"pageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/job#widget_id Job#widget_id}.
	WidgetId *string `field:"required" json:"widgetId" yaml:"widgetId"`
}

