// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appssettingscustomtemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppsSettingsCustomTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/apps_settings_custom_template#git_provider AppsSettingsCustomTemplate#git_provider}.
	GitProvider *string `field:"required" json:"gitProvider" yaml:"gitProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/apps_settings_custom_template#git_repo AppsSettingsCustomTemplate#git_repo}.
	GitRepo *string `field:"required" json:"gitRepo" yaml:"gitRepo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/apps_settings_custom_template#manifest AppsSettingsCustomTemplate#manifest}.
	Manifest *AppsSettingsCustomTemplateManifest `field:"required" json:"manifest" yaml:"manifest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/apps_settings_custom_template#name AppsSettingsCustomTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/apps_settings_custom_template#path AppsSettingsCustomTemplate#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/apps_settings_custom_template#description AppsSettingsCustomTemplate#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/apps_settings_custom_template#provider_config AppsSettingsCustomTemplate#provider_config}.
	ProviderConfig *AppsSettingsCustomTemplateProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

