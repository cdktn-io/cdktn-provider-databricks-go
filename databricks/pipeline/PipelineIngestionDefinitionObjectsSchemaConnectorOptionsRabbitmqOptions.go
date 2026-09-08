// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsSchemaConnectorOptionsRabbitmqOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.131.0/docs/resources/pipeline#queue Pipeline#queue}.
	Queue *string `field:"required" json:"queue" yaml:"queue"`
}

