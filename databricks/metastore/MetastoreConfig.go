// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package metastore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MetastoreConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/metastore#default_data_access_config_id Metastore#default_data_access_config_id}.
	DefaultDataAccessConfigId *string `field:"optional" json:"defaultDataAccessConfigId" yaml:"defaultDataAccessConfigId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/metastore#delta_sharing_organization_name Metastore#delta_sharing_organization_name}.
	DeltaSharingOrganizationName *string `field:"optional" json:"deltaSharingOrganizationName" yaml:"deltaSharingOrganizationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/metastore#delta_sharing_recipient_token_lifetime_in_seconds Metastore#delta_sharing_recipient_token_lifetime_in_seconds}.
	DeltaSharingRecipientTokenLifetimeInSeconds *float64 `field:"optional" json:"deltaSharingRecipientTokenLifetimeInSeconds" yaml:"deltaSharingRecipientTokenLifetimeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/metastore#delta_sharing_scope Metastore#delta_sharing_scope}.
	DeltaSharingScope *string `field:"optional" json:"deltaSharingScope" yaml:"deltaSharingScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/metastore#external_access_enabled Metastore#external_access_enabled}.
	ExternalAccessEnabled interface{} `field:"optional" json:"externalAccessEnabled" yaml:"externalAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/metastore#force_destroy Metastore#force_destroy}.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/metastore#id Metastore#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/metastore#name Metastore#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/metastore#owner Metastore#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/metastore#privilege_model_version Metastore#privilege_model_version}.
	PrivilegeModelVersion *string `field:"optional" json:"privilegeModelVersion" yaml:"privilegeModelVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/metastore#region Metastore#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/metastore#storage_root Metastore#storage_root}.
	StorageRoot *string `field:"optional" json:"storageRoot" yaml:"storageRoot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/metastore#storage_root_credential_id Metastore#storage_root_credential_id}.
	StorageRootCredentialId *string `field:"optional" json:"storageRootCredentialId" yaml:"storageRootCredentialId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/resources/metastore#storage_root_credential_name Metastore#storage_root_credential_name}.
	StorageRootCredentialName *string `field:"optional" json:"storageRootCredentialName" yaml:"storageRootCredentialName"`
}

