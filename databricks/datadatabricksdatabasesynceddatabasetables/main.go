// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasesynceddatabasetables

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTables",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTables)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceName", GoGetter: "InstanceName"},
			_jsii_.MemberProperty{JsiiProperty: "instanceNameInput", GoGetter: "InstanceNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pageSize", GoGetter: "PageSize"},
			_jsii_.MemberProperty{JsiiProperty: "pageSizeInput", GoGetter: "PageSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "providerConfig", GoGetter: "ProviderConfig"},
			_jsii_.MemberProperty{JsiiProperty: "providerConfigInput", GoGetter: "ProviderConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putProviderConfig", GoMethod: "PutProviderConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPageSize", GoMethod: "ResetPageSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetProviderConfig", GoMethod: "ResetProviderConfig"},
			_jsii_.MemberProperty{JsiiProperty: "syncedTables", GoGetter: "SyncedTables"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTables{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesConfig",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesProviderConfig",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesProviderConfigOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesProviderConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceId", GoMethod: "ResetWorkspaceId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceId", GoGetter: "WorkspaceId"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceIdInput", GoGetter: "WorkspaceIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesProviderConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTables",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTables)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatus",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusContinuousUpdateStatus",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusContinuousUpdateStatus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusContinuousUpdateStatusInitialPipelineSyncProgress",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusContinuousUpdateStatusInitialPipelineSyncProgress)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusContinuousUpdateStatusInitialPipelineSyncProgressOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusContinuousUpdateStatusInitialPipelineSyncProgressOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "estimatedCompletionTimeSeconds", GoGetter: "EstimatedCompletionTimeSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersionCurrentlyProcessing", GoGetter: "LatestVersionCurrentlyProcessing"},
			_jsii_.MemberProperty{JsiiProperty: "provisioningPhase", GoGetter: "ProvisioningPhase"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "syncedRowCount", GoGetter: "SyncedRowCount"},
			_jsii_.MemberProperty{JsiiProperty: "syncProgressCompletion", GoGetter: "SyncProgressCompletion"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "totalRowCount", GoGetter: "TotalRowCount"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusContinuousUpdateStatusInitialPipelineSyncProgressOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusContinuousUpdateStatusOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusContinuousUpdateStatusOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "initialPipelineSyncProgress", GoGetter: "InitialPipelineSyncProgress"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastProcessedCommitVersion", GoGetter: "LastProcessedCommitVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timestamp", GoGetter: "Timestamp"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusContinuousUpdateStatusOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusFailedStatus",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusFailedStatus)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusFailedStatusOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusFailedStatusOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastProcessedCommitVersion", GoGetter: "LastProcessedCommitVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timestamp", GoGetter: "Timestamp"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusFailedStatusOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusLastSync",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusLastSync)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusLastSyncDeltaTableSyncInfo",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusLastSyncDeltaTableSyncInfo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusLastSyncDeltaTableSyncInfoOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusLastSyncDeltaTableSyncInfoOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deltaCommitTimestamp", GoGetter: "DeltaCommitTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "deltaCommitVersion", GoGetter: "DeltaCommitVersion"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusLastSyncDeltaTableSyncInfoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusLastSyncOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusLastSyncOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deltaTableSyncInfo", GoGetter: "DeltaTableSyncInfo"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "syncEndTimestamp", GoGetter: "SyncEndTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "syncStartTimestamp", GoGetter: "SyncStartTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusLastSyncOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "continuousUpdateStatus", GoGetter: "ContinuousUpdateStatus"},
			_jsii_.MemberProperty{JsiiProperty: "continuousUpdateStatusInput", GoGetter: "ContinuousUpdateStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "detailedState", GoGetter: "DetailedState"},
			_jsii_.MemberProperty{JsiiProperty: "failedStatus", GoGetter: "FailedStatus"},
			_jsii_.MemberProperty{JsiiProperty: "failedStatusInput", GoGetter: "FailedStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastSync", GoGetter: "LastSync"},
			_jsii_.MemberProperty{JsiiProperty: "message", GoGetter: "Message"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineId", GoGetter: "PipelineId"},
			_jsii_.MemberProperty{JsiiProperty: "provisioningStatus", GoGetter: "ProvisioningStatus"},
			_jsii_.MemberProperty{JsiiProperty: "provisioningStatusInput", GoGetter: "ProvisioningStatusInput"},
			_jsii_.MemberMethod{JsiiMethod: "putContinuousUpdateStatus", GoMethod: "PutContinuousUpdateStatus"},
			_jsii_.MemberMethod{JsiiMethod: "putFailedStatus", GoMethod: "PutFailedStatus"},
			_jsii_.MemberMethod{JsiiMethod: "putProvisioningStatus", GoMethod: "PutProvisioningStatus"},
			_jsii_.MemberMethod{JsiiMethod: "putTriggeredUpdateStatus", GoMethod: "PutTriggeredUpdateStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetContinuousUpdateStatus", GoMethod: "ResetContinuousUpdateStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailedStatus", GoMethod: "ResetFailedStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetProvisioningStatus", GoMethod: "ResetProvisioningStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetTriggeredUpdateStatus", GoMethod: "ResetTriggeredUpdateStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggeredUpdateStatus", GoGetter: "TriggeredUpdateStatus"},
			_jsii_.MemberProperty{JsiiProperty: "triggeredUpdateStatusInput", GoGetter: "TriggeredUpdateStatusInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusProvisioningStatus",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusProvisioningStatus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusProvisioningStatusInitialPipelineSyncProgress",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusProvisioningStatusInitialPipelineSyncProgress)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusProvisioningStatusInitialPipelineSyncProgressOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusProvisioningStatusInitialPipelineSyncProgressOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "estimatedCompletionTimeSeconds", GoGetter: "EstimatedCompletionTimeSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersionCurrentlyProcessing", GoGetter: "LatestVersionCurrentlyProcessing"},
			_jsii_.MemberProperty{JsiiProperty: "provisioningPhase", GoGetter: "ProvisioningPhase"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "syncedRowCount", GoGetter: "SyncedRowCount"},
			_jsii_.MemberProperty{JsiiProperty: "syncProgressCompletion", GoGetter: "SyncProgressCompletion"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "totalRowCount", GoGetter: "TotalRowCount"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusProvisioningStatusInitialPipelineSyncProgressOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusProvisioningStatusOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusProvisioningStatusOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "initialPipelineSyncProgress", GoGetter: "InitialPipelineSyncProgress"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusProvisioningStatusOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusTriggeredUpdateStatus",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusTriggeredUpdateStatus)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusTriggeredUpdateStatusOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusTriggeredUpdateStatusOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastProcessedCommitVersion", GoGetter: "LastProcessedCommitVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timestamp", GoGetter: "Timestamp"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggeredUpdateProgress", GoGetter: "TriggeredUpdateProgress"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusTriggeredUpdateStatusOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusTriggeredUpdateStatusTriggeredUpdateProgress",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusTriggeredUpdateStatusTriggeredUpdateProgress)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusTriggeredUpdateStatusTriggeredUpdateProgressOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusTriggeredUpdateStatusTriggeredUpdateProgressOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "estimatedCompletionTimeSeconds", GoGetter: "EstimatedCompletionTimeSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersionCurrentlyProcessing", GoGetter: "LatestVersionCurrentlyProcessing"},
			_jsii_.MemberProperty{JsiiProperty: "provisioningPhase", GoGetter: "ProvisioningPhase"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "syncedRowCount", GoGetter: "SyncedRowCount"},
			_jsii_.MemberProperty{JsiiProperty: "syncProgressCompletion", GoGetter: "SyncProgressCompletion"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "totalRowCount", GoGetter: "TotalRowCount"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusTriggeredUpdateStatusTriggeredUpdateProgressOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesList",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "databaseInstanceName", GoGetter: "DatabaseInstanceName"},
			_jsii_.MemberProperty{JsiiProperty: "dataSynchronizationStatus", GoGetter: "DataSynchronizationStatus"},
			_jsii_.MemberProperty{JsiiProperty: "effectiveDatabaseInstanceName", GoGetter: "EffectiveDatabaseInstanceName"},
			_jsii_.MemberProperty{JsiiProperty: "effectiveLogicalDatabaseName", GoGetter: "EffectiveLogicalDatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "logicalDatabaseName", GoGetter: "LogicalDatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "providerConfig", GoGetter: "ProviderConfig"},
			_jsii_.MemberProperty{JsiiProperty: "providerConfigInput", GoGetter: "ProviderConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putProviderConfig", GoMethod: "PutProviderConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetProviderConfig", GoMethod: "ResetProviderConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unityCatalogProvisioningState", GoGetter: "UnityCatalogProvisioningState"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesProviderConfig",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesProviderConfigOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesProviderConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceId", GoMethod: "ResetWorkspaceId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceId", GoGetter: "WorkspaceId"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceIdInput", GoGetter: "WorkspaceIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesProviderConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpec",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecNewPipelineSpec",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecNewPipelineSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecNewPipelineSpecOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecNewPipelineSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "budgetPolicyId", GoGetter: "BudgetPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "budgetPolicyIdInput", GoGetter: "BudgetPolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetBudgetPolicyId", GoMethod: "ResetBudgetPolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageCatalog", GoMethod: "ResetStorageCatalog"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageSchema", GoMethod: "ResetStorageSchema"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "storageCatalog", GoGetter: "StorageCatalog"},
			_jsii_.MemberProperty{JsiiProperty: "storageCatalogInput", GoGetter: "StorageCatalogInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageSchema", GoGetter: "StorageSchema"},
			_jsii_.MemberProperty{JsiiProperty: "storageSchemaInput", GoGetter: "StorageSchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecNewPipelineSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acceleratedSync", GoGetter: "AcceleratedSync"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratedSyncInput", GoGetter: "AcceleratedSyncInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "createDatabaseObjectsIfMissing", GoGetter: "CreateDatabaseObjectsIfMissing"},
			_jsii_.MemberProperty{JsiiProperty: "createDatabaseObjectsIfMissingInput", GoGetter: "CreateDatabaseObjectsIfMissingInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "existingPipelineId", GoGetter: "ExistingPipelineId"},
			_jsii_.MemberProperty{JsiiProperty: "existingPipelineIdInput", GoGetter: "ExistingPipelineIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "newPipelineSpec", GoGetter: "NewPipelineSpec"},
			_jsii_.MemberProperty{JsiiProperty: "newPipelineSpecInput", GoGetter: "NewPipelineSpecInput"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKeyColumns", GoGetter: "PrimaryKeyColumns"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKeyColumnsInput", GoGetter: "PrimaryKeyColumnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putNewPipelineSpec", GoMethod: "PutNewPipelineSpec"},
			_jsii_.MemberMethod{JsiiMethod: "putTypeOverrides", GoMethod: "PutTypeOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcceleratedSync", GoMethod: "ResetAcceleratedSync"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreateDatabaseObjectsIfMissing", GoMethod: "ResetCreateDatabaseObjectsIfMissing"},
			_jsii_.MemberMethod{JsiiMethod: "resetExistingPipelineId", GoMethod: "ResetExistingPipelineId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNewPipelineSpec", GoMethod: "ResetNewPipelineSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimaryKeyColumns", GoMethod: "ResetPrimaryKeyColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetSchedulingPolicy", GoMethod: "ResetSchedulingPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceTableFullName", GoMethod: "ResetSourceTableFullName"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeseriesKey", GoMethod: "ResetTimeseriesKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetTypeOverrides", GoMethod: "ResetTypeOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPolicy", GoGetter: "SchedulingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPolicyInput", GoGetter: "SchedulingPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceTableFullName", GoGetter: "SourceTableFullName"},
			_jsii_.MemberProperty{JsiiProperty: "sourceTableFullNameInput", GoGetter: "SourceTableFullNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeseriesKey", GoGetter: "TimeseriesKey"},
			_jsii_.MemberProperty{JsiiProperty: "timeseriesKeyInput", GoGetter: "TimeseriesKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "typeOverrides", GoGetter: "TypeOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "typeOverridesInput", GoGetter: "TypeOverridesInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverrides",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverrides)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesList",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference",
		reflect.TypeOf((*DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "columnName", GoGetter: "ColumnName"},
			_jsii_.MemberProperty{JsiiProperty: "columnNameInput", GoGetter: "ColumnNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "pgType", GoGetter: "PgType"},
			_jsii_.MemberProperty{JsiiProperty: "pgTypeInput", GoGetter: "PgTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetSize", GoMethod: "ResetSize"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "size", GoGetter: "Size"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInput", GoGetter: "SizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecTypeOverridesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
