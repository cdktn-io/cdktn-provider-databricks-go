// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package disasterrecoveryfailovergroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/disasterrecoveryfailovergroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList interface {
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
	Get(index *float64) DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList
type jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList {
	_init_.Initialize()

	if err := validateNewDisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList{}

	_jsii_.Create(
		"@cdktn/provider-databricks.disasterRecoveryFailoverGroup.DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList_Override(d DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.disasterRecoveryFailoverGroup.DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList) Get(index *float64) DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DisasterRecoveryFailoverGroupUnityCatalogAssetsLocationMappingsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

