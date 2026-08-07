// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicy


type DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessAllowRulesDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/account_network_policy#account_api DataDatabricksAccountNetworkPolicy#account_api}.
	AccountApi *DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessAllowRulesDestinationAccountApi `field:"optional" json:"accountApi" yaml:"accountApi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/account_network_policy#account_databricks_one DataDatabricksAccountNetworkPolicy#account_databricks_one}.
	AccountDatabricksOne *DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessAllowRulesDestinationAccountDatabricksOne `field:"optional" json:"accountDatabricksOne" yaml:"accountDatabricksOne"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/account_network_policy#account_ui DataDatabricksAccountNetworkPolicy#account_ui}.
	AccountUi *DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessAllowRulesDestinationAccountUi `field:"optional" json:"accountUi" yaml:"accountUi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/account_network_policy#all_destinations DataDatabricksAccountNetworkPolicy#all_destinations}.
	AllDestinations interface{} `field:"optional" json:"allDestinations" yaml:"allDestinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/account_network_policy#apps_runtime DataDatabricksAccountNetworkPolicy#apps_runtime}.
	AppsRuntime *DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessAllowRulesDestinationAppsRuntime `field:"optional" json:"appsRuntime" yaml:"appsRuntime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/account_network_policy#lakebase_runtime DataDatabricksAccountNetworkPolicy#lakebase_runtime}.
	LakebaseRuntime *DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessAllowRulesDestinationLakebaseRuntime `field:"optional" json:"lakebaseRuntime" yaml:"lakebaseRuntime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/account_network_policy#workspace_api DataDatabricksAccountNetworkPolicy#workspace_api}.
	WorkspaceApi *DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessAllowRulesDestinationWorkspaceApi `field:"optional" json:"workspaceApi" yaml:"workspaceApi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/data-sources/account_network_policy#workspace_ui DataDatabricksAccountNetworkPolicy#workspace_ui}.
	WorkspaceUi *DataDatabricksAccountNetworkPolicyIngressDryRunPublicAccessAllowRulesDestinationWorkspaceUi `field:"optional" json:"workspaceUi" yaml:"workspaceUi"`
}

