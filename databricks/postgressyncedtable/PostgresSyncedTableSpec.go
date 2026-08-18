// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgressyncedtable


type PostgresSyncedTableSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_synced_table#accelerated_sync PostgresSyncedTable#accelerated_sync}.
	AcceleratedSync interface{} `field:"optional" json:"acceleratedSync" yaml:"acceleratedSync"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_synced_table#branch PostgresSyncedTable#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_synced_table#create_database_objects_if_missing PostgresSyncedTable#create_database_objects_if_missing}.
	CreateDatabaseObjectsIfMissing interface{} `field:"optional" json:"createDatabaseObjectsIfMissing" yaml:"createDatabaseObjectsIfMissing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_synced_table#existing_pipeline_id PostgresSyncedTable#existing_pipeline_id}.
	ExistingPipelineId *string `field:"optional" json:"existingPipelineId" yaml:"existingPipelineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_synced_table#extra_columns PostgresSyncedTable#extra_columns}.
	ExtraColumns interface{} `field:"optional" json:"extraColumns" yaml:"extraColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_synced_table#new_pipeline_spec PostgresSyncedTable#new_pipeline_spec}.
	NewPipelineSpec *PostgresSyncedTableSpecNewPipelineSpec `field:"optional" json:"newPipelineSpec" yaml:"newPipelineSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_synced_table#postgres_database PostgresSyncedTable#postgres_database}.
	PostgresDatabase *string `field:"optional" json:"postgresDatabase" yaml:"postgresDatabase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_synced_table#primary_key_columns PostgresSyncedTable#primary_key_columns}.
	PrimaryKeyColumns *[]*string `field:"optional" json:"primaryKeyColumns" yaml:"primaryKeyColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_synced_table#scheduling_policy PostgresSyncedTable#scheduling_policy}.
	SchedulingPolicy *string `field:"optional" json:"schedulingPolicy" yaml:"schedulingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_synced_table#source_table_full_name PostgresSyncedTable#source_table_full_name}.
	SourceTableFullName *string `field:"optional" json:"sourceTableFullName" yaml:"sourceTableFullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_synced_table#timeseries_key PostgresSyncedTable#timeseries_key}.
	TimeseriesKey *string `field:"optional" json:"timeseriesKey" yaml:"timeseriesKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.127.0/docs/resources/postgres_synced_table#type_overrides PostgresSyncedTable#type_overrides}.
	TypeOverrides interface{} `field:"optional" json:"typeOverrides" yaml:"typeOverrides"`
}

