// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspaceiamworkspaceassignmentsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksworkspaceiamworkspaceassignmentsv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference interface {
	cdktn.ComplexObject
	AccountId() *string
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
	EffectiveEntitlements() *[]*string
	Entitlements() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignments
	SetInternalValue(val *DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignments)
	PrincipalId() *float64
	SetPrincipalId(val *float64)
	PrincipalIdInput() *float64
	PrincipalType() *string
	ProviderConfig() DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsProviderConfigOutputReference
	ProviderConfigInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WorkspaceId() *float64
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
	PutProviderConfig(value *DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsProviderConfig)
	ResetProviderConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference
type jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) EffectiveEntitlements() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"effectiveEntitlements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) Entitlements() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entitlements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) InternalValue() *DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignments {
	var returns *DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignments
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) PrincipalId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) PrincipalIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"principalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) PrincipalType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) ProviderConfig() DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsProviderConfigOutputReference {
	var returns DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) WorkspaceId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}


func NewDataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksWorkspaceIamWorkspaceAssignmentsV2.DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference_Override(d DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksWorkspaceIamWorkspaceAssignmentsV2.DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference)SetInternalValue(val *DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignments) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference)SetPrincipalId(val *float64) {
	if err := j.validateSetPrincipalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principalId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) PutProviderConfig(value *DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsProviderConfig) {
	if err := d.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksWorkspaceIamWorkspaceAssignmentsV2WorkspaceAssignmentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

