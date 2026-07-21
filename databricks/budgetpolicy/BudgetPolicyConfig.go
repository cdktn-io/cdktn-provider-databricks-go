// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budgetpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BudgetPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/budget_policy#binding_workspace_ids BudgetPolicy#binding_workspace_ids}.
	BindingWorkspaceIds *[]*float64 `field:"optional" json:"bindingWorkspaceIds" yaml:"bindingWorkspaceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/budget_policy#custom_tags BudgetPolicy#custom_tags}.
	CustomTags interface{} `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/budget_policy#policy_id BudgetPolicy#policy_id}.
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/budget_policy#policy_name BudgetPolicy#policy_name}.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

