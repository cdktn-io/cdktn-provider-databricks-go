// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsTableConnectorOptionsLinkedinAdsOptionsCustomReportOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#finder Pipeline#finder}.
	Finder *string `field:"required" json:"finder" yaml:"finder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#entity_granularity Pipeline#entity_granularity}.
	EntityGranularity *[]*string `field:"optional" json:"entityGranularity" yaml:"entityGranularity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#metrics Pipeline#metrics}.
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/pipeline#time_granularity Pipeline#time_granularity}.
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

