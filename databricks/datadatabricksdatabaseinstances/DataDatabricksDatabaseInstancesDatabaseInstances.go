// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabaseinstances


type DataDatabricksDatabaseInstancesDatabaseInstances struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/database_instances#name DataDatabricksDatabaseInstances#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/database_instances#provider_config DataDatabricksDatabaseInstances#provider_config}.
	ProviderConfig *DataDatabricksDatabaseInstancesDatabaseInstancesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

