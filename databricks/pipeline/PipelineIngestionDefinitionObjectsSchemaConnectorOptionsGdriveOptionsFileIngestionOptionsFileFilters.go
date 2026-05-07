// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsFileIngestionOptionsFileFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/pipeline#modified_after Pipeline#modified_after}.
	ModifiedAfter *string `field:"optional" json:"modifiedAfter" yaml:"modifiedAfter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/pipeline#modified_before Pipeline#modified_before}.
	ModifiedBefore *string `field:"optional" json:"modifiedBefore" yaml:"modifiedBefore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/pipeline#path_filter Pipeline#path_filter}.
	PathFilter *string `field:"optional" json:"pathFilter" yaml:"pathFilter"`
}

