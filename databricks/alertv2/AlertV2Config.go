// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package alertv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AlertV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/alert_v2#display_name AlertV2#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/alert_v2#evaluation AlertV2#evaluation}.
	Evaluation *AlertV2Evaluation `field:"required" json:"evaluation" yaml:"evaluation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/alert_v2#query_text AlertV2#query_text}.
	QueryText *string `field:"required" json:"queryText" yaml:"queryText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/alert_v2#schedule AlertV2#schedule}.
	Schedule *AlertV2Schedule `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/alert_v2#warehouse_id AlertV2#warehouse_id}.
	WarehouseId *string `field:"required" json:"warehouseId" yaml:"warehouseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/alert_v2#custom_description AlertV2#custom_description}.
	CustomDescription *string `field:"optional" json:"customDescription" yaml:"customDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/alert_v2#custom_summary AlertV2#custom_summary}.
	CustomSummary *string `field:"optional" json:"customSummary" yaml:"customSummary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/alert_v2#parent_path AlertV2#parent_path}.
	ParentPath *string `field:"optional" json:"parentPath" yaml:"parentPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/alert_v2#provider_config AlertV2#provider_config}.
	ProviderConfig *AlertV2ProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/alert_v2#purge_on_delete AlertV2#purge_on_delete}.
	PurgeOnDelete interface{} `field:"optional" json:"purgeOnDelete" yaml:"purgeOnDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/alert_v2#run_as AlertV2#run_as}.
	RunAs *AlertV2RunAs `field:"optional" json:"runAs" yaml:"runAs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/alert_v2#run_as_user_name AlertV2#run_as_user_name}.
	RunAsUserName *string `field:"optional" json:"runAsUserName" yaml:"runAsUserName"`
}

