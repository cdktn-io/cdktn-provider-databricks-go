// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/postgresrole/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PostgresRoleStatusOutputReference interface {
	cdktn.ComplexObject
	Attributes() PostgresRoleStatusAttributesOutputReference
	AttributesInput() interface{}
	AuthMethod() *string
	SetAuthMethod(val *string)
	AuthMethodInput() *string
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
	IdentityType() *string
	SetIdentityType(val *string)
	IdentityTypeInput() *string
	InternalValue() *PostgresRoleStatus
	SetInternalValue(val *PostgresRoleStatus)
	MembershipRoles() *[]*string
	SetMembershipRoles(val *[]*string)
	MembershipRolesInput() *[]*string
	PostgresRole() *string
	SetPostgresRole(val *string)
	PostgresRoleInput() *string
	RoleId() *string
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
	PutAttributes(value *PostgresRoleStatusAttributes)
	ResetAttributes()
	ResetAuthMethod()
	ResetIdentityType()
	ResetMembershipRoles()
	ResetPostgresRole()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PostgresRoleStatusOutputReference
type jsiiProxy_PostgresRoleStatusOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) Attributes() PostgresRoleStatusAttributesOutputReference {
	var returns PostgresRoleStatusAttributesOutputReference
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) AttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) AuthMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) AuthMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) IdentityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) IdentityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) InternalValue() *PostgresRoleStatus {
	var returns *PostgresRoleStatus
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) MembershipRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"membershipRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) MembershipRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"membershipRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) PostgresRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) PostgresRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) RoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPostgresRoleStatusOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PostgresRoleStatusOutputReference {
	_init_.Initialize()

	if err := validateNewPostgresRoleStatusOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PostgresRoleStatusOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresRole.PostgresRoleStatusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPostgresRoleStatusOutputReference_Override(p PostgresRoleStatusOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresRole.PostgresRoleStatusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference)SetAuthMethod(val *string) {
	if err := j.validateSetAuthMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authMethod",
		val,
	)
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference)SetIdentityType(val *string) {
	if err := j.validateSetIdentityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityType",
		val,
	)
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference)SetInternalValue(val *PostgresRoleStatus) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference)SetMembershipRoles(val *[]*string) {
	if err := j.validateSetMembershipRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"membershipRoles",
		val,
	)
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference)SetPostgresRole(val *string) {
	if err := j.validateSetPostgresRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postgresRole",
		val,
	)
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PostgresRoleStatusOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PostgresRoleStatusOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresRoleStatusOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PostgresRoleStatusOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PostgresRoleStatusOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PostgresRoleStatusOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PostgresRoleStatusOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PostgresRoleStatusOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PostgresRoleStatusOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PostgresRoleStatusOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PostgresRoleStatusOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PostgresRoleStatusOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresRoleStatusOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PostgresRoleStatusOutputReference) PutAttributes(value *PostgresRoleStatusAttributes) {
	if err := p.validatePutAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAttributes",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresRoleStatusOutputReference) ResetAttributes() {
	_jsii_.InvokeVoid(
		p,
		"resetAttributes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresRoleStatusOutputReference) ResetAuthMethod() {
	_jsii_.InvokeVoid(
		p,
		"resetAuthMethod",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresRoleStatusOutputReference) ResetIdentityType() {
	_jsii_.InvokeVoid(
		p,
		"resetIdentityType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresRoleStatusOutputReference) ResetMembershipRoles() {
	_jsii_.InvokeVoid(
		p,
		"resetMembershipRoles",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresRoleStatusOutputReference) ResetPostgresRole() {
	_jsii_.InvokeVoid(
		p,
		"resetPostgresRole",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresRoleStatusOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PostgresRoleStatusOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

