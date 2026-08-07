// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspaceSettingV2Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#aibi_dashboard_embedding_access_policy WorkspaceSettingV2#aibi_dashboard_embedding_access_policy}.
	AibiDashboardEmbeddingAccessPolicy *WorkspaceSettingV2AibiDashboardEmbeddingAccessPolicy `field:"optional" json:"aibiDashboardEmbeddingAccessPolicy" yaml:"aibiDashboardEmbeddingAccessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#aibi_dashboard_embedding_approved_domains WorkspaceSettingV2#aibi_dashboard_embedding_approved_domains}.
	AibiDashboardEmbeddingApprovedDomains *WorkspaceSettingV2AibiDashboardEmbeddingApprovedDomains `field:"optional" json:"aibiDashboardEmbeddingApprovedDomains" yaml:"aibiDashboardEmbeddingApprovedDomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#allowed_apps_user_api_scopes WorkspaceSettingV2#allowed_apps_user_api_scopes}.
	AllowedAppsUserApiScopes *WorkspaceSettingV2AllowedAppsUserApiScopes `field:"optional" json:"allowedAppsUserApiScopes" yaml:"allowedAppsUserApiScopes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#automatic_cluster_update_workspace WorkspaceSettingV2#automatic_cluster_update_workspace}.
	AutomaticClusterUpdateWorkspace *WorkspaceSettingV2AutomaticClusterUpdateWorkspace `field:"optional" json:"automaticClusterUpdateWorkspace" yaml:"automaticClusterUpdateWorkspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#boolean_val WorkspaceSettingV2#boolean_val}.
	BooleanVal *WorkspaceSettingV2BooleanVal `field:"optional" json:"booleanVal" yaml:"booleanVal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#collaboration_platform_connectivity WorkspaceSettingV2#collaboration_platform_connectivity}.
	CollaborationPlatformConnectivity *WorkspaceSettingV2CollaborationPlatformConnectivity `field:"optional" json:"collaborationPlatformConnectivity" yaml:"collaborationPlatformConnectivity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#effective_aibi_dashboard_embedding_access_policy WorkspaceSettingV2#effective_aibi_dashboard_embedding_access_policy}.
	EffectiveAibiDashboardEmbeddingAccessPolicy *WorkspaceSettingV2EffectiveAibiDashboardEmbeddingAccessPolicy `field:"optional" json:"effectiveAibiDashboardEmbeddingAccessPolicy" yaml:"effectiveAibiDashboardEmbeddingAccessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#effective_aibi_dashboard_embedding_approved_domains WorkspaceSettingV2#effective_aibi_dashboard_embedding_approved_domains}.
	EffectiveAibiDashboardEmbeddingApprovedDomains *WorkspaceSettingV2EffectiveAibiDashboardEmbeddingApprovedDomains `field:"optional" json:"effectiveAibiDashboardEmbeddingApprovedDomains" yaml:"effectiveAibiDashboardEmbeddingApprovedDomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#effective_automatic_cluster_update_workspace WorkspaceSettingV2#effective_automatic_cluster_update_workspace}.
	EffectiveAutomaticClusterUpdateWorkspace *WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspace `field:"optional" json:"effectiveAutomaticClusterUpdateWorkspace" yaml:"effectiveAutomaticClusterUpdateWorkspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#effective_personal_compute WorkspaceSettingV2#effective_personal_compute}.
	EffectivePersonalCompute *WorkspaceSettingV2EffectivePersonalCompute `field:"optional" json:"effectivePersonalCompute" yaml:"effectivePersonalCompute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#effective_restrict_workspace_admins WorkspaceSettingV2#effective_restrict_workspace_admins}.
	EffectiveRestrictWorkspaceAdmins *WorkspaceSettingV2EffectiveRestrictWorkspaceAdmins `field:"optional" json:"effectiveRestrictWorkspaceAdmins" yaml:"effectiveRestrictWorkspaceAdmins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#integer_val WorkspaceSettingV2#integer_val}.
	IntegerVal *WorkspaceSettingV2IntegerVal `field:"optional" json:"integerVal" yaml:"integerVal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#name WorkspaceSettingV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#operational_email_custom_recipient WorkspaceSettingV2#operational_email_custom_recipient}.
	OperationalEmailCustomRecipient *WorkspaceSettingV2OperationalEmailCustomRecipient `field:"optional" json:"operationalEmailCustomRecipient" yaml:"operationalEmailCustomRecipient"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#personal_compute WorkspaceSettingV2#personal_compute}.
	PersonalCompute *WorkspaceSettingV2PersonalCompute `field:"optional" json:"personalCompute" yaml:"personalCompute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#provider_config WorkspaceSettingV2#provider_config}.
	ProviderConfig *WorkspaceSettingV2ProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#restrict_workspace_admins WorkspaceSettingV2#restrict_workspace_admins}.
	RestrictWorkspaceAdmins *WorkspaceSettingV2RestrictWorkspaceAdmins `field:"optional" json:"restrictWorkspaceAdmins" yaml:"restrictWorkspaceAdmins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.124.0/docs/resources/workspace_setting_v2#string_val WorkspaceSettingV2#string_val}.
	StringVal *WorkspaceSettingV2StringVal `field:"optional" json:"stringVal" yaml:"stringVal"`
}

