// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringkafkaconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FeatureEngineeringKafkaConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_kafka_config#auth_config FeatureEngineeringKafkaConfig#auth_config}.
	AuthConfig *FeatureEngineeringKafkaConfigAuthConfig `field:"required" json:"authConfig" yaml:"authConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_kafka_config#bootstrap_servers FeatureEngineeringKafkaConfig#bootstrap_servers}.
	BootstrapServers *string `field:"required" json:"bootstrapServers" yaml:"bootstrapServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_kafka_config#subscription_mode FeatureEngineeringKafkaConfig#subscription_mode}.
	SubscriptionMode *FeatureEngineeringKafkaConfigSubscriptionMode `field:"required" json:"subscriptionMode" yaml:"subscriptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_kafka_config#backfill_source FeatureEngineeringKafkaConfig#backfill_source}.
	BackfillSource *FeatureEngineeringKafkaConfigBackfillSource `field:"optional" json:"backfillSource" yaml:"backfillSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_kafka_config#extra_options FeatureEngineeringKafkaConfig#extra_options}.
	ExtraOptions *map[string]*string `field:"optional" json:"extraOptions" yaml:"extraOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_kafka_config#ingestion_config FeatureEngineeringKafkaConfig#ingestion_config}.
	IngestionConfig *FeatureEngineeringKafkaConfigIngestionConfig `field:"optional" json:"ingestionConfig" yaml:"ingestionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_kafka_config#key_schema FeatureEngineeringKafkaConfig#key_schema}.
	KeySchema *FeatureEngineeringKafkaConfigKeySchema `field:"optional" json:"keySchema" yaml:"keySchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_kafka_config#provider_config FeatureEngineeringKafkaConfig#provider_config}.
	ProviderConfig *FeatureEngineeringKafkaConfigProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.120.0/docs/resources/feature_engineering_kafka_config#value_schema FeatureEngineeringKafkaConfig#value_schema}.
	ValueSchema *FeatureEngineeringKafkaConfigValueSchema `field:"optional" json:"valueSchema" yaml:"valueSchema"`
}

