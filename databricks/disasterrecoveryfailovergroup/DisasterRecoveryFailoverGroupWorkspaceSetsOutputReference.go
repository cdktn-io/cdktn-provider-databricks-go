// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package disasterrecoveryfailovergroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/disasterrecoveryfailovergroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
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

// The jsii proxy struct for DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference
type jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) ReplicateWorkspaceAssets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicateWorkspaceAssets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) ReplicateWorkspaceAssetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicateWorkspaceAssetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) StableUrlNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stableUrlNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) StableUrlNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stableUrlNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) WorkspaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"workspaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) WorkspaceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"workspaceIdsInput",
		&returns,
	)
	return returns
}


func NewDisasterRecoveryFailoverGroupWorkspaceSetsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference {
	_init_.Initialize()

	if err := validateNewDisasterRecoveryFailoverGroupWorkspaceSetsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.disasterRecoveryFailoverGroup.DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDisasterRecoveryFailoverGroupWorkspaceSetsOutputReference_Override(d DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.disasterRecoveryFailoverGroup.DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference)SetReplicateWorkspaceAssets(val interface{}) {
	if err := j.validateSetReplicateWorkspaceAssetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicateWorkspaceAssets",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference)SetStableUrlNames(val *[]*string) {
	if err := j.validateSetStableUrlNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stableUrlNames",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference)SetWorkspaceIds(val *[]*string) {
	if err := j.validateSetWorkspaceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceIds",
		val,
	)
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) ResetStableUrlNames() {
	_jsii_.InvokeVoid(
		d,
		"resetStableUrlNames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroupWorkspaceSetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

