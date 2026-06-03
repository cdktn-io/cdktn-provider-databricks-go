// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accesscontrolruleset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccessControlRuleSetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/access_control_rule_set#name AccessControlRuleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies whether to use account-level or workspace-level API.
	//
	// Valid values are `account` and `workspace`. When not set, the API level is inferred from the provider host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/access_control_rule_set#api AccessControlRuleSet#api}
	Api *string `field:"optional" json:"api" yaml:"api"`
	// grant_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/access_control_rule_set#grant_rules AccessControlRuleSet#grant_rules}
	GrantRules interface{} `field:"optional" json:"grantRules" yaml:"grantRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/access_control_rule_set#id AccessControlRuleSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/access_control_rule_set#provider_config AccessControlRuleSet#provider_config}
	ProviderConfig *AccessControlRuleSetProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

