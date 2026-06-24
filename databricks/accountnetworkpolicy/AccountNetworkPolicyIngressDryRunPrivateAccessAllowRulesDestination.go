// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyIngressDryRunPrivateAccessAllowRulesDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/account_network_policy#account_api AccountNetworkPolicy#account_api}.
	AccountApi *AccountNetworkPolicyIngressDryRunPrivateAccessAllowRulesDestinationAccountApi `field:"optional" json:"accountApi" yaml:"accountApi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/account_network_policy#account_databricks_one AccountNetworkPolicy#account_databricks_one}.
	AccountDatabricksOne *AccountNetworkPolicyIngressDryRunPrivateAccessAllowRulesDestinationAccountDatabricksOne `field:"optional" json:"accountDatabricksOne" yaml:"accountDatabricksOne"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/account_network_policy#account_ui AccountNetworkPolicy#account_ui}.
	AccountUi *AccountNetworkPolicyIngressDryRunPrivateAccessAllowRulesDestinationAccountUi `field:"optional" json:"accountUi" yaml:"accountUi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/account_network_policy#all_destinations AccountNetworkPolicy#all_destinations}.
	AllDestinations interface{} `field:"optional" json:"allDestinations" yaml:"allDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/account_network_policy#apps_runtime AccountNetworkPolicy#apps_runtime}.
	AppsRuntime *AccountNetworkPolicyIngressDryRunPrivateAccessAllowRulesDestinationAppsRuntime `field:"optional" json:"appsRuntime" yaml:"appsRuntime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/account_network_policy#lakebase_runtime AccountNetworkPolicy#lakebase_runtime}.
	LakebaseRuntime *AccountNetworkPolicyIngressDryRunPrivateAccessAllowRulesDestinationLakebaseRuntime `field:"optional" json:"lakebaseRuntime" yaml:"lakebaseRuntime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/account_network_policy#workspace_api AccountNetworkPolicy#workspace_api}.
	WorkspaceApi *AccountNetworkPolicyIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceApi `field:"optional" json:"workspaceApi" yaml:"workspaceApi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/account_network_policy#workspace_ui AccountNetworkPolicy#workspace_ui}.
	WorkspaceUi *AccountNetworkPolicyIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceUi `field:"optional" json:"workspaceUi" yaml:"workspaceUi"`
}

