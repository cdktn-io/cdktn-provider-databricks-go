// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssecretucs


type DataDatabricksSecretUcsSecrets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/secret_ucs#full_name DataDatabricksSecretUcs#full_name}.
	FullName *string `field:"required" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/secret_ucs#provider_config DataDatabricksSecretUcs#provider_config}.
	ProviderConfig *DataDatabricksSecretUcsSecretsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

