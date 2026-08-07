// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgressyncedtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/postgressyncedtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PostgresSyncedTableSpecExtraColumnsOutputReference interface {
	cdktn.ComplexObject
	ColumnName() *string
	SetColumnName(val *string)
	ColumnNameInput() *string
	ColumnType() *string
	SetColumnType(val *string)
	ColumnTypeInput() *string
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
	Compute() *string
	SetCompute(val *string)
	ComputeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Maintenance() *string
	SetMaintenance(val *string)
	MaintenanceInput() *string
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
	ResetCompute()
	ResetMaintenance()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PostgresSyncedTableSpecExtraColumnsOutputReference
type jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) ColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) ColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) ColumnType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) ColumnTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) Compute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) ComputeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) Maintenance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) MaintenanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPostgresSyncedTableSpecExtraColumnsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PostgresSyncedTableSpecExtraColumnsOutputReference {
	_init_.Initialize()

	if err := validateNewPostgresSyncedTableSpecExtraColumnsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresSyncedTable.PostgresSyncedTableSpecExtraColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPostgresSyncedTableSpecExtraColumnsOutputReference_Override(p PostgresSyncedTableSpecExtraColumnsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresSyncedTable.PostgresSyncedTableSpecExtraColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference)SetColumnName(val *string) {
	if err := j.validateSetColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnName",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference)SetColumnType(val *string) {
	if err := j.validateSetColumnTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnType",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference)SetCompute(val *string) {
	if err := j.validateSetComputeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compute",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference)SetMaintenance(val *string) {
	if err := j.validateSetMaintenanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenance",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) ResetCompute() {
	_jsii_.InvokeVoid(
		p,
		"resetCompute",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) ResetMaintenance() {
	_jsii_.InvokeVoid(
		p,
		"resetMaintenance",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PostgresSyncedTableSpecExtraColumnsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

