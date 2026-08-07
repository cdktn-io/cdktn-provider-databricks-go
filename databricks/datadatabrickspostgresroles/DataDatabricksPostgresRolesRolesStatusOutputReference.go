// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresroles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabrickspostgresroles/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksPostgresRolesRolesStatusOutputReference interface {
	cdktn.ComplexObject
	Attributes() DataDatabricksPostgresRolesRolesStatusAttributesOutputReference
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
	InternalValue() *DataDatabricksPostgresRolesRolesStatus
	SetInternalValue(val *DataDatabricksPostgresRolesRolesStatus)
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
	PutAttributes(value *DataDatabricksPostgresRolesRolesStatusAttributes)
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

// The jsii proxy struct for DataDatabricksPostgresRolesRolesStatusOutputReference
type jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) Attributes() DataDatabricksPostgresRolesRolesStatusAttributesOutputReference {
	var returns DataDatabricksPostgresRolesRolesStatusAttributesOutputReference
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) AttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) AuthMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) AuthMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) IdentityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) IdentityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) InternalValue() *DataDatabricksPostgresRolesRolesStatus {
	var returns *DataDatabricksPostgresRolesRolesStatus
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) MembershipRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"membershipRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) MembershipRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"membershipRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) PostgresRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) PostgresRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) RoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksPostgresRolesRolesStatusOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksPostgresRolesRolesStatusOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksPostgresRolesRolesStatusOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresRoles.DataDatabricksPostgresRolesRolesStatusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksPostgresRolesRolesStatusOutputReference_Override(d DataDatabricksPostgresRolesRolesStatusOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresRoles.DataDatabricksPostgresRolesRolesStatusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference)SetAuthMethod(val *string) {
	if err := j.validateSetAuthMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authMethod",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference)SetIdentityType(val *string) {
	if err := j.validateSetIdentityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference)SetInternalValue(val *DataDatabricksPostgresRolesRolesStatus) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference)SetMembershipRoles(val *[]*string) {
	if err := j.validateSetMembershipRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"membershipRoles",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference)SetPostgresRole(val *string) {
	if err := j.validateSetPostgresRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postgresRole",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) PutAttributes(value *DataDatabricksPostgresRolesRolesStatusAttributes) {
	if err := d.validatePutAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) ResetAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) ResetAuthMethod() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthMethod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) ResetIdentityType() {
	_jsii_.InvokeVoid(
		d,
		"resetIdentityType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) ResetMembershipRoles() {
	_jsii_.InvokeVoid(
		d,
		"resetMembershipRoles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) ResetPostgresRole() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgresRole",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresRolesRolesStatusOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

