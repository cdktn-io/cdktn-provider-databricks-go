// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgressyncedtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/postgressyncedtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PostgresSyncedTableSpecOutputReference interface {
	cdktn.ComplexObject
	Branch() *string
	SetBranch(val *string)
	BranchInput() *string
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
	CreateDatabaseObjectsIfMissing() interface{}
	SetCreateDatabaseObjectsIfMissing(val interface{})
	CreateDatabaseObjectsIfMissingInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExistingPipelineId() *string
	SetExistingPipelineId(val *string)
	ExistingPipelineIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NewPipelineSpec() PostgresSyncedTableSpecNewPipelineSpecOutputReference
	NewPipelineSpecInput() interface{}
	PostgresDatabase() *string
	SetPostgresDatabase(val *string)
	PostgresDatabaseInput() *string
	PrimaryKeyColumns() *[]*string
	SetPrimaryKeyColumns(val *[]*string)
	PrimaryKeyColumnsInput() *[]*string
	SchedulingPolicy() *string
	SetSchedulingPolicy(val *string)
	SchedulingPolicyInput() *string
	SourceTableFullName() *string
	SetSourceTableFullName(val *string)
	SourceTableFullNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeseriesKey() *string
	SetTimeseriesKey(val *string)
	TimeseriesKeyInput() *string
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
	PutNewPipelineSpec(value *PostgresSyncedTableSpecNewPipelineSpec)
	ResetBranch()
	ResetCreateDatabaseObjectsIfMissing()
	ResetExistingPipelineId()
	ResetNewPipelineSpec()
	ResetPostgresDatabase()
	ResetPrimaryKeyColumns()
	ResetSchedulingPolicy()
	ResetSourceTableFullName()
	ResetTimeseriesKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PostgresSyncedTableSpecOutputReference
type jsiiProxy_PostgresSyncedTableSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) Branch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) BranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) CreateDatabaseObjectsIfMissing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatabaseObjectsIfMissing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) CreateDatabaseObjectsIfMissingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatabaseObjectsIfMissingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) ExistingPipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingPipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) ExistingPipelineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingPipelineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) NewPipelineSpec() PostgresSyncedTableSpecNewPipelineSpecOutputReference {
	var returns PostgresSyncedTableSpecNewPipelineSpecOutputReference
	_jsii_.Get(
		j,
		"newPipelineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) NewPipelineSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newPipelineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) PostgresDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) PostgresDatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) PrimaryKeyColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"primaryKeyColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) PrimaryKeyColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"primaryKeyColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) SchedulingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) SchedulingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) SourceTableFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTableFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) SourceTableFullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTableFullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) TimeseriesKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeseriesKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference) TimeseriesKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeseriesKeyInput",
		&returns,
	)
	return returns
}


func NewPostgresSyncedTableSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PostgresSyncedTableSpecOutputReference {
	_init_.Initialize()

	if err := validateNewPostgresSyncedTableSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PostgresSyncedTableSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresSyncedTable.PostgresSyncedTableSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPostgresSyncedTableSpecOutputReference_Override(p PostgresSyncedTableSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresSyncedTable.PostgresSyncedTableSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference)SetBranch(val *string) {
	if err := j.validateSetBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branch",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference)SetCreateDatabaseObjectsIfMissing(val interface{}) {
	if err := j.validateSetCreateDatabaseObjectsIfMissingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDatabaseObjectsIfMissing",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference)SetExistingPipelineId(val *string) {
	if err := j.validateSetExistingPipelineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"existingPipelineId",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference)SetPostgresDatabase(val *string) {
	if err := j.validateSetPostgresDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postgresDatabase",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference)SetPrimaryKeyColumns(val *[]*string) {
	if err := j.validateSetPrimaryKeyColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryKeyColumns",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference)SetSchedulingPolicy(val *string) {
	if err := j.validateSetSchedulingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulingPolicy",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference)SetSourceTableFullName(val *string) {
	if err := j.validateSetSourceTableFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTableFullName",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PostgresSyncedTableSpecOutputReference)SetTimeseriesKey(val *string) {
	if err := j.validateSetTimeseriesKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeseriesKey",
		val,
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) PutNewPipelineSpec(value *PostgresSyncedTableSpecNewPipelineSpec) {
	if err := p.validatePutNewPipelineSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNewPipelineSpec",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) ResetBranch() {
	_jsii_.InvokeVoid(
		p,
		"resetBranch",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) ResetCreateDatabaseObjectsIfMissing() {
	_jsii_.InvokeVoid(
		p,
		"resetCreateDatabaseObjectsIfMissing",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) ResetExistingPipelineId() {
	_jsii_.InvokeVoid(
		p,
		"resetExistingPipelineId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) ResetNewPipelineSpec() {
	_jsii_.InvokeVoid(
		p,
		"resetNewPipelineSpec",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) ResetPostgresDatabase() {
	_jsii_.InvokeVoid(
		p,
		"resetPostgresDatabase",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) ResetPrimaryKeyColumns() {
	_jsii_.InvokeVoid(
		p,
		"resetPrimaryKeyColumns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) ResetSchedulingPolicy() {
	_jsii_.InvokeVoid(
		p,
		"resetSchedulingPolicy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) ResetSourceTableFullName() {
	_jsii_.InvokeVoid(
		p,
		"resetSourceTableFullName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) ResetTimeseriesKey() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeseriesKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PostgresSyncedTableSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

