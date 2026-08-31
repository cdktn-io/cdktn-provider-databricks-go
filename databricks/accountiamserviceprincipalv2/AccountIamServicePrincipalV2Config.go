// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountiamserviceprincipalv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountIamServicePrincipalV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/account_iam_service_principal_v2#account_sp_status AccountIamServicePrincipalV2#account_sp_status}.
	AccountSpStatus *string `field:"required" json:"accountSpStatus" yaml:"accountSpStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/account_iam_service_principal_v2#display_name AccountIamServicePrincipalV2#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/account_iam_service_principal_v2#application_id AccountIamServicePrincipalV2#application_id}.
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.130.0/docs/resources/account_iam_service_principal_v2#external_id AccountIamServicePrincipalV2#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
}

