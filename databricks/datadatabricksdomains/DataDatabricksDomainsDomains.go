// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdomains


type DataDatabricksDomainsDomains struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/data-sources/domains#name DataDatabricksDomains#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/data-sources/domains#provider_config DataDatabricksDomains#provider_config}.
	ProviderConfig *DataDatabricksDomainsDomainsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

