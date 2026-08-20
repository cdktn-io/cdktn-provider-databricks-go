// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsSchemaFanoutOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/pipeline#fanout_by Pipeline#fanout_by}.
	FanoutBy *string `field:"optional" json:"fanoutBy" yaml:"fanoutBy"`
	// transforms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/pipeline#transforms Pipeline#transforms}
	Transforms interface{} `field:"optional" json:"transforms" yaml:"transforms"`
}

