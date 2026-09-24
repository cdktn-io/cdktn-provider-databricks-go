// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssandboxes


type DataDatabricksSandboxesSandboxes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/data-sources/sandboxes#name DataDatabricksSandboxes#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.134.0/docs/data-sources/sandboxes#provider_config DataDatabricksSandboxes#provider_config}.
	ProviderConfig *DataDatabricksSandboxesSandboxesProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

