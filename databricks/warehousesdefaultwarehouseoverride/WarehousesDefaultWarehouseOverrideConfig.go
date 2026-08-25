// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package warehousesdefaultwarehouseoverride

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WarehousesDefaultWarehouseOverrideConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/warehouses_default_warehouse_override#default_warehouse_override_id WarehousesDefaultWarehouseOverride#default_warehouse_override_id}.
	DefaultWarehouseOverrideId *string `field:"required" json:"defaultWarehouseOverrideId" yaml:"defaultWarehouseOverrideId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/warehouses_default_warehouse_override#type WarehousesDefaultWarehouseOverride#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/warehouses_default_warehouse_override#provider_config WarehousesDefaultWarehouseOverride#provider_config}.
	ProviderConfig *WarehousesDefaultWarehouseOverrideProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/warehouses_default_warehouse_override#warehouse_id WarehousesDefaultWarehouseOverride#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

