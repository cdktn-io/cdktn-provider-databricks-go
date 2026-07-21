// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package externallocation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ExternalLocationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#credential_name ExternalLocation#credential_name}.
	CredentialName *string `field:"required" json:"credentialName" yaml:"credentialName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#name ExternalLocation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#url ExternalLocation#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#comment ExternalLocation#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// effective_file_event_queue block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#effective_file_event_queue ExternalLocation#effective_file_event_queue}
	EffectiveFileEventQueue *ExternalLocationEffectiveFileEventQueue `field:"optional" json:"effectiveFileEventQueue" yaml:"effectiveFileEventQueue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#enable_file_events ExternalLocation#enable_file_events}.
	EnableFileEvents interface{} `field:"optional" json:"enableFileEvents" yaml:"enableFileEvents"`
	// encryption_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#encryption_details ExternalLocation#encryption_details}
	EncryptionDetails *ExternalLocationEncryptionDetails `field:"optional" json:"encryptionDetails" yaml:"encryptionDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#fallback ExternalLocation#fallback}.
	Fallback interface{} `field:"optional" json:"fallback" yaml:"fallback"`
	// file_event_queue block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#file_event_queue ExternalLocation#file_event_queue}
	FileEventQueue *ExternalLocationFileEventQueue `field:"optional" json:"fileEventQueue" yaml:"fileEventQueue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#force_destroy ExternalLocation#force_destroy}.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#force_update ExternalLocation#force_update}.
	ForceUpdate interface{} `field:"optional" json:"forceUpdate" yaml:"forceUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#id ExternalLocation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#isolation_mode ExternalLocation#isolation_mode}.
	IsolationMode *string `field:"optional" json:"isolationMode" yaml:"isolationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#metastore_id ExternalLocation#metastore_id}.
	MetastoreId *string `field:"optional" json:"metastoreId" yaml:"metastoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#owner ExternalLocation#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#provider_config ExternalLocation#provider_config}
	ProviderConfig *ExternalLocationProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#read_only ExternalLocation#read_only}.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.122.0/docs/resources/external_location#skip_validation ExternalLocation#skip_validation}.
	SkipValidation interface{} `field:"optional" json:"skipValidation" yaml:"skipValidation"`
}

