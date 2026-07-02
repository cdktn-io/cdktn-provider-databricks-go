// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rfaaccessrequestdestinations

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RfaAccessRequestDestinationsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/rfa_access_request_destinations#securable RfaAccessRequestDestinations#securable}.
	Securable *RfaAccessRequestDestinationsSecurable `field:"required" json:"securable" yaml:"securable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/rfa_access_request_destinations#destinations RfaAccessRequestDestinations#destinations}.
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/rfa_access_request_destinations#provider_config RfaAccessRequestDestinations#provider_config}.
	ProviderConfig *RfaAccessRequestDestinationsProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

