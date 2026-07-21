// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyIngressPrivateAccessAllowRulesOrigin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/account_network_policy#all_private_access AccountNetworkPolicy#all_private_access}.
	AllPrivateAccess interface{} `field:"optional" json:"allPrivateAccess" yaml:"allPrivateAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/account_network_policy#all_registered_endpoints AccountNetworkPolicy#all_registered_endpoints}.
	AllRegisteredEndpoints interface{} `field:"optional" json:"allRegisteredEndpoints" yaml:"allRegisteredEndpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/account_network_policy#azure_workspace_private_link AccountNetworkPolicy#azure_workspace_private_link}.
	AzureWorkspacePrivateLink interface{} `field:"optional" json:"azureWorkspacePrivateLink" yaml:"azureWorkspacePrivateLink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/account_network_policy#endpoints AccountNetworkPolicy#endpoints}.
	Endpoints *AccountNetworkPolicyIngressPrivateAccessAllowRulesOriginEndpoints `field:"optional" json:"endpoints" yaml:"endpoints"`
}

