// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package globalinitscript

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlobalInitScriptConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/global_init_script#name GlobalInitScript#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/global_init_script#content_base64 GlobalInitScript#content_base64}.
	ContentBase64 *string `field:"optional" json:"contentBase64" yaml:"contentBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/global_init_script#enabled GlobalInitScript#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/global_init_script#id GlobalInitScript#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/global_init_script#md5 GlobalInitScript#md5}.
	Md5 *string `field:"optional" json:"md5" yaml:"md5"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/global_init_script#position GlobalInitScript#position}.
	Position *float64 `field:"optional" json:"position" yaml:"position"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/global_init_script#provider_config GlobalInitScript#provider_config}
	ProviderConfig *GlobalInitScriptProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/global_init_script#source GlobalInitScript#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/global_init_script#timeouts GlobalInitScript#timeouts}
	Timeouts *GlobalInitScriptTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

