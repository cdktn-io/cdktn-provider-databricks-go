// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgrescdfconfigs


type DataDatabricksPostgresCdfConfigsCdfConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/postgres_cdf_configs#name DataDatabricksPostgresCdfConfigs#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/postgres_cdf_configs#provider_config DataDatabricksPostgresCdfConfigs#provider_config}.
	ProviderConfig *DataDatabricksPostgresCdfConfigsCdfConfigsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

