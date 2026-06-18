// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appssettingscustomtemplate


type AppsSettingsCustomTemplateManifest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/apps_settings_custom_template#name AppsSettingsCustomTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/apps_settings_custom_template#version AppsSettingsCustomTemplate#version}.
	Version *float64 `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/apps_settings_custom_template#description AppsSettingsCustomTemplate#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/apps_settings_custom_template#resource_specs AppsSettingsCustomTemplate#resource_specs}.
	ResourceSpecs interface{} `field:"optional" json:"resourceSpecs" yaml:"resourceSpecs"`
}

