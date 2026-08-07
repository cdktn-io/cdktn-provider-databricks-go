// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package modelserving

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/modelserving/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference interface {
	cdktn.ComplexObject
	ApiKeyAuth() ModelServingConfigServedEntitiesExternalModelCustomProviderConfigApiKeyAuthOutputReference
	ApiKeyAuthInput() *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigApiKeyAuth
	BearerTokenAuth() ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference
	BearerTokenAuthInput() *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuth
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
	CustomProviderUrl() *string
	SetCustomProviderUrl(val *string)
	CustomProviderUrlInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ModelServingConfigServedEntitiesExternalModelCustomProviderConfig
	SetInternalValue(val *ModelServingConfigServedEntitiesExternalModelCustomProviderConfig)
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
	PutApiKeyAuth(value *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigApiKeyAuth)
	PutBearerTokenAuth(value *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuth)
	ResetApiKeyAuth()
	ResetBearerTokenAuth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference
type jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) ApiKeyAuth() ModelServingConfigServedEntitiesExternalModelCustomProviderConfigApiKeyAuthOutputReference {
	var returns ModelServingConfigServedEntitiesExternalModelCustomProviderConfigApiKeyAuthOutputReference
	_jsii_.Get(
		j,
		"apiKeyAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) ApiKeyAuthInput() *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigApiKeyAuth {
	var returns *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigApiKeyAuth
	_jsii_.Get(
		j,
		"apiKeyAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) BearerTokenAuth() ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference {
	var returns ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuthOutputReference
	_jsii_.Get(
		j,
		"bearerTokenAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) BearerTokenAuthInput() *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuth {
	var returns *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuth
	_jsii_.Get(
		j,
		"bearerTokenAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) CustomProviderUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customProviderUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) CustomProviderUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customProviderUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) InternalValue() *ModelServingConfigServedEntitiesExternalModelCustomProviderConfig {
	var returns *ModelServingConfigServedEntitiesExternalModelCustomProviderConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference {
	_init_.Initialize()

	if err := validateNewModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference_Override(m ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference)SetCustomProviderUrl(val *string) {
	if err := j.validateSetCustomProviderUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customProviderUrl",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference)SetInternalValue(val *ModelServingConfigServedEntitiesExternalModelCustomProviderConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) PutApiKeyAuth(value *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigApiKeyAuth) {
	if err := m.validatePutApiKeyAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putApiKeyAuth",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) PutBearerTokenAuth(value *ModelServingConfigServedEntitiesExternalModelCustomProviderConfigBearerTokenAuth) {
	if err := m.validatePutBearerTokenAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBearerTokenAuth",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) ResetApiKeyAuth() {
	_jsii_.InvokeVoid(
		m,
		"resetApiKeyAuth",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) ResetBearerTokenAuth() {
	_jsii_.InvokeVoid(
		m,
		"resetBearerTokenAuth",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

