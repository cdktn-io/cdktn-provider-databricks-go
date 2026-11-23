// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchalertconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LakewatchAlertConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_alert_config#parent LakewatchAlertConfig#parent}.
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_alert_config#default_security_alert_collation_window LakewatchAlertConfig#default_security_alert_collation_window}.
	DefaultSecurityAlertCollationWindow *string `field:"optional" json:"defaultSecurityAlertCollationWindow" yaml:"defaultSecurityAlertCollationWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_alert_config#default_system_alert_collation_window LakewatchAlertConfig#default_system_alert_collation_window}.
	DefaultSystemAlertCollationWindow *string `field:"optional" json:"defaultSystemAlertCollationWindow" yaml:"defaultSystemAlertCollationWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_alert_config#provider_config LakewatchAlertConfig#provider_config}.
	ProviderConfig *LakewatchAlertConfigProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_alert_config#security_alert_export_config LakewatchAlertConfig#security_alert_export_config}.
	SecurityAlertExportConfig *LakewatchAlertConfigSecurityAlertExportConfig `field:"optional" json:"securityAlertExportConfig" yaml:"securityAlertExportConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.145.0/docs/resources/lakewatch_alert_config#system_alert_export_config LakewatchAlertConfig#system_alert_export_config}.
	SystemAlertExportConfig *LakewatchAlertConfigSystemAlertExportConfig `field:"optional" json:"systemAlertExportConfig" yaml:"systemAlertExportConfig"`
}

