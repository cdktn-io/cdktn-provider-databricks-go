// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/account_network_policies#account_api DataDatabricksAccountNetworkPolicies#account_api}.
	AccountApi *DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountApi `field:"optional" json:"accountApi" yaml:"accountApi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/account_network_policies#account_databricks_one DataDatabricksAccountNetworkPolicies#account_databricks_one}.
	AccountDatabricksOne *DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountDatabricksOne `field:"optional" json:"accountDatabricksOne" yaml:"accountDatabricksOne"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/account_network_policies#account_ui DataDatabricksAccountNetworkPolicies#account_ui}.
	AccountUi *DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountUi `field:"optional" json:"accountUi" yaml:"accountUi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/account_network_policies#all_destinations DataDatabricksAccountNetworkPolicies#all_destinations}.
	AllDestinations interface{} `field:"optional" json:"allDestinations" yaml:"allDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/account_network_policies#apps_runtime DataDatabricksAccountNetworkPolicies#apps_runtime}.
	AppsRuntime *DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAppsRuntime `field:"optional" json:"appsRuntime" yaml:"appsRuntime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/account_network_policies#lakebase_runtime DataDatabricksAccountNetworkPolicies#lakebase_runtime}.
	LakebaseRuntime *DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationLakebaseRuntime `field:"optional" json:"lakebaseRuntime" yaml:"lakebaseRuntime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/account_network_policies#workspace_api DataDatabricksAccountNetworkPolicies#workspace_api}.
	WorkspaceApi *DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationWorkspaceApi `field:"optional" json:"workspaceApi" yaml:"workspaceApi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.116.0/docs/data-sources/account_network_policies#workspace_ui DataDatabricksAccountNetworkPolicies#workspace_ui}.
	WorkspaceUi *DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationWorkspaceUi `field:"optional" json:"workspaceUi" yaml:"workspaceUi"`
}

