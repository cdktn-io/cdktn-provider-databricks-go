// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v16/datadatabricksapp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksAppAppPendingDeploymentOutputReference interface {
	cdktn.ComplexObject
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	CreateTime() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Creator() *string
	DeploymentArtifacts() DataDatabricksAppAppPendingDeploymentDeploymentArtifactsOutputReference
	DeploymentId() *string
	SetDeploymentId(val *string)
	DeploymentIdInput() *string
	EnvVars() DataDatabricksAppAppPendingDeploymentEnvVarsList
	EnvVarsInput() interface{}
	// Experimental.
	Fqn() *string
	GitSource() DataDatabricksAppAppPendingDeploymentGitSourceOutputReference
	GitSourceInput() interface{}
	InternalValue() *DataDatabricksAppAppPendingDeployment
	SetInternalValue(val *DataDatabricksAppAppPendingDeployment)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	SourceCodePath() *string
	SetSourceCodePath(val *string)
	SourceCodePathInput() *string
	Status() DataDatabricksAppAppPendingDeploymentStatusOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UpdateTime() *string
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
	PutEnvVars(value interface{})
	PutGitSource(value *DataDatabricksAppAppPendingDeploymentGitSource)
	ResetCommand()
	ResetDeploymentId()
	ResetEnvVars()
	ResetGitSource()
	ResetMode()
	ResetSourceCodePath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAppAppPendingDeploymentOutputReference
type jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) Creator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) DeploymentArtifacts() DataDatabricksAppAppPendingDeploymentDeploymentArtifactsOutputReference {
	var returns DataDatabricksAppAppPendingDeploymentDeploymentArtifactsOutputReference
	_jsii_.Get(
		j,
		"deploymentArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) DeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) DeploymentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) EnvVars() DataDatabricksAppAppPendingDeploymentEnvVarsList {
	var returns DataDatabricksAppAppPendingDeploymentEnvVarsList
	_jsii_.Get(
		j,
		"envVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) EnvVarsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) GitSource() DataDatabricksAppAppPendingDeploymentGitSourceOutputReference {
	var returns DataDatabricksAppAppPendingDeploymentGitSourceOutputReference
	_jsii_.Get(
		j,
		"gitSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) GitSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gitSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) InternalValue() *DataDatabricksAppAppPendingDeployment {
	var returns *DataDatabricksAppAppPendingDeployment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) SourceCodePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) SourceCodePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) Status() DataDatabricksAppAppPendingDeploymentStatusOutputReference {
	var returns DataDatabricksAppAppPendingDeploymentStatusOutputReference
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


func NewDataDatabricksAppAppPendingDeploymentOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksAppAppPendingDeploymentOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAppAppPendingDeploymentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksApp.DataDatabricksAppAppPendingDeploymentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAppAppPendingDeploymentOutputReference_Override(d DataDatabricksAppAppPendingDeploymentOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksApp.DataDatabricksAppAppPendingDeploymentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference)SetDeploymentId(val *string) {
	if err := j.validateSetDeploymentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference)SetInternalValue(val *DataDatabricksAppAppPendingDeployment) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference)SetSourceCodePath(val *string) {
	if err := j.validateSetSourceCodePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCodePath",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) PutEnvVars(value interface{}) {
	if err := d.validatePutEnvVarsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEnvVars",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) PutGitSource(value *DataDatabricksAppAppPendingDeploymentGitSource) {
	if err := d.validatePutGitSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGitSource",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		d,
		"resetCommand",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) ResetDeploymentId() {
	_jsii_.InvokeVoid(
		d,
		"resetDeploymentId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) ResetEnvVars() {
	_jsii_.InvokeVoid(
		d,
		"resetEnvVars",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) ResetGitSource() {
	_jsii_.InvokeVoid(
		d,
		"resetGitSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		d,
		"resetMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) ResetSourceCodePath() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceCodePath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAppAppPendingDeploymentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

