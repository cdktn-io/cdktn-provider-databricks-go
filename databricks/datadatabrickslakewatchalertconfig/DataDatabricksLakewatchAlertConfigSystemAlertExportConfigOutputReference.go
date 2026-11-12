// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickslakewatchalertconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/datadatabrickslakewatchalertconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksLakewatchAlertConfigSystemAlertExportConfig
	SetInternalValue(val *DataDatabricksLakewatchAlertConfigSystemAlertExportConfig)
	Jira() DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference
	JiraInput() interface{}
	ServiceNow() DataDatabricksLakewatchAlertConfigSystemAlertExportConfigServiceNowOutputReference
	ServiceNowInput() interface{}
	Slack() DataDatabricksLakewatchAlertConfigSystemAlertExportConfigSlackOutputReference
	SlackInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Tines() DataDatabricksLakewatchAlertConfigSystemAlertExportConfigTinesOutputReference
	TinesInput() interface{}
	Webhook() DataDatabricksLakewatchAlertConfigSystemAlertExportConfigWebhookOutputReference
	WebhookInput() interface{}
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	WorkspaceIdInput() *string
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
	PutJira(value *DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJira)
	PutServiceNow(value *DataDatabricksLakewatchAlertConfigSystemAlertExportConfigServiceNow)
	PutSlack(value *DataDatabricksLakewatchAlertConfigSystemAlertExportConfigSlack)
	PutTines(value *DataDatabricksLakewatchAlertConfigSystemAlertExportConfigTines)
	PutWebhook(value *DataDatabricksLakewatchAlertConfigSystemAlertExportConfigWebhook)
	ResetEnabled()
	ResetJira()
	ResetServiceNow()
	ResetSlack()
	ResetTines()
	ResetWebhook()
	ResetWorkspaceId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference
type jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) InternalValue() *DataDatabricksLakewatchAlertConfigSystemAlertExportConfig {
	var returns *DataDatabricksLakewatchAlertConfigSystemAlertExportConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) Jira() DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference {
	var returns DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJiraOutputReference
	_jsii_.Get(
		j,
		"jira",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) JiraInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jiraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) ServiceNow() DataDatabricksLakewatchAlertConfigSystemAlertExportConfigServiceNowOutputReference {
	var returns DataDatabricksLakewatchAlertConfigSystemAlertExportConfigServiceNowOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) ServiceNowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) Slack() DataDatabricksLakewatchAlertConfigSystemAlertExportConfigSlackOutputReference {
	var returns DataDatabricksLakewatchAlertConfigSystemAlertExportConfigSlackOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) SlackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) Tines() DataDatabricksLakewatchAlertConfigSystemAlertExportConfigTinesOutputReference {
	var returns DataDatabricksLakewatchAlertConfigSystemAlertExportConfigTinesOutputReference
	_jsii_.Get(
		j,
		"tines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) TinesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) Webhook() DataDatabricksLakewatchAlertConfigSystemAlertExportConfigWebhookOutputReference {
	var returns DataDatabricksLakewatchAlertConfigSystemAlertExportConfigWebhookOutputReference
	_jsii_.Get(
		j,
		"webhook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) WebhookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webhookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchAlertConfig.DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference_Override(d DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksLakewatchAlertConfig.DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference)SetInternalValue(val *DataDatabricksLakewatchAlertConfigSystemAlertExportConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) PutJira(value *DataDatabricksLakewatchAlertConfigSystemAlertExportConfigJira) {
	if err := d.validatePutJiraParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJira",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) PutServiceNow(value *DataDatabricksLakewatchAlertConfigSystemAlertExportConfigServiceNow) {
	if err := d.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) PutSlack(value *DataDatabricksLakewatchAlertConfigSystemAlertExportConfigSlack) {
	if err := d.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSlack",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) PutTines(value *DataDatabricksLakewatchAlertConfigSystemAlertExportConfigTines) {
	if err := d.validatePutTinesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTines",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) PutWebhook(value *DataDatabricksLakewatchAlertConfigSystemAlertExportConfigWebhook) {
	if err := d.validatePutWebhookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWebhook",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) ResetJira() {
	_jsii_.InvokeVoid(
		d,
		"resetJira",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		d,
		"resetSlack",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) ResetTines() {
	_jsii_.InvokeVoid(
		d,
		"resetTines",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) ResetWebhook() {
	_jsii_.InvokeVoid(
		d,
		"resetWebhook",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) ResetWorkspaceId() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkspaceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksLakewatchAlertConfigSystemAlertExportConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

