// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksrfaaccessrequestdestinations

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksRfaAccessRequestDestinationsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/rfa_access_request_destinations#full_name DataDatabricksRfaAccessRequestDestinations#full_name}.
	FullName *string `field:"required" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/rfa_access_request_destinations#securable_type DataDatabricksRfaAccessRequestDestinations#securable_type}.
	SecurableType *string `field:"required" json:"securableType" yaml:"securableType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/rfa_access_request_destinations#provider_config DataDatabricksRfaAccessRequestDestinations#provider_config}.
	ProviderConfig *DataDatabricksRfaAccessRequestDestinationsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

