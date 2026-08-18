// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsTableConnectorOptionsSharepointOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/pipeline#entity_type Pipeline#entity_type}.
	EntityType *string `field:"optional" json:"entityType" yaml:"entityType"`
	// file_ingestion_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/pipeline#file_ingestion_options Pipeline#file_ingestion_options}
	FileIngestionOptions *PipelineIngestionDefinitionObjectsTableConnectorOptionsSharepointOptionsFileIngestionOptions `field:"optional" json:"fileIngestionOptions" yaml:"fileIngestionOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/pipeline#url Pipeline#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

