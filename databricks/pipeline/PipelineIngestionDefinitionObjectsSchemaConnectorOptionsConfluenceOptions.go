// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsSchemaConnectorOptionsConfluenceOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/resources/pipeline#include_confluence_spaces Pipeline#include_confluence_spaces}.
	IncludeConfluenceSpaces *[]*string `field:"optional" json:"includeConfluenceSpaces" yaml:"includeConfluenceSpaces"`
}

