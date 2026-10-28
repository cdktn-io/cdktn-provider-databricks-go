// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskDashboardTaskSubscription struct {
	// attachment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/job#attachment Job#attachment}
	Attachment *JobTaskForEachTaskTaskDashboardTaskSubscriptionAttachment `field:"optional" json:"attachment" yaml:"attachment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/job#custom_subject Job#custom_subject}.
	CustomSubject *string `field:"optional" json:"customSubject" yaml:"customSubject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/job#paused Job#paused}.
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// subscribers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.140.0/docs/resources/job#subscribers Job#subscribers}
	Subscribers interface{} `field:"optional" json:"subscribers" yaml:"subscribers"`
}

