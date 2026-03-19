// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfeatureengineeringfeatures


type DataDatabricksFeatureEngineeringFeaturesFeaturesSourceKafkaSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/feature_engineering_features#name DataDatabricksFeatureEngineeringFeatures#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/feature_engineering_features#entity_column_identifiers DataDatabricksFeatureEngineeringFeatures#entity_column_identifiers}.
	EntityColumnIdentifiers interface{} `field:"optional" json:"entityColumnIdentifiers" yaml:"entityColumnIdentifiers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/feature_engineering_features#filter_condition DataDatabricksFeatureEngineeringFeatures#filter_condition}.
	FilterCondition *string `field:"optional" json:"filterCondition" yaml:"filterCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.112.0/docs/data-sources/feature_engineering_features#timeseries_column_identifier DataDatabricksFeatureEngineeringFeatures#timeseries_column_identifier}.
	TimeseriesColumnIdentifier *DataDatabricksFeatureEngineeringFeaturesFeaturesSourceKafkaSourceTimeseriesColumnIdentifier `field:"optional" json:"timeseriesColumnIdentifier" yaml:"timeseriesColumnIdentifier"`
}

