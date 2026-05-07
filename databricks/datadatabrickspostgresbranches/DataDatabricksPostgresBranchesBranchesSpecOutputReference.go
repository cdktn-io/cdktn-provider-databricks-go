// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresbranches

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabrickspostgresbranches/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksPostgresBranchesBranchesSpecOutputReference interface {
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
	ExpireTime() *string
	SetExpireTime(val *string)
	ExpireTimeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksPostgresBranchesBranchesSpec
	SetInternalValue(val *DataDatabricksPostgresBranchesBranchesSpec)
	IsProtected() interface{}
	SetIsProtected(val interface{})
	IsProtectedInput() interface{}
	NoExpiry() interface{}
	SetNoExpiry(val interface{})
	NoExpiryInput() interface{}
	SourceBranch() *string
	SetSourceBranch(val *string)
	SourceBranchInput() *string
	SourceBranchLsn() *string
	SetSourceBranchLsn(val *string)
	SourceBranchLsnInput() *string
	SourceBranchTime() *string
	SetSourceBranchTime(val *string)
	SourceBranchTimeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Ttl() *string
	SetTtl(val *string)
	TtlInput() *string
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
	ResetExpireTime()
	ResetIsProtected()
	ResetNoExpiry()
	ResetSourceBranch()
	ResetSourceBranchLsn()
	ResetSourceBranchTime()
	ResetTtl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksPostgresBranchesBranchesSpecOutputReference
type jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) ExpireTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expireTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) ExpireTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expireTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) InternalValue() *DataDatabricksPostgresBranchesBranchesSpec {
	var returns *DataDatabricksPostgresBranchesBranchesSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) IsProtected() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isProtected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) IsProtectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isProtectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) NoExpiry() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) NoExpiryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) SourceBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) SourceBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) SourceBranchLsn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBranchLsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) SourceBranchLsnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBranchLsnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) SourceBranchTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBranchTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) SourceBranchTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBranchTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) TtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksPostgresBranchesBranchesSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksPostgresBranchesBranchesSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksPostgresBranchesBranchesSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresBranches.DataDatabricksPostgresBranchesBranchesSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksPostgresBranchesBranchesSpecOutputReference_Override(d DataDatabricksPostgresBranchesBranchesSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresBranches.DataDatabricksPostgresBranchesBranchesSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference)SetExpireTime(val *string) {
	if err := j.validateSetExpireTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expireTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference)SetInternalValue(val *DataDatabricksPostgresBranchesBranchesSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference)SetIsProtected(val interface{}) {
	if err := j.validateSetIsProtectedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isProtected",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference)SetNoExpiry(val interface{}) {
	if err := j.validateSetNoExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noExpiry",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference)SetSourceBranch(val *string) {
	if err := j.validateSetSourceBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceBranch",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference)SetSourceBranchLsn(val *string) {
	if err := j.validateSetSourceBranchLsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceBranchLsn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference)SetSourceBranchTime(val *string) {
	if err := j.validateSetSourceBranchTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceBranchTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference)SetTtl(val *string) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) ResetExpireTime() {
	_jsii_.InvokeVoid(
		d,
		"resetExpireTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) ResetIsProtected() {
	_jsii_.InvokeVoid(
		d,
		"resetIsProtected",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) ResetNoExpiry() {
	_jsii_.InvokeVoid(
		d,
		"resetNoExpiry",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) ResetSourceBranch() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceBranch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) ResetSourceBranchLsn() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceBranchLsn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) ResetSourceBranchTime() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceBranchTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) ResetTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresBranchesBranchesSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

