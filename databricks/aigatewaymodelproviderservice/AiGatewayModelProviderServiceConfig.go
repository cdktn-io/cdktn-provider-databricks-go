// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aigatewaymodelproviderservice

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiGatewayModelProviderServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/ai_gateway_model_provider_service#model_provider_service_id AiGatewayModelProviderService#model_provider_service_id}.
	ModelProviderServiceId *string `field:"required" json:"modelProviderServiceId" yaml:"modelProviderServiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/ai_gateway_model_provider_service#parent AiGatewayModelProviderService#parent}.
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/ai_gateway_model_provider_service#comment AiGatewayModelProviderService#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/ai_gateway_model_provider_service#config AiGatewayModelProviderService#config}.
	Config *AiGatewayModelProviderServiceConfigA `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/ai_gateway_model_provider_service#owner AiGatewayModelProviderService#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.128.0/docs/resources/ai_gateway_model_provider_service#provider_config AiGatewayModelProviderService#provider_config}.
	ProviderConfig *AiGatewayModelProviderServiceProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
}

