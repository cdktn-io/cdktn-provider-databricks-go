// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsTableConnectorOptionsGoogleAdsOptionsCustomReportOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#resource Pipeline#resource}.
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#metrics Pipeline#metrics}.
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#resource_fields Pipeline#resource_fields}.
	ResourceFields *[]*string `field:"optional" json:"resourceFields" yaml:"resourceFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/resources/pipeline#segments Pipeline#segments}.
	Segments *[]*string `field:"optional" json:"segments" yaml:"segments"`
}

