// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgressyncedtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/postgressyncedtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PostgresSyncedTableSpecNewPipelineSpecOutputReference interface {
	cdktn.ComplexObject
	BudgetPolicyId() *string
	SetBudgetPolicyId(val *string)
	BudgetPolicyIdInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StorageCatalog() *string
	SetStorageCatalog(val *string)
	StorageCatalogInput() *string
	StorageSchema() *string
	SetStorageSchema(val *string)
	StorageSchemaInput() *string
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
	ResetBudgetPolicyId()
	ResetStorageCatalog()
	ResetStorageSchema()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PostgresSyncedTableSpecNewPipelineSpecOutputReference
type jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) BudgetPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) BudgetPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) StorageCatalog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageCatalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) StorageCatalogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageCatalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) StorageSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) StorageSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPostgresSyncedTableSpecNewPipelineSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PostgresSyncedTableSpecNewPipelineSpecOutputReference {
	_init_.Initialize()

	if err := validateNewPostgresSyncedTableSpecNewPipelineSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresSyncedTable.PostgresSyncedTableSpecNewPipelineSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPostgresSyncedTableSpecNewPipelineSpecOutputReference_Override(p PostgresSyncedTableSpecNewPipelineSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresSyncedTable.PostgresSyncedTableSpecNewPipelineSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference)SetBudgetPolicyId(val *string) {
	if err := j.validateSetBudgetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"budgetPolicyId",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference)SetStorageCatalog(val *string) {
	if err := j.validateSetStorageCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCatalog",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference)SetStorageSchema(val *string) {
	if err := j.validateSetStorageSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSchema",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) ResetBudgetPolicyId() {
	_jsii_.InvokeVoid(
		p,
		"resetBudgetPolicyId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) ResetStorageCatalog() {
	_jsii_.InvokeVoid(
		p,
		"resetStorageCatalog",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) ResetStorageSchema() {
	_jsii_.InvokeVoid(
		p,
		"resetStorageSchema",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PostgresSyncedTableSpecNewPipelineSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

