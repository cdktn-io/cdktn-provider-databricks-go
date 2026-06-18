// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPolicies",
		reflect.TypeOf((*DataDatabricksAccountNetworkPolicies)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
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
			j := jsiiProxy_DataDatabricksAccountNetworkPolicies{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesConfig",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItems",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgress",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccess",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccess)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedDatabricksDestinations",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedDatabricksDestinations)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedDatabricksDestinationsList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedDatabricksDestinationsList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedDatabricksDestinationsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedDatabricksDestinationsOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedDatabricksDestinationsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceIds", GoMethod: "ResetWorkspaceIds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceIds", GoGetter: "WorkspaceIds"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceIdsInput", GoGetter: "WorkspaceIdsInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedDatabricksDestinationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedInternetDestinations",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedInternetDestinations)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedInternetDestinationsList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedInternetDestinationsList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedInternetDestinationsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedInternetDestinationsOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedInternetDestinationsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internetDestinationType", GoGetter: "InternetDestinationType"},
			_jsii_.MemberProperty{JsiiProperty: "internetDestinationTypeInput", GoGetter: "InternetDestinationTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resetInternetDestinationType", GoMethod: "ResetInternetDestinationType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedInternetDestinationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedStorageDestinations",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedStorageDestinations)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedStorageDestinationsList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedStorageDestinationsList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedStorageDestinationsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedStorageDestinationsOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedStorageDestinationsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "azureStorageAccount", GoGetter: "AzureStorageAccount"},
			_jsii_.MemberProperty{JsiiProperty: "azureStorageAccountInput", GoGetter: "AzureStorageAccountInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureStorageService", GoGetter: "AzureStorageService"},
			_jsii_.MemberProperty{JsiiProperty: "azureStorageServiceInput", GoGetter: "AzureStorageServiceInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketNameInput", GoGetter: "BucketNameInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureStorageAccount", GoMethod: "ResetAzureStorageAccount"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureStorageService", GoMethod: "ResetAzureStorageService"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketName", GoMethod: "ResetBucketName"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageDestinationType", GoMethod: "ResetStorageDestinationType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "storageDestinationType", GoGetter: "StorageDestinationType"},
			_jsii_.MemberProperty{JsiiProperty: "storageDestinationTypeInput", GoGetter: "StorageDestinationTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedStorageDestinationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessBlockedInternetDestinations",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessBlockedInternetDestinations)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessBlockedInternetDestinationsList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessBlockedInternetDestinationsList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessBlockedInternetDestinationsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessBlockedInternetDestinationsOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessBlockedInternetDestinationsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internetDestinationType", GoGetter: "InternetDestinationType"},
			_jsii_.MemberProperty{JsiiProperty: "internetDestinationTypeInput", GoGetter: "InternetDestinationTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resetInternetDestinationType", GoMethod: "ResetInternetDestinationType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessBlockedInternetDestinationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowedDatabricksDestinations", GoGetter: "AllowedDatabricksDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allowedDatabricksDestinationsInput", GoGetter: "AllowedDatabricksDestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedInternetDestinations", GoGetter: "AllowedInternetDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allowedInternetDestinationsInput", GoGetter: "AllowedInternetDestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedStorageDestinations", GoGetter: "AllowedStorageDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allowedStorageDestinationsInput", GoGetter: "AllowedStorageDestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "blockedInternetDestinations", GoGetter: "BlockedInternetDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "blockedInternetDestinationsInput", GoGetter: "BlockedInternetDestinationsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "policyEnforcement", GoGetter: "PolicyEnforcement"},
			_jsii_.MemberProperty{JsiiProperty: "policyEnforcementInput", GoGetter: "PolicyEnforcementInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAllowedDatabricksDestinations", GoMethod: "PutAllowedDatabricksDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "putAllowedInternetDestinations", GoMethod: "PutAllowedInternetDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "putAllowedStorageDestinations", GoMethod: "PutAllowedStorageDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "putBlockedInternetDestinations", GoMethod: "PutBlockedInternetDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "putPolicyEnforcement", GoMethod: "PutPolicyEnforcement"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedDatabricksDestinations", GoMethod: "ResetAllowedDatabricksDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedInternetDestinations", GoMethod: "ResetAllowedInternetDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedStorageDestinations", GoMethod: "ResetAllowedStorageDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetBlockedInternetDestinations", GoMethod: "ResetBlockedInternetDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicyEnforcement", GoMethod: "ResetPolicyEnforcement"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionMode", GoGetter: "RestrictionMode"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionModeInput", GoGetter: "RestrictionModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessPolicyEnforcement",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessPolicyEnforcement)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessPolicyEnforcementOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessPolicyEnforcementOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dryRunModeProductFilter", GoGetter: "DryRunModeProductFilter"},
			_jsii_.MemberProperty{JsiiProperty: "dryRunModeProductFilterInput", GoGetter: "DryRunModeProductFilterInput"},
			_jsii_.MemberProperty{JsiiProperty: "enforcementMode", GoGetter: "EnforcementMode"},
			_jsii_.MemberProperty{JsiiProperty: "enforcementModeInput", GoGetter: "EnforcementModeInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDryRunModeProductFilter", GoMethod: "ResetDryRunModeProductFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnforcementMode", GoMethod: "ResetEnforcementMode"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessPolicyEnforcementOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsEgressOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsEgressOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "networkAccess", GoGetter: "NetworkAccess"},
			_jsii_.MemberProperty{JsiiProperty: "networkAccessInput", GoGetter: "NetworkAccessInput"},
			_jsii_.MemberMethod{JsiiMethod: "putNetworkAccess", GoMethod: "PutNetworkAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkAccess", GoMethod: "ResetNetworkAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsEgressOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngress",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccess",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccess)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRules",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesAuthentication",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesAuthenticationIdentities",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesAuthenticationIdentities)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesAuthenticationIdentitiesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesAuthenticationIdentitiesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesAuthenticationIdentitiesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesAuthenticationIdentitiesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesAuthenticationIdentitiesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberProperty{JsiiProperty: "principalIdInput", GoGetter: "PrincipalIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "principalType", GoGetter: "PrincipalType"},
			_jsii_.MemberProperty{JsiiProperty: "principalTypeInput", GoGetter: "PrincipalTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalId", GoMethod: "ResetPrincipalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalType", GoMethod: "ResetPrincipalType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesAuthenticationIdentitiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesAuthenticationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesAuthenticationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "identities", GoGetter: "Identities"},
			_jsii_.MemberProperty{JsiiProperty: "identitiesInput", GoGetter: "IdentitiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "identityType", GoGetter: "IdentityType"},
			_jsii_.MemberProperty{JsiiProperty: "identityTypeInput", GoGetter: "IdentityTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIdentities", GoMethod: "PutIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentities", GoMethod: "ResetIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityType", GoMethod: "ResetIdentityType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestination",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAccountApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAccountApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAccountApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAccountApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAccountApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAccountDatabricksOne",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAccountDatabricksOne)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAccountDatabricksOneOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAccountDatabricksOneOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAccountDatabricksOneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAccountUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAccountUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAccountUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAccountUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAccountUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAppsRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAppsRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAppsRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAppsRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationAppsRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationLakebaseRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationLakebaseRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationLakebaseRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationLakebaseRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationLakebaseRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountApi", GoGetter: "AccountApi"},
			_jsii_.MemberProperty{JsiiProperty: "accountApiInput", GoGetter: "AccountApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOne", GoGetter: "AccountDatabricksOne"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOneInput", GoGetter: "AccountDatabricksOneInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountUi", GoGetter: "AccountUi"},
			_jsii_.MemberProperty{JsiiProperty: "accountUiInput", GoGetter: "AccountUiInput"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntime", GoGetter: "AppsRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntimeInput", GoGetter: "AppsRuntimeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntime", GoGetter: "LakebaseRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntimeInput", GoGetter: "LakebaseRuntimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountApi", GoMethod: "PutAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountDatabricksOne", GoMethod: "PutAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountUi", GoMethod: "PutAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "putAppsRuntime", GoMethod: "PutAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putLakebaseRuntime", GoMethod: "PutLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceApi", GoMethod: "PutWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceUi", GoMethod: "PutWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountApi", GoMethod: "ResetAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountDatabricksOne", GoMethod: "ResetAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountUi", GoMethod: "ResetAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppsRuntime", GoMethod: "ResetAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetLakebaseRuntime", GoMethod: "ResetLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceApi", GoMethod: "ResetWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceUi", GoMethod: "ResetWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApi", GoGetter: "WorkspaceApi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApiInput", GoGetter: "WorkspaceApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUi", GoGetter: "WorkspaceUi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUiInput", GoGetter: "WorkspaceUiInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationWorkspaceApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationWorkspaceApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationWorkspaceApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationWorkspaceApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationWorkspaceApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationWorkspaceUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationWorkspaceUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationWorkspaceUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationWorkspaceUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesDestinationWorkspaceUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesOrigin",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesOrigin)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesOriginOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesOriginOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allSourceWorkspaces", GoGetter: "AllSourceWorkspaces"},
			_jsii_.MemberProperty{JsiiProperty: "allSourceWorkspacesInput", GoGetter: "AllSourceWorkspacesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putSelectedWorkspaces", GoMethod: "PutSelectedWorkspaces"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllSourceWorkspaces", GoMethod: "ResetAllSourceWorkspaces"},
			_jsii_.MemberMethod{JsiiMethod: "resetSelectedWorkspaces", GoMethod: "ResetSelectedWorkspaces"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "selectedWorkspaces", GoGetter: "SelectedWorkspaces"},
			_jsii_.MemberProperty{JsiiProperty: "selectedWorkspacesInput", GoGetter: "SelectedWorkspacesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesOriginOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesOriginSelectedWorkspaces",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesOriginSelectedWorkspaces)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesOriginSelectedWorkspacesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesOriginSelectedWorkspacesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceIds", GoMethod: "ResetWorkspaceIds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceIds", GoGetter: "WorkspaceIds"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceIdsInput", GoGetter: "WorkspaceIdsInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesOriginSelectedWorkspacesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationInput", GoGetter: "AuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberProperty{JsiiProperty: "originInput", GoGetter: "OriginInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthentication", GoMethod: "PutAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "putDestination", GoMethod: "PutDestination"},
			_jsii_.MemberMethod{JsiiMethod: "putOrigin", GoMethod: "PutOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthentication", GoMethod: "ResetAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabel", GoMethod: "ResetLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrigin", GoMethod: "ResetOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessAllowRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRules",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesAuthentication",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesAuthenticationIdentities",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesAuthenticationIdentities)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesAuthenticationIdentitiesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesAuthenticationIdentitiesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesAuthenticationIdentitiesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesAuthenticationIdentitiesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesAuthenticationIdentitiesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberProperty{JsiiProperty: "principalIdInput", GoGetter: "PrincipalIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "principalType", GoGetter: "PrincipalType"},
			_jsii_.MemberProperty{JsiiProperty: "principalTypeInput", GoGetter: "PrincipalTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalId", GoMethod: "ResetPrincipalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalType", GoMethod: "ResetPrincipalType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesAuthenticationIdentitiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesAuthenticationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesAuthenticationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "identities", GoGetter: "Identities"},
			_jsii_.MemberProperty{JsiiProperty: "identitiesInput", GoGetter: "IdentitiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "identityType", GoGetter: "IdentityType"},
			_jsii_.MemberProperty{JsiiProperty: "identityTypeInput", GoGetter: "IdentityTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIdentities", GoMethod: "PutIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentities", GoMethod: "ResetIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityType", GoMethod: "ResetIdentityType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestination",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAccountApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAccountApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAccountApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAccountApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAccountApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAccountDatabricksOne",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAccountDatabricksOne)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAccountDatabricksOneOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAccountDatabricksOneOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAccountDatabricksOneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAccountUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAccountUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAccountUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAccountUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAccountUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAppsRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAppsRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAppsRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAppsRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationAppsRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationLakebaseRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationLakebaseRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationLakebaseRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationLakebaseRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationLakebaseRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountApi", GoGetter: "AccountApi"},
			_jsii_.MemberProperty{JsiiProperty: "accountApiInput", GoGetter: "AccountApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOne", GoGetter: "AccountDatabricksOne"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOneInput", GoGetter: "AccountDatabricksOneInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountUi", GoGetter: "AccountUi"},
			_jsii_.MemberProperty{JsiiProperty: "accountUiInput", GoGetter: "AccountUiInput"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntime", GoGetter: "AppsRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntimeInput", GoGetter: "AppsRuntimeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntime", GoGetter: "LakebaseRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntimeInput", GoGetter: "LakebaseRuntimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountApi", GoMethod: "PutAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountDatabricksOne", GoMethod: "PutAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountUi", GoMethod: "PutAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "putAppsRuntime", GoMethod: "PutAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putLakebaseRuntime", GoMethod: "PutLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceApi", GoMethod: "PutWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceUi", GoMethod: "PutWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountApi", GoMethod: "ResetAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountDatabricksOne", GoMethod: "ResetAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountUi", GoMethod: "ResetAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppsRuntime", GoMethod: "ResetAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetLakebaseRuntime", GoMethod: "ResetLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceApi", GoMethod: "ResetWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceUi", GoMethod: "ResetWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApi", GoGetter: "WorkspaceApi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApiInput", GoGetter: "WorkspaceApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUi", GoGetter: "WorkspaceUi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUiInput", GoGetter: "WorkspaceUiInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesDestinationWorkspaceUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesOrigin",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesOrigin)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesOriginOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesOriginOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allSourceWorkspaces", GoGetter: "AllSourceWorkspaces"},
			_jsii_.MemberProperty{JsiiProperty: "allSourceWorkspacesInput", GoGetter: "AllSourceWorkspacesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putSelectedWorkspaces", GoMethod: "PutSelectedWorkspaces"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllSourceWorkspaces", GoMethod: "ResetAllSourceWorkspaces"},
			_jsii_.MemberMethod{JsiiMethod: "resetSelectedWorkspaces", GoMethod: "ResetSelectedWorkspaces"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "selectedWorkspaces", GoGetter: "SelectedWorkspaces"},
			_jsii_.MemberProperty{JsiiProperty: "selectedWorkspacesInput", GoGetter: "SelectedWorkspacesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesOriginOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesOriginSelectedWorkspaces",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesOriginSelectedWorkspaces)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesOriginSelectedWorkspacesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesOriginSelectedWorkspacesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceIds", GoMethod: "ResetWorkspaceIds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceIds", GoGetter: "WorkspaceIds"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceIdsInput", GoGetter: "WorkspaceIdsInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesOriginSelectedWorkspacesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationInput", GoGetter: "AuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberProperty{JsiiProperty: "originInput", GoGetter: "OriginInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthentication", GoMethod: "PutAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "putDestination", GoMethod: "PutDestination"},
			_jsii_.MemberMethod{JsiiMethod: "putOrigin", GoMethod: "PutOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthentication", GoMethod: "ResetAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabel", GoMethod: "ResetLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrigin", GoMethod: "ResetOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessDenyRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowRules", GoGetter: "AllowRules"},
			_jsii_.MemberProperty{JsiiProperty: "allowRulesInput", GoGetter: "AllowRulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "denyRules", GoGetter: "DenyRules"},
			_jsii_.MemberProperty{JsiiProperty: "denyRulesInput", GoGetter: "DenyRulesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putAllowRules", GoMethod: "PutAllowRules"},
			_jsii_.MemberMethod{JsiiMethod: "putDenyRules", GoMethod: "PutDenyRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowRules", GoMethod: "ResetAllowRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetDenyRules", GoMethod: "ResetDenyRules"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionMode", GoGetter: "RestrictionMode"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionModeInput", GoGetter: "RestrictionModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressCrossWorkspaceAccessOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRun",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRun)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccess",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccess)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRules",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesAuthentication",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesAuthenticationIdentities",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesAuthenticationIdentities)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesAuthenticationIdentitiesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesAuthenticationIdentitiesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesAuthenticationIdentitiesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesAuthenticationIdentitiesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesAuthenticationIdentitiesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberProperty{JsiiProperty: "principalIdInput", GoGetter: "PrincipalIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "principalType", GoGetter: "PrincipalType"},
			_jsii_.MemberProperty{JsiiProperty: "principalTypeInput", GoGetter: "PrincipalTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalId", GoMethod: "ResetPrincipalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalType", GoMethod: "ResetPrincipalType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesAuthenticationIdentitiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesAuthenticationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesAuthenticationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "identities", GoGetter: "Identities"},
			_jsii_.MemberProperty{JsiiProperty: "identitiesInput", GoGetter: "IdentitiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "identityType", GoGetter: "IdentityType"},
			_jsii_.MemberProperty{JsiiProperty: "identityTypeInput", GoGetter: "IdentityTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIdentities", GoMethod: "PutIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentities", GoMethod: "ResetIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityType", GoMethod: "ResetIdentityType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestination",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountDatabricksOne",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountDatabricksOne)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountDatabricksOneOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountDatabricksOneOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountDatabricksOneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAccountUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAppsRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAppsRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAppsRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAppsRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationAppsRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationLakebaseRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationLakebaseRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationLakebaseRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationLakebaseRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationLakebaseRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountApi", GoGetter: "AccountApi"},
			_jsii_.MemberProperty{JsiiProperty: "accountApiInput", GoGetter: "AccountApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOne", GoGetter: "AccountDatabricksOne"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOneInput", GoGetter: "AccountDatabricksOneInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountUi", GoGetter: "AccountUi"},
			_jsii_.MemberProperty{JsiiProperty: "accountUiInput", GoGetter: "AccountUiInput"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntime", GoGetter: "AppsRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntimeInput", GoGetter: "AppsRuntimeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntime", GoGetter: "LakebaseRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntimeInput", GoGetter: "LakebaseRuntimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountApi", GoMethod: "PutAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountDatabricksOne", GoMethod: "PutAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountUi", GoMethod: "PutAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "putAppsRuntime", GoMethod: "PutAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putLakebaseRuntime", GoMethod: "PutLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceApi", GoMethod: "PutWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceUi", GoMethod: "PutWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountApi", GoMethod: "ResetAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountDatabricksOne", GoMethod: "ResetAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountUi", GoMethod: "ResetAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppsRuntime", GoMethod: "ResetAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetLakebaseRuntime", GoMethod: "ResetLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceApi", GoMethod: "ResetWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceUi", GoMethod: "ResetWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApi", GoGetter: "WorkspaceApi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApiInput", GoGetter: "WorkspaceApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUi", GoGetter: "WorkspaceUi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUiInput", GoGetter: "WorkspaceUiInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesDestinationWorkspaceUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesOrigin",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesOrigin)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesOriginOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesOriginOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allSourceWorkspaces", GoGetter: "AllSourceWorkspaces"},
			_jsii_.MemberProperty{JsiiProperty: "allSourceWorkspacesInput", GoGetter: "AllSourceWorkspacesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putSelectedWorkspaces", GoMethod: "PutSelectedWorkspaces"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllSourceWorkspaces", GoMethod: "ResetAllSourceWorkspaces"},
			_jsii_.MemberMethod{JsiiMethod: "resetSelectedWorkspaces", GoMethod: "ResetSelectedWorkspaces"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "selectedWorkspaces", GoGetter: "SelectedWorkspaces"},
			_jsii_.MemberProperty{JsiiProperty: "selectedWorkspacesInput", GoGetter: "SelectedWorkspacesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesOriginOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesOriginSelectedWorkspaces",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesOriginSelectedWorkspaces)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesOriginSelectedWorkspacesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesOriginSelectedWorkspacesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceIds", GoMethod: "ResetWorkspaceIds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceIds", GoGetter: "WorkspaceIds"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceIdsInput", GoGetter: "WorkspaceIdsInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesOriginSelectedWorkspacesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationInput", GoGetter: "AuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberProperty{JsiiProperty: "originInput", GoGetter: "OriginInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthentication", GoMethod: "PutAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "putDestination", GoMethod: "PutDestination"},
			_jsii_.MemberMethod{JsiiMethod: "putOrigin", GoMethod: "PutOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthentication", GoMethod: "ResetAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabel", GoMethod: "ResetLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrigin", GoMethod: "ResetOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessAllowRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRules",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesAuthentication",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationIdentities",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationIdentities)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationIdentitiesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationIdentitiesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationIdentitiesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationIdentitiesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationIdentitiesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberProperty{JsiiProperty: "principalIdInput", GoGetter: "PrincipalIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "principalType", GoGetter: "PrincipalType"},
			_jsii_.MemberProperty{JsiiProperty: "principalTypeInput", GoGetter: "PrincipalTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalId", GoMethod: "ResetPrincipalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalType", GoMethod: "ResetPrincipalType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationIdentitiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "identities", GoGetter: "Identities"},
			_jsii_.MemberProperty{JsiiProperty: "identitiesInput", GoGetter: "IdentitiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "identityType", GoGetter: "IdentityType"},
			_jsii_.MemberProperty{JsiiProperty: "identityTypeInput", GoGetter: "IdentityTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIdentities", GoMethod: "PutIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentities", GoMethod: "ResetIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityType", GoMethod: "ResetIdentityType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestination",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAccountApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAccountApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAccountApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAccountApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAccountApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAccountDatabricksOne",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAccountDatabricksOne)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAccountDatabricksOneOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAccountDatabricksOneOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAccountDatabricksOneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAccountUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAccountUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAccountUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAccountUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAccountUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAppsRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAppsRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAppsRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAppsRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationAppsRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationLakebaseRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationLakebaseRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationLakebaseRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationLakebaseRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationLakebaseRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountApi", GoGetter: "AccountApi"},
			_jsii_.MemberProperty{JsiiProperty: "accountApiInput", GoGetter: "AccountApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOne", GoGetter: "AccountDatabricksOne"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOneInput", GoGetter: "AccountDatabricksOneInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountUi", GoGetter: "AccountUi"},
			_jsii_.MemberProperty{JsiiProperty: "accountUiInput", GoGetter: "AccountUiInput"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntime", GoGetter: "AppsRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntimeInput", GoGetter: "AppsRuntimeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntime", GoGetter: "LakebaseRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntimeInput", GoGetter: "LakebaseRuntimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountApi", GoMethod: "PutAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountDatabricksOne", GoMethod: "PutAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountUi", GoMethod: "PutAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "putAppsRuntime", GoMethod: "PutAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putLakebaseRuntime", GoMethod: "PutLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceApi", GoMethod: "PutWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceUi", GoMethod: "PutWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountApi", GoMethod: "ResetAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountDatabricksOne", GoMethod: "ResetAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountUi", GoMethod: "ResetAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppsRuntime", GoMethod: "ResetAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetLakebaseRuntime", GoMethod: "ResetLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceApi", GoMethod: "ResetWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceUi", GoMethod: "ResetWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApi", GoGetter: "WorkspaceApi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApiInput", GoGetter: "WorkspaceApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUi", GoGetter: "WorkspaceUi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUiInput", GoGetter: "WorkspaceUiInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationWorkspaceApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationWorkspaceApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationWorkspaceApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationWorkspaceApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationWorkspaceApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationWorkspaceUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationWorkspaceUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationWorkspaceUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationWorkspaceUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesDestinationWorkspaceUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesOrigin",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesOrigin)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesOriginOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesOriginOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allSourceWorkspaces", GoGetter: "AllSourceWorkspaces"},
			_jsii_.MemberProperty{JsiiProperty: "allSourceWorkspacesInput", GoGetter: "AllSourceWorkspacesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putSelectedWorkspaces", GoMethod: "PutSelectedWorkspaces"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllSourceWorkspaces", GoMethod: "ResetAllSourceWorkspaces"},
			_jsii_.MemberMethod{JsiiMethod: "resetSelectedWorkspaces", GoMethod: "ResetSelectedWorkspaces"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "selectedWorkspaces", GoGetter: "SelectedWorkspaces"},
			_jsii_.MemberProperty{JsiiProperty: "selectedWorkspacesInput", GoGetter: "SelectedWorkspacesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesOriginOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesOriginSelectedWorkspaces",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesOriginSelectedWorkspaces)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesOriginSelectedWorkspacesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesOriginSelectedWorkspacesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceIds", GoMethod: "ResetWorkspaceIds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceIds", GoGetter: "WorkspaceIds"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceIdsInput", GoGetter: "WorkspaceIdsInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesOriginSelectedWorkspacesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationInput", GoGetter: "AuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberProperty{JsiiProperty: "originInput", GoGetter: "OriginInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthentication", GoMethod: "PutAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "putDestination", GoMethod: "PutDestination"},
			_jsii_.MemberMethod{JsiiMethod: "putOrigin", GoMethod: "PutOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthentication", GoMethod: "ResetAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabel", GoMethod: "ResetLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrigin", GoMethod: "ResetOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessDenyRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowRules", GoGetter: "AllowRules"},
			_jsii_.MemberProperty{JsiiProperty: "allowRulesInput", GoGetter: "AllowRulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "denyRules", GoGetter: "DenyRules"},
			_jsii_.MemberProperty{JsiiProperty: "denyRulesInput", GoGetter: "DenyRulesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putAllowRules", GoMethod: "PutAllowRules"},
			_jsii_.MemberMethod{JsiiMethod: "putDenyRules", GoMethod: "PutDenyRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowRules", GoMethod: "ResetAllowRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetDenyRules", GoMethod: "ResetDenyRules"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionMode", GoGetter: "RestrictionMode"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionModeInput", GoGetter: "RestrictionModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunCrossWorkspaceAccessOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "crossWorkspaceAccess", GoGetter: "CrossWorkspaceAccess"},
			_jsii_.MemberProperty{JsiiProperty: "crossWorkspaceAccessInput", GoGetter: "CrossWorkspaceAccessInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "privateAccess", GoGetter: "PrivateAccess"},
			_jsii_.MemberProperty{JsiiProperty: "privateAccessInput", GoGetter: "PrivateAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "publicAccess", GoGetter: "PublicAccess"},
			_jsii_.MemberProperty{JsiiProperty: "publicAccessInput", GoGetter: "PublicAccessInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCrossWorkspaceAccess", GoMethod: "PutCrossWorkspaceAccess"},
			_jsii_.MemberMethod{JsiiMethod: "putPrivateAccess", GoMethod: "PutPrivateAccess"},
			_jsii_.MemberMethod{JsiiMethod: "putPublicAccess", GoMethod: "PutPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetCrossWorkspaceAccess", GoMethod: "ResetCrossWorkspaceAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateAccess", GoMethod: "ResetPrivateAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicAccess", GoMethod: "ResetPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccess",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccess)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRules",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesAuthentication",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesAuthenticationIdentities",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesAuthenticationIdentities)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesAuthenticationIdentitiesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesAuthenticationIdentitiesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesAuthenticationIdentitiesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesAuthenticationIdentitiesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesAuthenticationIdentitiesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberProperty{JsiiProperty: "principalIdInput", GoGetter: "PrincipalIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "principalType", GoGetter: "PrincipalType"},
			_jsii_.MemberProperty{JsiiProperty: "principalTypeInput", GoGetter: "PrincipalTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalId", GoMethod: "ResetPrincipalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalType", GoMethod: "ResetPrincipalType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesAuthenticationIdentitiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesAuthenticationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesAuthenticationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "identities", GoGetter: "Identities"},
			_jsii_.MemberProperty{JsiiProperty: "identitiesInput", GoGetter: "IdentitiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "identityType", GoGetter: "IdentityType"},
			_jsii_.MemberProperty{JsiiProperty: "identityTypeInput", GoGetter: "IdentityTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIdentities", GoMethod: "PutIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentities", GoMethod: "ResetIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityType", GoMethod: "ResetIdentityType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestination",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountDatabricksOne",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountDatabricksOne)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountDatabricksOneOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountDatabricksOneOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountDatabricksOneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAccountUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAppsRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAppsRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAppsRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAppsRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationAppsRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationLakebaseRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationLakebaseRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationLakebaseRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationLakebaseRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationLakebaseRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountApi", GoGetter: "AccountApi"},
			_jsii_.MemberProperty{JsiiProperty: "accountApiInput", GoGetter: "AccountApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOne", GoGetter: "AccountDatabricksOne"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOneInput", GoGetter: "AccountDatabricksOneInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountUi", GoGetter: "AccountUi"},
			_jsii_.MemberProperty{JsiiProperty: "accountUiInput", GoGetter: "AccountUiInput"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntime", GoGetter: "AppsRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntimeInput", GoGetter: "AppsRuntimeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntime", GoGetter: "LakebaseRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntimeInput", GoGetter: "LakebaseRuntimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountApi", GoMethod: "PutAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountDatabricksOne", GoMethod: "PutAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountUi", GoMethod: "PutAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "putAppsRuntime", GoMethod: "PutAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putLakebaseRuntime", GoMethod: "PutLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceApi", GoMethod: "PutWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceUi", GoMethod: "PutWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountApi", GoMethod: "ResetAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountDatabricksOne", GoMethod: "ResetAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountUi", GoMethod: "ResetAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppsRuntime", GoMethod: "ResetAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetLakebaseRuntime", GoMethod: "ResetLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceApi", GoMethod: "ResetWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceUi", GoMethod: "ResetWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApi", GoGetter: "WorkspaceApi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApiInput", GoGetter: "WorkspaceApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUi", GoGetter: "WorkspaceUi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUiInput", GoGetter: "WorkspaceUiInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesDestinationWorkspaceUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesOrigin",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesOrigin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesOriginEndpoints",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesOriginEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesOriginEndpointsOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesOriginEndpointsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endpointIds", GoGetter: "EndpointIds"},
			_jsii_.MemberProperty{JsiiProperty: "endpointIdsInput", GoGetter: "EndpointIdsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetEndpointIds", GoMethod: "ResetEndpointIds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesOriginEndpointsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesOriginOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesOriginOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allPrivateAccess", GoGetter: "AllPrivateAccess"},
			_jsii_.MemberProperty{JsiiProperty: "allPrivateAccessInput", GoGetter: "AllPrivateAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "allRegisteredEndpoints", GoGetter: "AllRegisteredEndpoints"},
			_jsii_.MemberProperty{JsiiProperty: "allRegisteredEndpointsInput", GoGetter: "AllRegisteredEndpointsInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureWorkspacePrivateLink", GoGetter: "AzureWorkspacePrivateLink"},
			_jsii_.MemberProperty{JsiiProperty: "azureWorkspacePrivateLinkInput", GoGetter: "AzureWorkspacePrivateLinkInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
			_jsii_.MemberProperty{JsiiProperty: "endpointsInput", GoGetter: "EndpointsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putEndpoints", GoMethod: "PutEndpoints"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllPrivateAccess", GoMethod: "ResetAllPrivateAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllRegisteredEndpoints", GoMethod: "ResetAllRegisteredEndpoints"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureWorkspacePrivateLink", GoMethod: "ResetAzureWorkspacePrivateLink"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndpoints", GoMethod: "ResetEndpoints"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesOriginOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationInput", GoGetter: "AuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberProperty{JsiiProperty: "originInput", GoGetter: "OriginInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthentication", GoMethod: "PutAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "putDestination", GoMethod: "PutDestination"},
			_jsii_.MemberMethod{JsiiMethod: "putOrigin", GoMethod: "PutOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthentication", GoMethod: "ResetAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabel", GoMethod: "ResetLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrigin", GoMethod: "ResetOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessAllowRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRules",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesAuthentication",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesAuthenticationIdentities",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesAuthenticationIdentities)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberProperty{JsiiProperty: "principalIdInput", GoGetter: "PrincipalIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "principalType", GoGetter: "PrincipalType"},
			_jsii_.MemberProperty{JsiiProperty: "principalTypeInput", GoGetter: "PrincipalTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalId", GoMethod: "ResetPrincipalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalType", GoMethod: "ResetPrincipalType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesAuthenticationIdentitiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesAuthenticationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesAuthenticationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "identities", GoGetter: "Identities"},
			_jsii_.MemberProperty{JsiiProperty: "identitiesInput", GoGetter: "IdentitiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "identityType", GoGetter: "IdentityType"},
			_jsii_.MemberProperty{JsiiProperty: "identityTypeInput", GoGetter: "IdentityTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIdentities", GoMethod: "PutIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentities", GoMethod: "ResetIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityType", GoMethod: "ResetIdentityType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestination",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAccountApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAccountApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAccountApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAccountApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAccountApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAccountDatabricksOne",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAccountDatabricksOne)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAccountDatabricksOneOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAccountDatabricksOneOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAccountDatabricksOneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAccountUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAccountUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAccountUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAccountUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAccountUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAppsRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAppsRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAppsRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAppsRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationAppsRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationLakebaseRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationLakebaseRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationLakebaseRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationLakebaseRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationLakebaseRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountApi", GoGetter: "AccountApi"},
			_jsii_.MemberProperty{JsiiProperty: "accountApiInput", GoGetter: "AccountApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOne", GoGetter: "AccountDatabricksOne"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOneInput", GoGetter: "AccountDatabricksOneInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountUi", GoGetter: "AccountUi"},
			_jsii_.MemberProperty{JsiiProperty: "accountUiInput", GoGetter: "AccountUiInput"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntime", GoGetter: "AppsRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntimeInput", GoGetter: "AppsRuntimeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntime", GoGetter: "LakebaseRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntimeInput", GoGetter: "LakebaseRuntimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountApi", GoMethod: "PutAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountDatabricksOne", GoMethod: "PutAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountUi", GoMethod: "PutAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "putAppsRuntime", GoMethod: "PutAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putLakebaseRuntime", GoMethod: "PutLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceApi", GoMethod: "PutWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceUi", GoMethod: "PutWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountApi", GoMethod: "ResetAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountDatabricksOne", GoMethod: "ResetAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountUi", GoMethod: "ResetAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppsRuntime", GoMethod: "ResetAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetLakebaseRuntime", GoMethod: "ResetLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceApi", GoMethod: "ResetWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceUi", GoMethod: "ResetWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApi", GoGetter: "WorkspaceApi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApiInput", GoGetter: "WorkspaceApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUi", GoGetter: "WorkspaceUi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUiInput", GoGetter: "WorkspaceUiInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationWorkspaceApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationWorkspaceApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationWorkspaceApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationWorkspaceApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationWorkspaceApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationWorkspaceUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationWorkspaceUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationWorkspaceUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationWorkspaceUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesDestinationWorkspaceUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesOrigin",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesOrigin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesOriginEndpoints",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesOriginEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesOriginEndpointsOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesOriginEndpointsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endpointIds", GoGetter: "EndpointIds"},
			_jsii_.MemberProperty{JsiiProperty: "endpointIdsInput", GoGetter: "EndpointIdsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetEndpointIds", GoMethod: "ResetEndpointIds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesOriginEndpointsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesOriginOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesOriginOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allPrivateAccess", GoGetter: "AllPrivateAccess"},
			_jsii_.MemberProperty{JsiiProperty: "allPrivateAccessInput", GoGetter: "AllPrivateAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "allRegisteredEndpoints", GoGetter: "AllRegisteredEndpoints"},
			_jsii_.MemberProperty{JsiiProperty: "allRegisteredEndpointsInput", GoGetter: "AllRegisteredEndpointsInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureWorkspacePrivateLink", GoGetter: "AzureWorkspacePrivateLink"},
			_jsii_.MemberProperty{JsiiProperty: "azureWorkspacePrivateLinkInput", GoGetter: "AzureWorkspacePrivateLinkInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
			_jsii_.MemberProperty{JsiiProperty: "endpointsInput", GoGetter: "EndpointsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putEndpoints", GoMethod: "PutEndpoints"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllPrivateAccess", GoMethod: "ResetAllPrivateAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllRegisteredEndpoints", GoMethod: "ResetAllRegisteredEndpoints"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureWorkspacePrivateLink", GoMethod: "ResetAzureWorkspacePrivateLink"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndpoints", GoMethod: "ResetEndpoints"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesOriginOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationInput", GoGetter: "AuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberProperty{JsiiProperty: "originInput", GoGetter: "OriginInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthentication", GoMethod: "PutAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "putDestination", GoMethod: "PutDestination"},
			_jsii_.MemberMethod{JsiiMethod: "putOrigin", GoMethod: "PutOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthentication", GoMethod: "ResetAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabel", GoMethod: "ResetLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrigin", GoMethod: "ResetOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessDenyRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowRules", GoGetter: "AllowRules"},
			_jsii_.MemberProperty{JsiiProperty: "allowRulesInput", GoGetter: "AllowRulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "denyRules", GoGetter: "DenyRules"},
			_jsii_.MemberProperty{JsiiProperty: "denyRulesInput", GoGetter: "DenyRulesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putAllowRules", GoMethod: "PutAllowRules"},
			_jsii_.MemberMethod{JsiiMethod: "putDenyRules", GoMethod: "PutDenyRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowRules", GoMethod: "ResetAllowRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetDenyRules", GoMethod: "ResetDenyRules"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionMode", GoGetter: "RestrictionMode"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionModeInput", GoGetter: "RestrictionModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPrivateAccessOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccess",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccess)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRules",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesAuthentication",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesAuthenticationIdentities",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesAuthenticationIdentities)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesAuthenticationIdentitiesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesAuthenticationIdentitiesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesAuthenticationIdentitiesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesAuthenticationIdentitiesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesAuthenticationIdentitiesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberProperty{JsiiProperty: "principalIdInput", GoGetter: "PrincipalIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "principalType", GoGetter: "PrincipalType"},
			_jsii_.MemberProperty{JsiiProperty: "principalTypeInput", GoGetter: "PrincipalTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalId", GoMethod: "ResetPrincipalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalType", GoMethod: "ResetPrincipalType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesAuthenticationIdentitiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesAuthenticationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesAuthenticationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "identities", GoGetter: "Identities"},
			_jsii_.MemberProperty{JsiiProperty: "identitiesInput", GoGetter: "IdentitiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "identityType", GoGetter: "IdentityType"},
			_jsii_.MemberProperty{JsiiProperty: "identityTypeInput", GoGetter: "IdentityTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIdentities", GoMethod: "PutIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentities", GoMethod: "ResetIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityType", GoMethod: "ResetIdentityType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestination",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAccountApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAccountApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAccountApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAccountApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAccountApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAccountDatabricksOne",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAccountDatabricksOne)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAccountDatabricksOneOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAccountDatabricksOneOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAccountDatabricksOneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAccountUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAccountUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAccountUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAccountUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAccountUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAppsRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAppsRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAppsRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAppsRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationAppsRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationLakebaseRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationLakebaseRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationLakebaseRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationLakebaseRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationLakebaseRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountApi", GoGetter: "AccountApi"},
			_jsii_.MemberProperty{JsiiProperty: "accountApiInput", GoGetter: "AccountApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOne", GoGetter: "AccountDatabricksOne"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOneInput", GoGetter: "AccountDatabricksOneInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountUi", GoGetter: "AccountUi"},
			_jsii_.MemberProperty{JsiiProperty: "accountUiInput", GoGetter: "AccountUiInput"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntime", GoGetter: "AppsRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntimeInput", GoGetter: "AppsRuntimeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntime", GoGetter: "LakebaseRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntimeInput", GoGetter: "LakebaseRuntimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountApi", GoMethod: "PutAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountDatabricksOne", GoMethod: "PutAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountUi", GoMethod: "PutAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "putAppsRuntime", GoMethod: "PutAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putLakebaseRuntime", GoMethod: "PutLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceApi", GoMethod: "PutWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceUi", GoMethod: "PutWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountApi", GoMethod: "ResetAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountDatabricksOne", GoMethod: "ResetAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountUi", GoMethod: "ResetAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppsRuntime", GoMethod: "ResetAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetLakebaseRuntime", GoMethod: "ResetLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceApi", GoMethod: "ResetWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceUi", GoMethod: "ResetWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApi", GoGetter: "WorkspaceApi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApiInput", GoGetter: "WorkspaceApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUi", GoGetter: "WorkspaceUi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUiInput", GoGetter: "WorkspaceUiInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationWorkspaceApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationWorkspaceApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationWorkspaceApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationWorkspaceApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationWorkspaceApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationWorkspaceUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationWorkspaceUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationWorkspaceUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationWorkspaceUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesDestinationWorkspaceUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOrigin",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOrigin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOriginExcludedIpRanges",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOriginExcludedIpRanges)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOriginExcludedIpRangesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOriginExcludedIpRangesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ipRanges", GoGetter: "IpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "ipRangesInput", GoGetter: "IpRangesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpRanges", GoMethod: "ResetIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOriginExcludedIpRangesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOriginIncludedIpRanges",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOriginIncludedIpRanges)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOriginIncludedIpRangesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOriginIncludedIpRangesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ipRanges", GoGetter: "IpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "ipRangesInput", GoGetter: "IpRangesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpRanges", GoMethod: "ResetIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOriginIncludedIpRangesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOriginOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOriginOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allIpRanges", GoGetter: "AllIpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "allIpRangesInput", GoGetter: "AllIpRangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludedIpRanges", GoGetter: "ExcludedIpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "excludedIpRangesInput", GoGetter: "ExcludedIpRangesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includedIpRanges", GoGetter: "IncludedIpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "includedIpRangesInput", GoGetter: "IncludedIpRangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putExcludedIpRanges", GoMethod: "PutExcludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "putIncludedIpRanges", GoMethod: "PutIncludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllIpRanges", GoMethod: "ResetAllIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludedIpRanges", GoMethod: "ResetExcludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludedIpRanges", GoMethod: "ResetIncludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOriginOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationInput", GoGetter: "AuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberProperty{JsiiProperty: "originInput", GoGetter: "OriginInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthentication", GoMethod: "PutAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "putDestination", GoMethod: "PutDestination"},
			_jsii_.MemberMethod{JsiiMethod: "putOrigin", GoMethod: "PutOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthentication", GoMethod: "ResetAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabel", GoMethod: "ResetLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrigin", GoMethod: "ResetOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessAllowRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRules",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesAuthentication",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesAuthenticationIdentities",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesAuthenticationIdentities)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesAuthenticationIdentitiesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesAuthenticationIdentitiesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesAuthenticationIdentitiesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesAuthenticationIdentitiesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesAuthenticationIdentitiesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberProperty{JsiiProperty: "principalIdInput", GoGetter: "PrincipalIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "principalType", GoGetter: "PrincipalType"},
			_jsii_.MemberProperty{JsiiProperty: "principalTypeInput", GoGetter: "PrincipalTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalId", GoMethod: "ResetPrincipalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalType", GoMethod: "ResetPrincipalType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesAuthenticationIdentitiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesAuthenticationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesAuthenticationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "identities", GoGetter: "Identities"},
			_jsii_.MemberProperty{JsiiProperty: "identitiesInput", GoGetter: "IdentitiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "identityType", GoGetter: "IdentityType"},
			_jsii_.MemberProperty{JsiiProperty: "identityTypeInput", GoGetter: "IdentityTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIdentities", GoMethod: "PutIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentities", GoMethod: "ResetIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityType", GoMethod: "ResetIdentityType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestination",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAccountApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAccountApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAccountApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAccountApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAccountApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAccountDatabricksOne",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAccountDatabricksOne)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAccountDatabricksOneOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAccountDatabricksOneOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAccountDatabricksOneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAccountUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAccountUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAccountUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAccountUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAccountUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAppsRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAppsRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAppsRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAppsRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationAppsRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationLakebaseRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationLakebaseRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationLakebaseRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationLakebaseRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationLakebaseRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountApi", GoGetter: "AccountApi"},
			_jsii_.MemberProperty{JsiiProperty: "accountApiInput", GoGetter: "AccountApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOne", GoGetter: "AccountDatabricksOne"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOneInput", GoGetter: "AccountDatabricksOneInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountUi", GoGetter: "AccountUi"},
			_jsii_.MemberProperty{JsiiProperty: "accountUiInput", GoGetter: "AccountUiInput"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntime", GoGetter: "AppsRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntimeInput", GoGetter: "AppsRuntimeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntime", GoGetter: "LakebaseRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntimeInput", GoGetter: "LakebaseRuntimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountApi", GoMethod: "PutAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountDatabricksOne", GoMethod: "PutAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountUi", GoMethod: "PutAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "putAppsRuntime", GoMethod: "PutAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putLakebaseRuntime", GoMethod: "PutLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceApi", GoMethod: "PutWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceUi", GoMethod: "PutWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountApi", GoMethod: "ResetAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountDatabricksOne", GoMethod: "ResetAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountUi", GoMethod: "ResetAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppsRuntime", GoMethod: "ResetAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetLakebaseRuntime", GoMethod: "ResetLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceApi", GoMethod: "ResetWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceUi", GoMethod: "ResetWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApi", GoGetter: "WorkspaceApi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApiInput", GoGetter: "WorkspaceApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUi", GoGetter: "WorkspaceUi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUiInput", GoGetter: "WorkspaceUiInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationWorkspaceApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationWorkspaceApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationWorkspaceApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationWorkspaceApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationWorkspaceApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationWorkspaceUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationWorkspaceUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationWorkspaceUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationWorkspaceUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesDestinationWorkspaceUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOrigin",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOrigin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginExcludedIpRanges",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginExcludedIpRanges)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginExcludedIpRangesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginExcludedIpRangesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ipRanges", GoGetter: "IpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "ipRangesInput", GoGetter: "IpRangesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpRanges", GoMethod: "ResetIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginExcludedIpRangesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRanges",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRanges)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ipRanges", GoGetter: "IpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "ipRangesInput", GoGetter: "IpRangesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpRanges", GoMethod: "ResetIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginIncludedIpRangesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allIpRanges", GoGetter: "AllIpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "allIpRangesInput", GoGetter: "AllIpRangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludedIpRanges", GoGetter: "ExcludedIpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "excludedIpRangesInput", GoGetter: "ExcludedIpRangesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includedIpRanges", GoGetter: "IncludedIpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "includedIpRangesInput", GoGetter: "IncludedIpRangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putExcludedIpRanges", GoMethod: "PutExcludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "putIncludedIpRanges", GoMethod: "PutIncludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllIpRanges", GoMethod: "ResetAllIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludedIpRanges", GoMethod: "ResetExcludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludedIpRanges", GoMethod: "ResetIncludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOriginOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationInput", GoGetter: "AuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberProperty{JsiiProperty: "originInput", GoGetter: "OriginInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthentication", GoMethod: "PutAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "putDestination", GoMethod: "PutDestination"},
			_jsii_.MemberMethod{JsiiMethod: "putOrigin", GoMethod: "PutOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthentication", GoMethod: "ResetAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabel", GoMethod: "ResetLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrigin", GoMethod: "ResetOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessDenyRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowRules", GoGetter: "AllowRules"},
			_jsii_.MemberProperty{JsiiProperty: "allowRulesInput", GoGetter: "AllowRulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "denyRules", GoGetter: "DenyRules"},
			_jsii_.MemberProperty{JsiiProperty: "denyRulesInput", GoGetter: "DenyRulesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putAllowRules", GoMethod: "PutAllowRules"},
			_jsii_.MemberMethod{JsiiMethod: "putDenyRules", GoMethod: "PutDenyRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowRules", GoMethod: "ResetAllowRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetDenyRules", GoMethod: "ResetDenyRules"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionMode", GoGetter: "RestrictionMode"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionModeInput", GoGetter: "RestrictionModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressDryRunPublicAccessOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "crossWorkspaceAccess", GoGetter: "CrossWorkspaceAccess"},
			_jsii_.MemberProperty{JsiiProperty: "crossWorkspaceAccessInput", GoGetter: "CrossWorkspaceAccessInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "privateAccess", GoGetter: "PrivateAccess"},
			_jsii_.MemberProperty{JsiiProperty: "privateAccessInput", GoGetter: "PrivateAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "publicAccess", GoGetter: "PublicAccess"},
			_jsii_.MemberProperty{JsiiProperty: "publicAccessInput", GoGetter: "PublicAccessInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCrossWorkspaceAccess", GoMethod: "PutCrossWorkspaceAccess"},
			_jsii_.MemberMethod{JsiiMethod: "putPrivateAccess", GoMethod: "PutPrivateAccess"},
			_jsii_.MemberMethod{JsiiMethod: "putPublicAccess", GoMethod: "PutPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetCrossWorkspaceAccess", GoMethod: "ResetCrossWorkspaceAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateAccess", GoMethod: "ResetPrivateAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicAccess", GoMethod: "ResetPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccess",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccess)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRules",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesAuthentication",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesAuthenticationIdentities",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesAuthenticationIdentities)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesAuthenticationIdentitiesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesAuthenticationIdentitiesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesAuthenticationIdentitiesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesAuthenticationIdentitiesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesAuthenticationIdentitiesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberProperty{JsiiProperty: "principalIdInput", GoGetter: "PrincipalIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "principalType", GoGetter: "PrincipalType"},
			_jsii_.MemberProperty{JsiiProperty: "principalTypeInput", GoGetter: "PrincipalTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalId", GoMethod: "ResetPrincipalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalType", GoMethod: "ResetPrincipalType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesAuthenticationIdentitiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesAuthenticationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesAuthenticationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "identities", GoGetter: "Identities"},
			_jsii_.MemberProperty{JsiiProperty: "identitiesInput", GoGetter: "IdentitiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "identityType", GoGetter: "IdentityType"},
			_jsii_.MemberProperty{JsiiProperty: "identityTypeInput", GoGetter: "IdentityTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIdentities", GoMethod: "PutIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentities", GoMethod: "ResetIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityType", GoMethod: "ResetIdentityType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestination",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAccountApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAccountApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAccountApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAccountApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAccountApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAccountDatabricksOne",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAccountDatabricksOne)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAccountDatabricksOneOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAccountDatabricksOneOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAccountDatabricksOneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAccountUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAccountUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAccountUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAccountUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAccountUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAppsRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAppsRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAppsRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAppsRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationAppsRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationLakebaseRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationLakebaseRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationLakebaseRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationLakebaseRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationLakebaseRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountApi", GoGetter: "AccountApi"},
			_jsii_.MemberProperty{JsiiProperty: "accountApiInput", GoGetter: "AccountApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOne", GoGetter: "AccountDatabricksOne"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOneInput", GoGetter: "AccountDatabricksOneInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountUi", GoGetter: "AccountUi"},
			_jsii_.MemberProperty{JsiiProperty: "accountUiInput", GoGetter: "AccountUiInput"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntime", GoGetter: "AppsRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntimeInput", GoGetter: "AppsRuntimeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntime", GoGetter: "LakebaseRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntimeInput", GoGetter: "LakebaseRuntimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountApi", GoMethod: "PutAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountDatabricksOne", GoMethod: "PutAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountUi", GoMethod: "PutAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "putAppsRuntime", GoMethod: "PutAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putLakebaseRuntime", GoMethod: "PutLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceApi", GoMethod: "PutWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceUi", GoMethod: "PutWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountApi", GoMethod: "ResetAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountDatabricksOne", GoMethod: "ResetAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountUi", GoMethod: "ResetAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppsRuntime", GoMethod: "ResetAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetLakebaseRuntime", GoMethod: "ResetLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceApi", GoMethod: "ResetWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceUi", GoMethod: "ResetWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApi", GoGetter: "WorkspaceApi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApiInput", GoGetter: "WorkspaceApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUi", GoGetter: "WorkspaceUi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUiInput", GoGetter: "WorkspaceUiInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationWorkspaceApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationWorkspaceApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationWorkspaceApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationWorkspaceApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationWorkspaceApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationWorkspaceUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationWorkspaceUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationWorkspaceUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationWorkspaceUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesDestinationWorkspaceUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesOrigin",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesOrigin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesOriginEndpoints",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesOriginEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesOriginEndpointsOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesOriginEndpointsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endpointIds", GoGetter: "EndpointIds"},
			_jsii_.MemberProperty{JsiiProperty: "endpointIdsInput", GoGetter: "EndpointIdsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetEndpointIds", GoMethod: "ResetEndpointIds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesOriginEndpointsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesOriginOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesOriginOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allPrivateAccess", GoGetter: "AllPrivateAccess"},
			_jsii_.MemberProperty{JsiiProperty: "allPrivateAccessInput", GoGetter: "AllPrivateAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "allRegisteredEndpoints", GoGetter: "AllRegisteredEndpoints"},
			_jsii_.MemberProperty{JsiiProperty: "allRegisteredEndpointsInput", GoGetter: "AllRegisteredEndpointsInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureWorkspacePrivateLink", GoGetter: "AzureWorkspacePrivateLink"},
			_jsii_.MemberProperty{JsiiProperty: "azureWorkspacePrivateLinkInput", GoGetter: "AzureWorkspacePrivateLinkInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
			_jsii_.MemberProperty{JsiiProperty: "endpointsInput", GoGetter: "EndpointsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putEndpoints", GoMethod: "PutEndpoints"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllPrivateAccess", GoMethod: "ResetAllPrivateAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllRegisteredEndpoints", GoMethod: "ResetAllRegisteredEndpoints"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureWorkspacePrivateLink", GoMethod: "ResetAzureWorkspacePrivateLink"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndpoints", GoMethod: "ResetEndpoints"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesOriginOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationInput", GoGetter: "AuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberProperty{JsiiProperty: "originInput", GoGetter: "OriginInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthentication", GoMethod: "PutAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "putDestination", GoMethod: "PutDestination"},
			_jsii_.MemberMethod{JsiiMethod: "putOrigin", GoMethod: "PutOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthentication", GoMethod: "ResetAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabel", GoMethod: "ResetLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrigin", GoMethod: "ResetOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessAllowRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRules",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesAuthentication",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesAuthenticationIdentities",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesAuthenticationIdentities)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesAuthenticationIdentitiesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesAuthenticationIdentitiesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesAuthenticationIdentitiesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesAuthenticationIdentitiesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesAuthenticationIdentitiesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberProperty{JsiiProperty: "principalIdInput", GoGetter: "PrincipalIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "principalType", GoGetter: "PrincipalType"},
			_jsii_.MemberProperty{JsiiProperty: "principalTypeInput", GoGetter: "PrincipalTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalId", GoMethod: "ResetPrincipalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalType", GoMethod: "ResetPrincipalType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesAuthenticationIdentitiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesAuthenticationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesAuthenticationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "identities", GoGetter: "Identities"},
			_jsii_.MemberProperty{JsiiProperty: "identitiesInput", GoGetter: "IdentitiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "identityType", GoGetter: "IdentityType"},
			_jsii_.MemberProperty{JsiiProperty: "identityTypeInput", GoGetter: "IdentityTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIdentities", GoMethod: "PutIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentities", GoMethod: "ResetIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityType", GoMethod: "ResetIdentityType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestination",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAccountApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAccountApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAccountApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAccountApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAccountApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAccountDatabricksOne",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAccountDatabricksOne)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAccountDatabricksOneOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAccountDatabricksOneOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAccountDatabricksOneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAccountUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAccountUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAccountUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAccountUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAccountUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAppsRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAppsRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAppsRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAppsRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationAppsRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationLakebaseRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationLakebaseRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationLakebaseRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationLakebaseRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationLakebaseRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountApi", GoGetter: "AccountApi"},
			_jsii_.MemberProperty{JsiiProperty: "accountApiInput", GoGetter: "AccountApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOne", GoGetter: "AccountDatabricksOne"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOneInput", GoGetter: "AccountDatabricksOneInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountUi", GoGetter: "AccountUi"},
			_jsii_.MemberProperty{JsiiProperty: "accountUiInput", GoGetter: "AccountUiInput"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntime", GoGetter: "AppsRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntimeInput", GoGetter: "AppsRuntimeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntime", GoGetter: "LakebaseRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntimeInput", GoGetter: "LakebaseRuntimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountApi", GoMethod: "PutAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountDatabricksOne", GoMethod: "PutAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountUi", GoMethod: "PutAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "putAppsRuntime", GoMethod: "PutAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putLakebaseRuntime", GoMethod: "PutLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceApi", GoMethod: "PutWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceUi", GoMethod: "PutWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountApi", GoMethod: "ResetAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountDatabricksOne", GoMethod: "ResetAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountUi", GoMethod: "ResetAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppsRuntime", GoMethod: "ResetAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetLakebaseRuntime", GoMethod: "ResetLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceApi", GoMethod: "ResetWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceUi", GoMethod: "ResetWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApi", GoGetter: "WorkspaceApi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApiInput", GoGetter: "WorkspaceApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUi", GoGetter: "WorkspaceUi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUiInput", GoGetter: "WorkspaceUiInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationWorkspaceApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationWorkspaceApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationWorkspaceApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationWorkspaceApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationWorkspaceApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationWorkspaceUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationWorkspaceUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationWorkspaceUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationWorkspaceUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesDestinationWorkspaceUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesOrigin",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesOrigin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesOriginEndpoints",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesOriginEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesOriginEndpointsOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesOriginEndpointsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endpointIds", GoGetter: "EndpointIds"},
			_jsii_.MemberProperty{JsiiProperty: "endpointIdsInput", GoGetter: "EndpointIdsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetEndpointIds", GoMethod: "ResetEndpointIds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesOriginEndpointsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesOriginOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesOriginOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allPrivateAccess", GoGetter: "AllPrivateAccess"},
			_jsii_.MemberProperty{JsiiProperty: "allPrivateAccessInput", GoGetter: "AllPrivateAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "allRegisteredEndpoints", GoGetter: "AllRegisteredEndpoints"},
			_jsii_.MemberProperty{JsiiProperty: "allRegisteredEndpointsInput", GoGetter: "AllRegisteredEndpointsInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureWorkspacePrivateLink", GoGetter: "AzureWorkspacePrivateLink"},
			_jsii_.MemberProperty{JsiiProperty: "azureWorkspacePrivateLinkInput", GoGetter: "AzureWorkspacePrivateLinkInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
			_jsii_.MemberProperty{JsiiProperty: "endpointsInput", GoGetter: "EndpointsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putEndpoints", GoMethod: "PutEndpoints"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllPrivateAccess", GoMethod: "ResetAllPrivateAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllRegisteredEndpoints", GoMethod: "ResetAllRegisteredEndpoints"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureWorkspacePrivateLink", GoMethod: "ResetAzureWorkspacePrivateLink"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndpoints", GoMethod: "ResetEndpoints"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesOriginOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationInput", GoGetter: "AuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberProperty{JsiiProperty: "originInput", GoGetter: "OriginInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthentication", GoMethod: "PutAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "putDestination", GoMethod: "PutDestination"},
			_jsii_.MemberMethod{JsiiMethod: "putOrigin", GoMethod: "PutOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthentication", GoMethod: "ResetAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabel", GoMethod: "ResetLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrigin", GoMethod: "ResetOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessDenyRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowRules", GoGetter: "AllowRules"},
			_jsii_.MemberProperty{JsiiProperty: "allowRulesInput", GoGetter: "AllowRulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "denyRules", GoGetter: "DenyRules"},
			_jsii_.MemberProperty{JsiiProperty: "denyRulesInput", GoGetter: "DenyRulesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putAllowRules", GoMethod: "PutAllowRules"},
			_jsii_.MemberMethod{JsiiMethod: "putDenyRules", GoMethod: "PutDenyRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowRules", GoMethod: "ResetAllowRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetDenyRules", GoMethod: "ResetDenyRules"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionMode", GoGetter: "RestrictionMode"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionModeInput", GoGetter: "RestrictionModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPrivateAccessOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccess",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccess)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRules",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesAuthentication",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesAuthenticationIdentities",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesAuthenticationIdentities)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesAuthenticationIdentitiesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesAuthenticationIdentitiesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesAuthenticationIdentitiesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesAuthenticationIdentitiesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesAuthenticationIdentitiesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberProperty{JsiiProperty: "principalIdInput", GoGetter: "PrincipalIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "principalType", GoGetter: "PrincipalType"},
			_jsii_.MemberProperty{JsiiProperty: "principalTypeInput", GoGetter: "PrincipalTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalId", GoMethod: "ResetPrincipalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalType", GoMethod: "ResetPrincipalType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesAuthenticationIdentitiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesAuthenticationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesAuthenticationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "identities", GoGetter: "Identities"},
			_jsii_.MemberProperty{JsiiProperty: "identitiesInput", GoGetter: "IdentitiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "identityType", GoGetter: "IdentityType"},
			_jsii_.MemberProperty{JsiiProperty: "identityTypeInput", GoGetter: "IdentityTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIdentities", GoMethod: "PutIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentities", GoMethod: "ResetIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityType", GoMethod: "ResetIdentityType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestination",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountDatabricksOne",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountDatabricksOne)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountDatabricksOneOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountDatabricksOneOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountDatabricksOneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAccountUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAppsRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAppsRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAppsRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAppsRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationAppsRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationLakebaseRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationLakebaseRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationLakebaseRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationLakebaseRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationLakebaseRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountApi", GoGetter: "AccountApi"},
			_jsii_.MemberProperty{JsiiProperty: "accountApiInput", GoGetter: "AccountApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOne", GoGetter: "AccountDatabricksOne"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOneInput", GoGetter: "AccountDatabricksOneInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountUi", GoGetter: "AccountUi"},
			_jsii_.MemberProperty{JsiiProperty: "accountUiInput", GoGetter: "AccountUiInput"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntime", GoGetter: "AppsRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntimeInput", GoGetter: "AppsRuntimeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntime", GoGetter: "LakebaseRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntimeInput", GoGetter: "LakebaseRuntimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountApi", GoMethod: "PutAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountDatabricksOne", GoMethod: "PutAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountUi", GoMethod: "PutAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "putAppsRuntime", GoMethod: "PutAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putLakebaseRuntime", GoMethod: "PutLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceApi", GoMethod: "PutWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceUi", GoMethod: "PutWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountApi", GoMethod: "ResetAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountDatabricksOne", GoMethod: "ResetAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountUi", GoMethod: "ResetAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppsRuntime", GoMethod: "ResetAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetLakebaseRuntime", GoMethod: "ResetLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceApi", GoMethod: "ResetWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceUi", GoMethod: "ResetWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApi", GoGetter: "WorkspaceApi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApiInput", GoGetter: "WorkspaceApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUi", GoGetter: "WorkspaceUi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUiInput", GoGetter: "WorkspaceUiInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationWorkspaceApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationWorkspaceApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationWorkspaceApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationWorkspaceApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationWorkspaceApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationWorkspaceUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationWorkspaceUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationWorkspaceUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationWorkspaceUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesDestinationWorkspaceUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOrigin",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOrigin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOriginExcludedIpRanges",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOriginExcludedIpRanges)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOriginExcludedIpRangesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOriginExcludedIpRangesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ipRanges", GoGetter: "IpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "ipRangesInput", GoGetter: "IpRangesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpRanges", GoMethod: "ResetIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOriginExcludedIpRangesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOriginIncludedIpRanges",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOriginIncludedIpRanges)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOriginIncludedIpRangesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOriginIncludedIpRangesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ipRanges", GoGetter: "IpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "ipRangesInput", GoGetter: "IpRangesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpRanges", GoMethod: "ResetIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOriginIncludedIpRangesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOriginOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOriginOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allIpRanges", GoGetter: "AllIpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "allIpRangesInput", GoGetter: "AllIpRangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludedIpRanges", GoGetter: "ExcludedIpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "excludedIpRangesInput", GoGetter: "ExcludedIpRangesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includedIpRanges", GoGetter: "IncludedIpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "includedIpRangesInput", GoGetter: "IncludedIpRangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putExcludedIpRanges", GoMethod: "PutExcludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "putIncludedIpRanges", GoMethod: "PutIncludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllIpRanges", GoMethod: "ResetAllIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludedIpRanges", GoMethod: "ResetExcludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludedIpRanges", GoMethod: "ResetIncludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOriginOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationInput", GoGetter: "AuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberProperty{JsiiProperty: "originInput", GoGetter: "OriginInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthentication", GoMethod: "PutAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "putDestination", GoMethod: "PutDestination"},
			_jsii_.MemberMethod{JsiiMethod: "putOrigin", GoMethod: "PutOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthentication", GoMethod: "ResetAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabel", GoMethod: "ResetLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrigin", GoMethod: "ResetOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessAllowRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRules",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesAuthentication",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesAuthenticationIdentities",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesAuthenticationIdentities)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesAuthenticationIdentitiesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesAuthenticationIdentitiesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesAuthenticationIdentitiesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesAuthenticationIdentitiesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesAuthenticationIdentitiesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberProperty{JsiiProperty: "principalIdInput", GoGetter: "PrincipalIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "principalType", GoGetter: "PrincipalType"},
			_jsii_.MemberProperty{JsiiProperty: "principalTypeInput", GoGetter: "PrincipalTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalId", GoMethod: "ResetPrincipalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrincipalType", GoMethod: "ResetPrincipalType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesAuthenticationIdentitiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesAuthenticationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesAuthenticationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "identities", GoGetter: "Identities"},
			_jsii_.MemberProperty{JsiiProperty: "identitiesInput", GoGetter: "IdentitiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "identityType", GoGetter: "IdentityType"},
			_jsii_.MemberProperty{JsiiProperty: "identityTypeInput", GoGetter: "IdentityTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIdentities", GoMethod: "PutIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentities", GoMethod: "ResetIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityType", GoMethod: "ResetIdentityType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestination",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAccountApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAccountApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAccountApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAccountApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAccountApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAccountDatabricksOne",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAccountDatabricksOne)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAccountDatabricksOneOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAccountDatabricksOneOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAccountDatabricksOneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAccountUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAccountUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAccountUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAccountUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAccountUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAppsRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAppsRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAppsRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAppsRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationAppsRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationLakebaseRuntime",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationLakebaseRuntime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationLakebaseRuntimeOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationLakebaseRuntimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationLakebaseRuntimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountApi", GoGetter: "AccountApi"},
			_jsii_.MemberProperty{JsiiProperty: "accountApiInput", GoGetter: "AccountApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOne", GoGetter: "AccountDatabricksOne"},
			_jsii_.MemberProperty{JsiiProperty: "accountDatabricksOneInput", GoGetter: "AccountDatabricksOneInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountUi", GoGetter: "AccountUi"},
			_jsii_.MemberProperty{JsiiProperty: "accountUiInput", GoGetter: "AccountUiInput"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntime", GoGetter: "AppsRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "appsRuntimeInput", GoGetter: "AppsRuntimeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntime", GoGetter: "LakebaseRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "lakebaseRuntimeInput", GoGetter: "LakebaseRuntimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountApi", GoMethod: "PutAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountDatabricksOne", GoMethod: "PutAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountUi", GoMethod: "PutAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "putAppsRuntime", GoMethod: "PutAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putLakebaseRuntime", GoMethod: "PutLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceApi", GoMethod: "PutWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceUi", GoMethod: "PutWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountApi", GoMethod: "ResetAccountApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountDatabricksOne", GoMethod: "ResetAccountDatabricksOne"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountUi", GoMethod: "ResetAccountUi"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppsRuntime", GoMethod: "ResetAppsRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetLakebaseRuntime", GoMethod: "ResetLakebaseRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceApi", GoMethod: "ResetWorkspaceApi"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceUi", GoMethod: "ResetWorkspaceUi"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApi", GoGetter: "WorkspaceApi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceApiInput", GoGetter: "WorkspaceApiInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUi", GoGetter: "WorkspaceUi"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUiInput", GoGetter: "WorkspaceUiInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationWorkspaceApi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationWorkspaceApi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationWorkspaceApiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationWorkspaceApiOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetScopeQualifier", GoMethod: "ResetScopeQualifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifier", GoGetter: "ScopeQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "scopeQualifierInput", GoGetter: "ScopeQualifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationWorkspaceApiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationWorkspaceUi",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationWorkspaceUi)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationWorkspaceUiOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationWorkspaceUiOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allDestinations", GoGetter: "AllDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "allDestinationsInput", GoGetter: "AllDestinationsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllDestinations", GoMethod: "ResetAllDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesDestinationWorkspaceUiOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOrigin",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOrigin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOriginExcludedIpRanges",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOriginExcludedIpRanges)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOriginExcludedIpRangesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOriginExcludedIpRangesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ipRanges", GoGetter: "IpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "ipRangesInput", GoGetter: "IpRangesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpRanges", GoMethod: "ResetIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOriginExcludedIpRangesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOriginIncludedIpRanges",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOriginIncludedIpRanges)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOriginIncludedIpRangesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOriginIncludedIpRangesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ipRanges", GoGetter: "IpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "ipRangesInput", GoGetter: "IpRangesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpRanges", GoMethod: "ResetIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOriginIncludedIpRangesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOriginOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOriginOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allIpRanges", GoGetter: "AllIpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "allIpRangesInput", GoGetter: "AllIpRangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludedIpRanges", GoGetter: "ExcludedIpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "excludedIpRangesInput", GoGetter: "ExcludedIpRangesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includedIpRanges", GoGetter: "IncludedIpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "includedIpRangesInput", GoGetter: "IncludedIpRangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putExcludedIpRanges", GoMethod: "PutExcludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "putIncludedIpRanges", GoMethod: "PutIncludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllIpRanges", GoMethod: "ResetAllIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludedIpRanges", GoMethod: "ResetExcludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludedIpRanges", GoMethod: "ResetIncludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOriginOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationInput", GoGetter: "AuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberProperty{JsiiProperty: "originInput", GoGetter: "OriginInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthentication", GoMethod: "PutAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "putDestination", GoMethod: "PutDestination"},
			_jsii_.MemberMethod{JsiiMethod: "putOrigin", GoMethod: "PutOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthentication", GoMethod: "ResetAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabel", GoMethod: "ResetLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrigin", GoMethod: "ResetOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessDenyRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowRules", GoGetter: "AllowRules"},
			_jsii_.MemberProperty{JsiiProperty: "allowRulesInput", GoGetter: "AllowRulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "denyRules", GoGetter: "DenyRules"},
			_jsii_.MemberProperty{JsiiProperty: "denyRulesInput", GoGetter: "DenyRulesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putAllowRules", GoMethod: "PutAllowRules"},
			_jsii_.MemberMethod{JsiiMethod: "putDenyRules", GoMethod: "PutDenyRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowRules", GoMethod: "ResetAllowRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetDenyRules", GoMethod: "ResetDenyRules"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionMode", GoGetter: "RestrictionMode"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionModeInput", GoGetter: "RestrictionModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsIngressPublicAccessOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsList",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsList)(nil)).Elem(),
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
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-databricks.dataDatabricksAccountNetworkPolicies.DataDatabricksAccountNetworkPoliciesItemsOutputReference",
		reflect.TypeOf((*DataDatabricksAccountNetworkPoliciesItemsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "egress", GoGetter: "Egress"},
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
			_jsii_.MemberProperty{JsiiProperty: "ingress", GoGetter: "Ingress"},
			_jsii_.MemberProperty{JsiiProperty: "ingressDryRun", GoGetter: "IngressDryRun"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "networkPolicyId", GoGetter: "NetworkPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "networkPolicyIdInput", GoGetter: "NetworkPolicyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksAccountNetworkPoliciesItemsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
