// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsTableConnectorOptionsJiraOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/pipeline#include_jira_spaces Pipeline#include_jira_spaces}.
	IncludeJiraSpaces *[]*string `field:"optional" json:"includeJiraSpaces" yaml:"includeJiraSpaces"`
}

