// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionSourceConfigurationsApiSourceConnectorConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/pipeline#configs Pipeline#configs}.
	Configs *map[string]*string `field:"optional" json:"configs" yaml:"configs"`
}

