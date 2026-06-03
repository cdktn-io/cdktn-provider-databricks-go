// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesOrigin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/account_network_policies#all_private_access DataDatabricksAccountNetworkPolicies#all_private_access}.
	AllPrivateAccess interface{} `field:"optional" json:"allPrivateAccess" yaml:"allPrivateAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/account_network_policies#all_registered_endpoints DataDatabricksAccountNetworkPolicies#all_registered_endpoints}.
	AllRegisteredEndpoints interface{} `field:"optional" json:"allRegisteredEndpoints" yaml:"allRegisteredEndpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/account_network_policies#azure_workspace_private_link DataDatabricksAccountNetworkPolicies#azure_workspace_private_link}.
	AzureWorkspacePrivateLink interface{} `field:"optional" json:"azureWorkspacePrivateLink" yaml:"azureWorkspacePrivateLink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/data-sources/account_network_policies#endpoints DataDatabricksAccountNetworkPolicies#endpoints}.
	Endpoints *DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesOriginEndpoints `field:"optional" json:"endpoints" yaml:"endpoints"`
}

