// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresbranch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/postgresbranch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PostgresBranchSpecOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
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

// The jsii proxy struct for PostgresBranchSpecOutputReference
type jsiiProxy_PostgresBranchSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) ExpireTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expireTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) ExpireTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expireTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) IsProtected() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isProtected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) IsProtectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isProtectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) NoExpiry() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) NoExpiryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) SourceBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) SourceBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) SourceBranchLsn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBranchLsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) SourceBranchLsnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBranchLsnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) SourceBranchTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBranchTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) SourceBranchTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBranchTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference) TtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}


func NewPostgresBranchSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PostgresBranchSpecOutputReference {
	_init_.Initialize()

	if err := validateNewPostgresBranchSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PostgresBranchSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresBranch.PostgresBranchSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPostgresBranchSpecOutputReference_Override(p PostgresBranchSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresBranch.PostgresBranchSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference)SetExpireTime(val *string) {
	if err := j.validateSetExpireTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expireTime",
		val,
	)
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference)SetIsProtected(val interface{}) {
	if err := j.validateSetIsProtectedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isProtected",
		val,
	)
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference)SetNoExpiry(val interface{}) {
	if err := j.validateSetNoExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noExpiry",
		val,
	)
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference)SetSourceBranch(val *string) {
	if err := j.validateSetSourceBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceBranch",
		val,
	)
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference)SetSourceBranchLsn(val *string) {
	if err := j.validateSetSourceBranchLsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceBranchLsn",
		val,
	)
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference)SetSourceBranchTime(val *string) {
	if err := j.validateSetSourceBranchTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceBranchTime",
		val,
	)
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PostgresBranchSpecOutputReference)SetTtl(val *string) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) ResetExpireTime() {
	_jsii_.InvokeVoid(
		p,
		"resetExpireTime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) ResetIsProtected() {
	_jsii_.InvokeVoid(
		p,
		"resetIsProtected",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) ResetNoExpiry() {
	_jsii_.InvokeVoid(
		p,
		"resetNoExpiry",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) ResetSourceBranch() {
	_jsii_.InvokeVoid(
		p,
		"resetSourceBranch",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) ResetSourceBranchLsn() {
	_jsii_.InvokeVoid(
		p,
		"resetSourceBranchLsn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) ResetSourceBranchTime() {
	_jsii_.InvokeVoid(
		p,
		"resetSourceBranchTime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) ResetTtl() {
	_jsii_.InvokeVoid(
		p,
		"resetTtl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresBranchSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

