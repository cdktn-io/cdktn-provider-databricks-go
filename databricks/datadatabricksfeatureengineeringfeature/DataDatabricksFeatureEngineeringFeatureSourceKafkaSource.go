// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeature


type DataDatabricksFeatureEngineeringFeatureSourceKafkaSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/feature_engineering_feature#name DataDatabricksFeatureEngineeringFeature#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/feature_engineering_feature#entity_column_identifiers DataDatabricksFeatureEngineeringFeature#entity_column_identifiers}.
	EntityColumnIdentifiers interface{} `field:"optional" json:"entityColumnIdentifiers" yaml:"entityColumnIdentifiers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/feature_engineering_feature#filter_condition DataDatabricksFeatureEngineeringFeature#filter_condition}.
	FilterCondition *string `field:"optional" json:"filterCondition" yaml:"filterCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/feature_engineering_feature#timeseries_column_identifier DataDatabricksFeatureEngineeringFeature#timeseries_column_identifier}.
	TimeseriesColumnIdentifier *DataDatabricksFeatureEngineeringFeatureSourceKafkaSourceTimeseriesColumnIdentifier `field:"optional" json:"timeseriesColumnIdentifier" yaml:"timeseriesColumnIdentifier"`
}

