// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsTableConnectorOptionsRedditAdsOptionsCustomReportOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#breakdowns Pipeline#breakdowns}.
	Breakdowns *[]*string `field:"optional" json:"breakdowns" yaml:"breakdowns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#fields Pipeline#fields}.
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
}

