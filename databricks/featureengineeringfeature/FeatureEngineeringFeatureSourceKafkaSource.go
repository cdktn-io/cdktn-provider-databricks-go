// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package featureengineeringfeature


type FeatureEngineeringFeatureSourceKafkaSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/feature_engineering_feature#name FeatureEngineeringFeature#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/feature_engineering_feature#entity_column_identifiers FeatureEngineeringFeature#entity_column_identifiers}.
	EntityColumnIdentifiers interface{} `field:"optional" json:"entityColumnIdentifiers" yaml:"entityColumnIdentifiers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/feature_engineering_feature#filter_condition FeatureEngineeringFeature#filter_condition}.
	FilterCondition *string `field:"optional" json:"filterCondition" yaml:"filterCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.114.2/docs/resources/feature_engineering_feature#timeseries_column_identifier FeatureEngineeringFeature#timeseries_column_identifier}.
	TimeseriesColumnIdentifier *FeatureEngineeringFeatureSourceKafkaSourceTimeseriesColumnIdentifier `field:"optional" json:"timeseriesColumnIdentifier" yaml:"timeseriesColumnIdentifier"`
}

