// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksprivatenetworkgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabricksprivatenetworkgateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference interface {
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
	CrossAccountRole() DataDatabricksPrivateNetworkGatewayAwsCloudConnectionCrossAccountRoleOutputReference
	CrossAccountRoleInput() *DataDatabricksPrivateNetworkGatewayAwsCloudConnectionCrossAccountRole
	// Experimental.
	Fqn() *string
	GatewaySubnets() DataDatabricksPrivateNetworkGatewayAwsCloudConnectionGatewaySubnetsList
	GatewaySubnetsInput() interface{}
	InternalValue() *DataDatabricksPrivateNetworkGatewayAwsCloudConnection
	SetInternalValue(val *DataDatabricksPrivateNetworkGatewayAwsCloudConnection)
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
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
	PutCrossAccountRole(value *DataDatabricksPrivateNetworkGatewayAwsCloudConnectionCrossAccountRole)
	PutGatewaySubnets(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference
type jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) CrossAccountRole() DataDatabricksPrivateNetworkGatewayAwsCloudConnectionCrossAccountRoleOutputReference {
	var returns DataDatabricksPrivateNetworkGatewayAwsCloudConnectionCrossAccountRoleOutputReference
	_jsii_.Get(
		j,
		"crossAccountRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) CrossAccountRoleInput() *DataDatabricksPrivateNetworkGatewayAwsCloudConnectionCrossAccountRole {
	var returns *DataDatabricksPrivateNetworkGatewayAwsCloudConnectionCrossAccountRole
	_jsii_.Get(
		j,
		"crossAccountRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) GatewaySubnets() DataDatabricksPrivateNetworkGatewayAwsCloudConnectionGatewaySubnetsList {
	var returns DataDatabricksPrivateNetworkGatewayAwsCloudConnectionGatewaySubnetsList
	_jsii_.Get(
		j,
		"gatewaySubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) GatewaySubnetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewaySubnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) InternalValue() *DataDatabricksPrivateNetworkGatewayAwsCloudConnection {
	var returns *DataDatabricksPrivateNetworkGatewayAwsCloudConnection
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPrivateNetworkGateway.DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference_Override(d DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPrivateNetworkGateway.DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference)SetInternalValue(val *DataDatabricksPrivateNetworkGatewayAwsCloudConnection) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) PutCrossAccountRole(value *DataDatabricksPrivateNetworkGatewayAwsCloudConnectionCrossAccountRole) {
	if err := d.validatePutCrossAccountRoleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCrossAccountRole",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) PutGatewaySubnets(value interface{}) {
	if err := d.validatePutGatewaySubnetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGatewaySubnets",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksPrivateNetworkGatewayAwsCloudConnectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

