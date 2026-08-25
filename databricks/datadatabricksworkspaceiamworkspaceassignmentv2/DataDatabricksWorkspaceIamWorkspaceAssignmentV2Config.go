// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspaceiamworkspaceassignmentv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksWorkspaceIamWorkspaceAssignmentV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/workspace_iam_workspace_assignment_v2#principal_id DataDatabricksWorkspaceIamWorkspaceAssignmentV2#principal_id}.
	PrincipalId *float64 `field:"required" json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/workspace_iam_workspace_assignment_v2#provider_config DataDatabricksWorkspaceIamWorkspaceAssignmentV2#provider_config}.
	ProviderConfig *DataDatabricksWorkspaceIamWorkspaceAssignmentV2ProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

