// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package onlinetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/onlinetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList
type jsiiProxy_OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewOnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList {
	_init_.Initialize()

	if err := validateNewOnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList{}

	_jsii_.Create(
		"@cdktn/provider-databricks.onlineTable.OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewOnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList_Override(o OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.onlineTable.OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		o,
	)
}

func (j *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (o *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := o.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		o,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList) Get(index *float64) OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressOutputReference {
	if err := o.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressOutputReference

	_jsii_.Invoke(
		o,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusTriggeredUpdateProgressList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

