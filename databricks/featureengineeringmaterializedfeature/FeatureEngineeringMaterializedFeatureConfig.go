// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringmaterializedfeature

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FeatureEngineeringMaterializedFeatureConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/feature_engineering_materialized_feature#feature_name FeatureEngineeringMaterializedFeature#feature_name}.
	FeatureName *string `field:"required" json:"featureName" yaml:"featureName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/feature_engineering_materialized_feature#cron_schedule FeatureEngineeringMaterializedFeature#cron_schedule}.
	CronSchedule *string `field:"optional" json:"cronSchedule" yaml:"cronSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/feature_engineering_materialized_feature#cron_schedule_trigger FeatureEngineeringMaterializedFeature#cron_schedule_trigger}.
	CronScheduleTrigger *FeatureEngineeringMaterializedFeatureCronScheduleTrigger `field:"optional" json:"cronScheduleTrigger" yaml:"cronScheduleTrigger"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/feature_engineering_materialized_feature#offline_store_config FeatureEngineeringMaterializedFeature#offline_store_config}.
	OfflineStoreConfig *FeatureEngineeringMaterializedFeatureOfflineStoreConfig `field:"optional" json:"offlineStoreConfig" yaml:"offlineStoreConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/feature_engineering_materialized_feature#online_store_config FeatureEngineeringMaterializedFeature#online_store_config}.
	OnlineStoreConfig *FeatureEngineeringMaterializedFeatureOnlineStoreConfig `field:"optional" json:"onlineStoreConfig" yaml:"onlineStoreConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/feature_engineering_materialized_feature#pipeline_schedule_state FeatureEngineeringMaterializedFeature#pipeline_schedule_state}.
	PipelineScheduleState *string `field:"optional" json:"pipelineScheduleState" yaml:"pipelineScheduleState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/feature_engineering_materialized_feature#provider_config FeatureEngineeringMaterializedFeature#provider_config}.
	ProviderConfig *FeatureEngineeringMaterializedFeatureProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/feature_engineering_materialized_feature#streaming_mode FeatureEngineeringMaterializedFeature#streaming_mode}.
	StreamingMode *FeatureEngineeringMaterializedFeatureStreamingMode `field:"optional" json:"streamingMode" yaml:"streamingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.119.0/docs/resources/feature_engineering_materialized_feature#table_trigger FeatureEngineeringMaterializedFeature#table_trigger}.
	TableTrigger *FeatureEngineeringMaterializedFeatureTableTrigger `field:"optional" json:"tableTrigger" yaml:"tableTrigger"`
}

