// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy


type DataDatabricksAccountNetworkPolicyIngressDryRunPrivateAccessAllowRulesOrigin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/account_network_policy#all_private_access DataDatabricksAccountNetworkPolicy#all_private_access}.
	AllPrivateAccess interface{} `field:"optional" json:"allPrivateAccess" yaml:"allPrivateAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/account_network_policy#all_registered_endpoints DataDatabricksAccountNetworkPolicy#all_registered_endpoints}.
	AllRegisteredEndpoints interface{} `field:"optional" json:"allRegisteredEndpoints" yaml:"allRegisteredEndpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/account_network_policy#azure_workspace_private_link DataDatabricksAccountNetworkPolicy#azure_workspace_private_link}.
	AzureWorkspacePrivateLink interface{} `field:"optional" json:"azureWorkspacePrivateLink" yaml:"azureWorkspacePrivateLink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/data-sources/account_network_policy#endpoints DataDatabricksAccountNetworkPolicy#endpoints}.
	Endpoints *DataDatabricksAccountNetworkPolicyIngressDryRunPrivateAccessAllowRulesOriginEndpoints `field:"optional" json:"endpoints" yaml:"endpoints"`
}

