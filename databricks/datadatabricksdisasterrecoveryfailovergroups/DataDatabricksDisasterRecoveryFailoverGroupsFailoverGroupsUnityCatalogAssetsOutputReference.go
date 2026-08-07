// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdisasterrecoveryfailovergroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksdisasterrecoveryfailovergroups/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference interface {
	cdktn.ComplexObject
	Catalogs() DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsCatalogsList
	CatalogsInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataReplicationWorkspaceSet() *string
	SetDataReplicationWorkspaceSet(val *string)
	DataReplicationWorkspaceSetInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssets
	SetInternalValue(val *DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssets)
	LocationMappings() DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsLocationMappingsList
	LocationMappingsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutCatalogs(value interface{})
	PutLocationMappings(value interface{})
	ResetLocationMappings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference
type jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) Catalogs() DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsCatalogsList {
	var returns DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsCatalogsList
	_jsii_.Get(
		j,
		"catalogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) CatalogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) DataReplicationWorkspaceSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataReplicationWorkspaceSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) DataReplicationWorkspaceSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataReplicationWorkspaceSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) InternalValue() *DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssets {
	var returns *DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssets
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) LocationMappings() DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsLocationMappingsList {
	var returns DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsLocationMappingsList
	_jsii_.Get(
		j,
		"locationMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) LocationMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locationMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksDisasterRecoveryFailoverGroups.DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference_Override(d DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksDisasterRecoveryFailoverGroups.DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference)SetDataReplicationWorkspaceSet(val *string) {
	if err := j.validateSetDataReplicationWorkspaceSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataReplicationWorkspaceSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference)SetInternalValue(val *DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssets) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) PutCatalogs(value interface{}) {
	if err := d.validatePutCatalogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCatalogs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) PutLocationMappings(value interface{}) {
	if err := d.validatePutLocationMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLocationMappings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) ResetLocationMappings() {
	_jsii_.InvokeVoid(
		d,
		"resetLocationMappings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsUnityCatalogAssetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

