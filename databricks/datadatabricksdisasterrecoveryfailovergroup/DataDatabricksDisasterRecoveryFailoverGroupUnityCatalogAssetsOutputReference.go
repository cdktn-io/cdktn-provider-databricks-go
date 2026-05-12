// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdisasterrecoveryfailovergroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabricksdisasterrecoveryfailovergroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference interface {
	cdktn.ComplexObject
	Catalogs() DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsCatalogsList
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
	InternalValue() *DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssets
	SetInternalValue(val *DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssets)
	LocationMappings() DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList
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

// The jsii proxy struct for DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference
type jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) Catalogs() DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsCatalogsList {
	var returns DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsCatalogsList
	_jsii_.Get(
		j,
		"catalogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) CatalogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) DataReplicationWorkspaceSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataReplicationWorkspaceSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) DataReplicationWorkspaceSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataReplicationWorkspaceSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) InternalValue() *DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssets {
	var returns *DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssets
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) LocationMappings() DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList {
	var returns DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList
	_jsii_.Get(
		j,
		"locationMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) LocationMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locationMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksDisasterRecoveryFailoverGroup.DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference_Override(d DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksDisasterRecoveryFailoverGroup.DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference)SetDataReplicationWorkspaceSet(val *string) {
	if err := j.validateSetDataReplicationWorkspaceSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataReplicationWorkspaceSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference)SetInternalValue(val *DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssets) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) PutCatalogs(value interface{}) {
	if err := d.validatePutCatalogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCatalogs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) PutLocationMappings(value interface{}) {
	if err := d.validatePutLocationMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLocationMappings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) ResetLocationMappings() {
	_jsii_.InvokeVoid(
		d,
		"resetLocationMappings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupUnityCatalogAssetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

