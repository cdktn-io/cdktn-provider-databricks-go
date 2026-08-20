// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package disasterrecoverystableurl

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DisasterRecoveryStableUrlConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/disaster_recovery_stable_url#initial_workspace_id DisasterRecoveryStableUrl#initial_workspace_id}.
	InitialWorkspaceId *string `field:"required" json:"initialWorkspaceId" yaml:"initialWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/disaster_recovery_stable_url#parent DisasterRecoveryStableUrl#parent}.
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/disaster_recovery_stable_url#stable_url_id DisasterRecoveryStableUrl#stable_url_id}.
	StableUrlId *string `field:"required" json:"stableUrlId" yaml:"stableUrlId"`
}

