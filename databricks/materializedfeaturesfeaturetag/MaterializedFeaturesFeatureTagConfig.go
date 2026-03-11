// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package materializedfeaturesfeaturetag

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MaterializedFeaturesFeatureTagConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.111.0/docs/resources/materialized_features_feature_tag#key MaterializedFeaturesFeatureTag#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.111.0/docs/resources/materialized_features_feature_tag#provider_config MaterializedFeaturesFeatureTag#provider_config}.
	ProviderConfig *MaterializedFeaturesFeatureTagProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.111.0/docs/resources/materialized_features_feature_tag#value MaterializedFeaturesFeatureTag#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

