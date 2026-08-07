// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdisasterrecoveryfailovergroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksdisasterrecoveryfailovergroups/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference interface {
	cdktn.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSets
	SetInternalValue(val *DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSets)
	Name() *string
	SetName(val *string)
	NameInput() *string
	ReplicateWorkspaceAssets() interface{}
	SetReplicateWorkspaceAssets(val interface{})
	ReplicateWorkspaceAssetsInput() interface{}
	StableUrlNames() *[]*string
	SetStableUrlNames(val *[]*string)
	StableUrlNamesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WorkspaceIds() *[]*string
	SetWorkspaceIds(val *[]*string)
	WorkspaceIdsInput() *[]*string
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
	ResetReplicateWorkspaceAssets()
	ResetStableUrlNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference
type jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) InternalValue() *DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSets {
	var returns *DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSets
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) ReplicateWorkspaceAssets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicateWorkspaceAssets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) ReplicateWorkspaceAssetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicateWorkspaceAssetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) StableUrlNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stableUrlNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) StableUrlNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stableUrlNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) WorkspaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"workspaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) WorkspaceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"workspaceIdsInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksDisasterRecoveryFailoverGroups.DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference_Override(d DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksDisasterRecoveryFailoverGroups.DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference)SetInternalValue(val *DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSets) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference)SetReplicateWorkspaceAssets(val interface{}) {
	if err := j.validateSetReplicateWorkspaceAssetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicateWorkspaceAssets",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference)SetStableUrlNames(val *[]*string) {
	if err := j.validateSetStableUrlNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stableUrlNames",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference)SetWorkspaceIds(val *[]*string) {
	if err := j.validateSetWorkspaceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceIds",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) ResetReplicateWorkspaceAssets() {
	_jsii_.InvokeVoid(
		d,
		"resetReplicateWorkspaceAssets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) ResetStableUrlNames() {
	_jsii_.InvokeVoid(
		d,
		"resetStableUrlNames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksDisasterRecoveryFailoverGroupsFailoverGroupsWorkspaceSetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

