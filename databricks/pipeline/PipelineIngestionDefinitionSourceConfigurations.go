// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionSourceConfigurations struct {
	// catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.113.0/docs/resources/pipeline#catalog Pipeline#catalog}
	Catalog *PipelineIngestionDefinitionSourceConfigurationsCatalog `field:"optional" json:"catalog" yaml:"catalog"`
}

