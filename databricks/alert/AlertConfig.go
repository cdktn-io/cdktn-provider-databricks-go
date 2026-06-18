// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package alert

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AlertConfig struct {
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
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/alert#condition Alert#condition}
	Condition *AlertCondition `field:"required" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/alert#display_name Alert#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/alert#query_id Alert#query_id}.
	QueryId *string `field:"required" json:"queryId" yaml:"queryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/alert#custom_body Alert#custom_body}.
	CustomBody *string `field:"optional" json:"customBody" yaml:"customBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/alert#custom_subject Alert#custom_subject}.
	CustomSubject *string `field:"optional" json:"customSubject" yaml:"customSubject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/alert#notify_on_ok Alert#notify_on_ok}.
	NotifyOnOk interface{} `field:"optional" json:"notifyOnOk" yaml:"notifyOnOk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/alert#owner_user_name Alert#owner_user_name}.
	OwnerUserName *string `field:"optional" json:"ownerUserName" yaml:"ownerUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/alert#parent_path Alert#parent_path}.
	ParentPath *string `field:"optional" json:"parentPath" yaml:"parentPath"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/alert#provider_config Alert#provider_config}
	ProviderConfig *AlertProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/alert#seconds_to_retrigger Alert#seconds_to_retrigger}.
	SecondsToRetrigger *float64 `field:"optional" json:"secondsToRetrigger" yaml:"secondsToRetrigger"`
}

