// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/datadatabrickspostgresrole/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksPostgresRoleSpecOutputReference interface {
	cdktn.ComplexObject
	Attributes() DataDatabricksPostgresRoleSpecAttributesOutputReference
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
	InternalValue() *DataDatabricksPostgresRoleSpec
	SetInternalValue(val *DataDatabricksPostgresRoleSpec)
	MembershipRoles() *[]*string
	SetMembershipRoles(val *[]*string)
	MembershipRolesInput() *[]*string
	PostgresRole() *string
	SetPostgresRole(val *string)
	PostgresRoleInput() *string
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
	PutAttributes(value *DataDatabricksPostgresRoleSpecAttributes)
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

// The jsii proxy struct for DataDatabricksPostgresRoleSpecOutputReference
type jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) Attributes() DataDatabricksPostgresRoleSpecAttributesOutputReference {
	var returns DataDatabricksPostgresRoleSpecAttributesOutputReference
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) AttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) AuthMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) AuthMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) IdentityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) IdentityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) InternalValue() *DataDatabricksPostgresRoleSpec {
	var returns *DataDatabricksPostgresRoleSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) MembershipRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"membershipRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) MembershipRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"membershipRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) PostgresRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) PostgresRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksPostgresRoleSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksPostgresRoleSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksPostgresRoleSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresRole.DataDatabricksPostgresRoleSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksPostgresRoleSpecOutputReference_Override(d DataDatabricksPostgresRoleSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresRole.DataDatabricksPostgresRoleSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference)SetAuthMethod(val *string) {
	if err := j.validateSetAuthMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authMethod",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference)SetIdentityType(val *string) {
	if err := j.validateSetIdentityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference)SetInternalValue(val *DataDatabricksPostgresRoleSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference)SetMembershipRoles(val *[]*string) {
	if err := j.validateSetMembershipRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"membershipRoles",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference)SetPostgresRole(val *string) {
	if err := j.validateSetPostgresRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postgresRole",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) PutAttributes(value *DataDatabricksPostgresRoleSpecAttributes) {
	if err := d.validatePutAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) ResetAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) ResetAuthMethod() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthMethod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) ResetIdentityType() {
	_jsii_.InvokeVoid(
		d,
		"resetIdentityType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) ResetMembershipRoles() {
	_jsii_.InvokeVoid(
		d,
		"resetMembershipRoles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) ResetPostgresRole() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgresRole",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresRoleSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

