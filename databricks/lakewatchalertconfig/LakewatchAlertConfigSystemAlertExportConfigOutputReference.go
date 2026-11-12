// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchalertconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/lakewatchalertconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LakewatchAlertConfigSystemAlertExportConfigOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Jira() LakewatchAlertConfigSystemAlertExportConfigJiraOutputReference
	JiraInput() interface{}
	ServiceNow() LakewatchAlertConfigSystemAlertExportConfigServiceNowOutputReference
	ServiceNowInput() interface{}
	Slack() LakewatchAlertConfigSystemAlertExportConfigSlackOutputReference
	SlackInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Tines() LakewatchAlertConfigSystemAlertExportConfigTinesOutputReference
	TinesInput() interface{}
	Webhook() LakewatchAlertConfigSystemAlertExportConfigWebhookOutputReference
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
	PutJira(value *LakewatchAlertConfigSystemAlertExportConfigJira)
	PutServiceNow(value *LakewatchAlertConfigSystemAlertExportConfigServiceNow)
	PutSlack(value *LakewatchAlertConfigSystemAlertExportConfigSlack)
	PutTines(value *LakewatchAlertConfigSystemAlertExportConfigTines)
	PutWebhook(value *LakewatchAlertConfigSystemAlertExportConfigWebhook)
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

// The jsii proxy struct for LakewatchAlertConfigSystemAlertExportConfigOutputReference
type jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) Jira() LakewatchAlertConfigSystemAlertExportConfigJiraOutputReference {
	var returns LakewatchAlertConfigSystemAlertExportConfigJiraOutputReference
	_jsii_.Get(
		j,
		"jira",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) JiraInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jiraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) ServiceNow() LakewatchAlertConfigSystemAlertExportConfigServiceNowOutputReference {
	var returns LakewatchAlertConfigSystemAlertExportConfigServiceNowOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) ServiceNowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) Slack() LakewatchAlertConfigSystemAlertExportConfigSlackOutputReference {
	var returns LakewatchAlertConfigSystemAlertExportConfigSlackOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) SlackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) Tines() LakewatchAlertConfigSystemAlertExportConfigTinesOutputReference {
	var returns LakewatchAlertConfigSystemAlertExportConfigTinesOutputReference
	_jsii_.Get(
		j,
		"tines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) TinesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) Webhook() LakewatchAlertConfigSystemAlertExportConfigWebhookOutputReference {
	var returns LakewatchAlertConfigSystemAlertExportConfigWebhookOutputReference
	_jsii_.Get(
		j,
		"webhook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) WebhookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webhookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}


func NewLakewatchAlertConfigSystemAlertExportConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LakewatchAlertConfigSystemAlertExportConfigOutputReference {
	_init_.Initialize()

	if err := validateNewLakewatchAlertConfigSystemAlertExportConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.lakewatchAlertConfig.LakewatchAlertConfigSystemAlertExportConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLakewatchAlertConfigSystemAlertExportConfigOutputReference_Override(l LakewatchAlertConfigSystemAlertExportConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.lakewatchAlertConfig.LakewatchAlertConfigSystemAlertExportConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) PutJira(value *LakewatchAlertConfigSystemAlertExportConfigJira) {
	if err := l.validatePutJiraParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putJira",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) PutServiceNow(value *LakewatchAlertConfigSystemAlertExportConfigServiceNow) {
	if err := l.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) PutSlack(value *LakewatchAlertConfigSystemAlertExportConfigSlack) {
	if err := l.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSlack",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) PutTines(value *LakewatchAlertConfigSystemAlertExportConfigTines) {
	if err := l.validatePutTinesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTines",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) PutWebhook(value *LakewatchAlertConfigSystemAlertExportConfigWebhook) {
	if err := l.validatePutWebhookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putWebhook",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) ResetJira() {
	_jsii_.InvokeVoid(
		l,
		"resetJira",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		l,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		l,
		"resetSlack",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) ResetTines() {
	_jsii_.InvokeVoid(
		l,
		"resetTines",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) ResetWebhook() {
	_jsii_.InvokeVoid(
		l,
		"resetWebhook",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) ResetWorkspaceId() {
	_jsii_.InvokeVoid(
		l,
		"resetWorkspaceId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchAlertConfigSystemAlertExportConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

