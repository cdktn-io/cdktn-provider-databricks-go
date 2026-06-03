// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mlflowwebhook

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MlflowWebhookConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/mlflow_webhook#events MlflowWebhook#events}.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/mlflow_webhook#description MlflowWebhook#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// http_url_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/mlflow_webhook#http_url_spec MlflowWebhook#http_url_spec}
	HttpUrlSpec *MlflowWebhookHttpUrlSpec `field:"optional" json:"httpUrlSpec" yaml:"httpUrlSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/mlflow_webhook#id MlflowWebhook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// job_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/mlflow_webhook#job_spec MlflowWebhook#job_spec}
	JobSpec *MlflowWebhookJobSpec `field:"optional" json:"jobSpec" yaml:"jobSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/mlflow_webhook#model_name MlflowWebhook#model_name}.
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/mlflow_webhook#provider_config MlflowWebhook#provider_config}
	ProviderConfig *MlflowWebhookProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.117.0/docs/resources/mlflow_webhook#status MlflowWebhook#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

