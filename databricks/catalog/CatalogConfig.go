// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalog

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CatalogConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#browse_only Catalog#browse_only}.
	BrowseOnly interface{} `field:"optional" json:"browseOnly" yaml:"browseOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#comment Catalog#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#connection_name Catalog#connection_name}.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#custom_max_retention_hours Catalog#custom_max_retention_hours}.
	CustomMaxRetentionHours *float64 `field:"optional" json:"customMaxRetentionHours" yaml:"customMaxRetentionHours"`
	// effective_predictive_optimization_flag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#effective_predictive_optimization_flag Catalog#effective_predictive_optimization_flag}
	EffectivePredictiveOptimizationFlag *CatalogEffectivePredictiveOptimizationFlag `field:"optional" json:"effectivePredictiveOptimizationFlag" yaml:"effectivePredictiveOptimizationFlag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#enable_predictive_optimization Catalog#enable_predictive_optimization}.
	EnablePredictiveOptimization *string `field:"optional" json:"enablePredictiveOptimization" yaml:"enablePredictiveOptimization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#force_destroy Catalog#force_destroy}.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#id Catalog#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#isolation_mode Catalog#isolation_mode}.
	IsolationMode *string `field:"optional" json:"isolationMode" yaml:"isolationMode"`
	// managed_encryption_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#managed_encryption_settings Catalog#managed_encryption_settings}
	ManagedEncryptionSettings *CatalogManagedEncryptionSettings `field:"optional" json:"managedEncryptionSettings" yaml:"managedEncryptionSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#metastore_id Catalog#metastore_id}.
	MetastoreId *string `field:"optional" json:"metastoreId" yaml:"metastoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#name Catalog#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#options Catalog#options}.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#owner Catalog#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#properties Catalog#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#provider_config Catalog#provider_config}
	ProviderConfig *CatalogProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#provider_name Catalog#provider_name}.
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
	// provisioning_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#provisioning_info Catalog#provisioning_info}
	ProvisioningInfo *CatalogProvisioningInfo `field:"optional" json:"provisioningInfo" yaml:"provisioningInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#share_name Catalog#share_name}.
	ShareName *string `field:"optional" json:"shareName" yaml:"shareName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.121.0/docs/resources/catalog#storage_root Catalog#storage_root}.
	StorageRoot *string `field:"optional" json:"storageRoot" yaml:"storageRoot"`
}

