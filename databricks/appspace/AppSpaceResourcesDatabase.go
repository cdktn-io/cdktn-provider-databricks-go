// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appspace


type AppSpaceResourcesDatabase struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/app_space#database_name AppSpace#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/app_space#instance_name AppSpace#instance_name}.
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/app_space#permission AppSpace#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

