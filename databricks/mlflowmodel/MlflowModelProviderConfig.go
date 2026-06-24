// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mlflowmodel


type MlflowModelProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/mlflow_model#workspace_id MlflowModel#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

