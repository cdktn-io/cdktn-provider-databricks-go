// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package systemschema


type SystemSchemaProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/system_schema#workspace_id SystemSchema#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

