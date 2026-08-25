// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package serviceprincipalsecret

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicePrincipalSecretConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/service_principal_secret#service_principal_id ServicePrincipalSecret#service_principal_id}.
	ServicePrincipalId *string `field:"required" json:"servicePrincipalId" yaml:"servicePrincipalId"`
	// Specifies whether to use account-level or workspace-level API.
	//
	// Valid values are `account` and `workspace`. When not set, the API level is inferred from the provider host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/service_principal_secret#api ServicePrincipalSecret#api}
	Api *string `field:"optional" json:"api" yaml:"api"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/service_principal_secret#create_time ServicePrincipalSecret#create_time}.
	CreateTime *string `field:"optional" json:"createTime" yaml:"createTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/service_principal_secret#expire_time ServicePrincipalSecret#expire_time}.
	ExpireTime *string `field:"optional" json:"expireTime" yaml:"expireTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/service_principal_secret#id ServicePrincipalSecret#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/service_principal_secret#lifetime ServicePrincipalSecret#lifetime}.
	Lifetime *string `field:"optional" json:"lifetime" yaml:"lifetime"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/service_principal_secret#provider_config ServicePrincipalSecret#provider_config}
	ProviderConfig *ServicePrincipalSecretProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/service_principal_secret#secret ServicePrincipalSecret#secret}.
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/service_principal_secret#secret_hash ServicePrincipalSecret#secret_hash}.
	SecretHash *string `field:"optional" json:"secretHash" yaml:"secretHash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/service_principal_secret#status ServicePrincipalSecret#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/service_principal_secret#time_rotating ServicePrincipalSecret#time_rotating}.
	TimeRotating *string `field:"optional" json:"timeRotating" yaml:"timeRotating"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/resources/service_principal_secret#update_time ServicePrincipalSecret#update_time}.
	UpdateTime *string `field:"optional" json:"updateTime" yaml:"updateTime"`
}

