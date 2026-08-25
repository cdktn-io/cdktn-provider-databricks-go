// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountNetworkPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/account_network_policy#account_id AccountNetworkPolicy#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/account_network_policy#egress AccountNetworkPolicy#egress}.
	Egress *AccountNetworkPolicyEgress `field:"optional" json:"egress" yaml:"egress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/account_network_policy#ingress AccountNetworkPolicy#ingress}.
	Ingress *AccountNetworkPolicyIngress `field:"optional" json:"ingress" yaml:"ingress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/account_network_policy#ingress_dry_run AccountNetworkPolicy#ingress_dry_run}.
	IngressDryRun *AccountNetworkPolicyIngressDryRun `field:"optional" json:"ingressDryRun" yaml:"ingressDryRun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/account_network_policy#network_policy_id AccountNetworkPolicy#network_policy_id}.
	NetworkPolicyId *string `field:"optional" json:"networkPolicyId" yaml:"networkPolicyId"`
}

