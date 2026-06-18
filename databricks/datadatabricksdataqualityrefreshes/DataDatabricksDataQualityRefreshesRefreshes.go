// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdataqualityrefreshes


type DataDatabricksDataQualityRefreshesRefreshes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/data_quality_refreshes#object_id DataDatabricksDataQualityRefreshes#object_id}.
	ObjectId *string `field:"required" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/data_quality_refreshes#object_type DataDatabricksDataQualityRefreshes#object_type}.
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/data_quality_refreshes#refresh_id DataDatabricksDataQualityRefreshes#refresh_id}.
	RefreshId *float64 `field:"required" json:"refreshId" yaml:"refreshId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/data-sources/data_quality_refreshes#provider_config DataDatabricksDataQualityRefreshes#provider_config}.
	ProviderConfig *DataDatabricksDataQualityRefreshesRefreshesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

