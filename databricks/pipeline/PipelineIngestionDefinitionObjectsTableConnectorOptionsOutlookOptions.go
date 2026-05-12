// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsTableConnectorOptionsOutlookOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/pipeline#attachment_mode Pipeline#attachment_mode}.
	AttachmentMode *string `field:"optional" json:"attachmentMode" yaml:"attachmentMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/pipeline#body_format Pipeline#body_format}.
	BodyFormat *string `field:"optional" json:"bodyFormat" yaml:"bodyFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/pipeline#folder_filter Pipeline#folder_filter}.
	FolderFilter *[]*string `field:"optional" json:"folderFilter" yaml:"folderFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/pipeline#include_folders Pipeline#include_folders}.
	IncludeFolders *[]*string `field:"optional" json:"includeFolders" yaml:"includeFolders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/pipeline#include_mailboxes Pipeline#include_mailboxes}.
	IncludeMailboxes *[]*string `field:"optional" json:"includeMailboxes" yaml:"includeMailboxes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/pipeline#include_senders Pipeline#include_senders}.
	IncludeSenders *[]*string `field:"optional" json:"includeSenders" yaml:"includeSenders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/pipeline#include_subjects Pipeline#include_subjects}.
	IncludeSubjects *[]*string `field:"optional" json:"includeSubjects" yaml:"includeSubjects"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/pipeline#sender_filter Pipeline#sender_filter}.
	SenderFilter *[]*string `field:"optional" json:"senderFilter" yaml:"senderFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/pipeline#start_date Pipeline#start_date}.
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.115.0/docs/resources/pipeline#subject_filter Pipeline#subject_filter}.
	SubjectFilter *[]*string `field:"optional" json:"subjectFilter" yaml:"subjectFilter"`
}

