// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package budget


type BudgetFilter struct {
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.111.0/docs/resources/budget#tags Budget#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// workspace_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.111.0/docs/resources/budget#workspace_id Budget#workspace_id}
	WorkspaceId *BudgetFilterWorkspaceId `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

