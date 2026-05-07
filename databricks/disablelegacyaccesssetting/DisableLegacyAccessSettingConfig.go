// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package disablelegacyaccesssetting

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DisableLegacyAccessSettingConfig struct {
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
	// disable_legacy_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/disable_legacy_access_setting#disable_legacy_access DisableLegacyAccessSetting#disable_legacy_access}
	DisableLegacyAccess *DisableLegacyAccessSettingDisableLegacyAccess `field:"required" json:"disableLegacyAccess" yaml:"disableLegacyAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/disable_legacy_access_setting#etag DisableLegacyAccessSetting#etag}.
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/disable_legacy_access_setting#id DisableLegacyAccessSetting#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/disable_legacy_access_setting#provider_config DisableLegacyAccessSetting#provider_config}
	ProviderConfig *DisableLegacyAccessSettingProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/disable_legacy_access_setting#setting_name DisableLegacyAccessSetting#setting_name}.
	SettingName *string `field:"optional" json:"settingName" yaml:"settingName"`
}

