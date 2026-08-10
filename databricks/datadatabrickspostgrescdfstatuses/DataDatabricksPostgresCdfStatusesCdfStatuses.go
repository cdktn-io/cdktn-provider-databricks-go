// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgrescdfstatuses


type DataDatabricksPostgresCdfStatusesCdfStatuses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/postgres_cdf_statuses#name DataDatabricksPostgresCdfStatuses#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.125.0/docs/data-sources/postgres_cdf_statuses#provider_config DataDatabricksPostgresCdfStatuses#provider_config}.
	ProviderConfig *DataDatabricksPostgresCdfStatusesCdfStatusesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

