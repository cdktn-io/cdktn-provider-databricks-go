// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksinstancepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/datadatabricksinstancepool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktn.IInterpolatingParent
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList
type jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList {
	_init_.Initialize()

	if err := validateNewDataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksInstancePool.DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList_Override(d DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksInstancePool.DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList) Get(index *float64) DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverrideList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

