// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mlflowwebhook


type MlflowWebhookProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.110.0/docs/resources/mlflow_webhook#workspace_id MlflowWebhook#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

