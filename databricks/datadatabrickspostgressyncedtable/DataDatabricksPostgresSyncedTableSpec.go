// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgressyncedtable


type DataDatabricksPostgresSyncedTableSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/postgres_synced_table#accelerated_sync DataDatabricksPostgresSyncedTable#accelerated_sync}.
	AcceleratedSync interface{} `field:"optional" json:"acceleratedSync" yaml:"acceleratedSync"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/postgres_synced_table#branch DataDatabricksPostgresSyncedTable#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/postgres_synced_table#create_database_objects_if_missing DataDatabricksPostgresSyncedTable#create_database_objects_if_missing}.
	CreateDatabaseObjectsIfMissing interface{} `field:"optional" json:"createDatabaseObjectsIfMissing" yaml:"createDatabaseObjectsIfMissing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/postgres_synced_table#existing_pipeline_id DataDatabricksPostgresSyncedTable#existing_pipeline_id}.
	ExistingPipelineId *string `field:"optional" json:"existingPipelineId" yaml:"existingPipelineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/postgres_synced_table#extra_columns DataDatabricksPostgresSyncedTable#extra_columns}.
	ExtraColumns interface{} `field:"optional" json:"extraColumns" yaml:"extraColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/postgres_synced_table#new_pipeline_spec DataDatabricksPostgresSyncedTable#new_pipeline_spec}.
	NewPipelineSpec *DataDatabricksPostgresSyncedTableSpecNewPipelineSpec `field:"optional" json:"newPipelineSpec" yaml:"newPipelineSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/postgres_synced_table#postgres_database DataDatabricksPostgresSyncedTable#postgres_database}.
	PostgresDatabase *string `field:"optional" json:"postgresDatabase" yaml:"postgresDatabase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/postgres_synced_table#primary_key_columns DataDatabricksPostgresSyncedTable#primary_key_columns}.
	PrimaryKeyColumns *[]*string `field:"optional" json:"primaryKeyColumns" yaml:"primaryKeyColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/postgres_synced_table#scheduling_policy DataDatabricksPostgresSyncedTable#scheduling_policy}.
	SchedulingPolicy *string `field:"optional" json:"schedulingPolicy" yaml:"schedulingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/postgres_synced_table#source_table_full_name DataDatabricksPostgresSyncedTable#source_table_full_name}.
	SourceTableFullName *string `field:"optional" json:"sourceTableFullName" yaml:"sourceTableFullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/postgres_synced_table#timeseries_key DataDatabricksPostgresSyncedTable#timeseries_key}.
	TimeseriesKey *string `field:"optional" json:"timeseriesKey" yaml:"timeseriesKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.129.0/docs/data-sources/postgres_synced_table#type_overrides DataDatabricksPostgresSyncedTable#type_overrides}.
	TypeOverrides interface{} `field:"optional" json:"typeOverrides" yaml:"typeOverrides"`
}

