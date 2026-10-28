// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package notificationdestination

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/notificationdestination/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NotificationDestinationConfigDatabricksAppSlackOutputReference interface {
	cdktn.ComplexObject
	ChannelId() *string
	SetChannelId(val *string)
	ChannelIdInput() *string
	ChannelIdSet() interface{}
	SetChannelIdSet(val interface{})
	ChannelIdSetInput() interface{}
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
	InternalValue() *NotificationDestinationConfigDatabricksAppSlack
	SetInternalValue(val *NotificationDestinationConfigDatabricksAppSlack)
	TeamId() *string
	SetTeamId(val *string)
	TeamIdInput() *string
	TeamIdSet() interface{}
	SetTeamIdSet(val interface{})
	TeamIdSetInput() interface{}
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
	ResetChannelId()
	ResetChannelIdSet()
	ResetTeamId()
	ResetTeamIdSet()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NotificationDestinationConfigDatabricksAppSlackOutputReference
type jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) ChannelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) ChannelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) ChannelIdSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"channelIdSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) ChannelIdSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"channelIdSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) InternalValue() *NotificationDestinationConfigDatabricksAppSlack {
	var returns *NotificationDestinationConfigDatabricksAppSlack
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) TeamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) TeamIdSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"teamIdSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) TeamIdSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"teamIdSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNotificationDestinationConfigDatabricksAppSlackOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) NotificationDestinationConfigDatabricksAppSlackOutputReference {
	_init_.Initialize()

	if err := validateNewNotificationDestinationConfigDatabricksAppSlackOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.notificationDestination.NotificationDestinationConfigDatabricksAppSlackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNotificationDestinationConfigDatabricksAppSlackOutputReference_Override(n NotificationDestinationConfigDatabricksAppSlackOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.notificationDestination.NotificationDestinationConfigDatabricksAppSlackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference)SetChannelId(val *string) {
	if err := j.validateSetChannelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelId",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference)SetChannelIdSet(val interface{}) {
	if err := j.validateSetChannelIdSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelIdSet",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference)SetInternalValue(val *NotificationDestinationConfigDatabricksAppSlack) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference)SetTeamId(val *string) {
	if err := j.validateSetTeamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference)SetTeamIdSet(val interface{}) {
	if err := j.validateSetTeamIdSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamIdSet",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) ResetChannelId() {
	_jsii_.InvokeVoid(
		n,
		"resetChannelId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) ResetChannelIdSet() {
	_jsii_.InvokeVoid(
		n,
		"resetChannelIdSet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) ResetTeamId() {
	_jsii_.InvokeVoid(
		n,
		"resetTeamId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) ResetTeamIdSet() {
	_jsii_.InvokeVoid(
		n,
		"resetTeamIdSet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationDestinationConfigDatabricksAppSlackOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

