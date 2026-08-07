// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualitymonitors


type DataDatabricksDataQualityMonitorsMonitors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/data_quality_monitors#object_id DataDatabricksDataQualityMonitors#object_id}.
	ObjectId *string `field:"required" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/data_quality_monitors#object_type DataDatabricksDataQualityMonitors#object_type}.
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/data_quality_monitors#provider_config DataDatabricksDataQualityMonitors#provider_config}.
	ProviderConfig *DataDatabricksDataQualityMonitorsMonitorsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

