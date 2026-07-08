// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksqualitymonitorsv2


type DataDatabricksQualityMonitorsV2QualityMonitors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/quality_monitors_v2#object_id DataDatabricksQualityMonitorsV2#object_id}.
	ObjectId *string `field:"required" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/quality_monitors_v2#object_type DataDatabricksQualityMonitorsV2#object_type}.
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/quality_monitors_v2#provider_config DataDatabricksQualityMonitorsV2#provider_config}.
	ProviderConfig *DataDatabricksQualityMonitorsV2QualityMonitorsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

