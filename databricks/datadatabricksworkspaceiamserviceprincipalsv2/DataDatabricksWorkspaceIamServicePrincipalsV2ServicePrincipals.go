// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspaceiamserviceprincipalsv2


type DataDatabricksWorkspaceIamServicePrincipalsV2ServicePrincipals struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/workspace_iam_service_principals_v2#name DataDatabricksWorkspaceIamServicePrincipalsV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.143.0/docs/data-sources/workspace_iam_service_principals_v2#provider_config DataDatabricksWorkspaceIamServicePrincipalsV2#provider_config}.
	ProviderConfig *DataDatabricksWorkspaceIamServicePrincipalsV2ServicePrincipalsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

