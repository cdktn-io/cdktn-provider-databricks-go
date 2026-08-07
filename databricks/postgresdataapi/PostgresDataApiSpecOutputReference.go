// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresdataapi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/postgresdataapi/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PostgresDataApiSpecOutputReference interface {
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
	DbAggregatesEnabled() interface{}
	SetDbAggregatesEnabled(val interface{})
	DbAggregatesEnabledInput() interface{}
	DbExtraSearchPath() *[]*string
	SetDbExtraSearchPath(val *[]*string)
	DbExtraSearchPathInput() *[]*string
	DbMaxRows() *float64
	SetDbMaxRows(val *float64)
	DbMaxRowsInput() *float64
	DbSchemas() *[]*string
	SetDbSchemas(val *[]*string)
	DbSchemasInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JwtCacheMaxLifetime() *string
	SetJwtCacheMaxLifetime(val *string)
	JwtCacheMaxLifetimeInput() *string
	JwtRoleClaimKey() *string
	SetJwtRoleClaimKey(val *string)
	JwtRoleClaimKeyInput() *string
	OpenapiMode() *string
	SetOpenapiMode(val *string)
	OpenapiModeInput() *string
	ServerCorsAllowedOrigins() *[]*string
	SetServerCorsAllowedOrigins(val *[]*string)
	ServerCorsAllowedOriginsInput() *[]*string
	ServerTimingEnabled() interface{}
	SetServerTimingEnabled(val interface{})
	ServerTimingEnabledInput() interface{}
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
	ResetDbAggregatesEnabled()
	ResetDbExtraSearchPath()
	ResetDbMaxRows()
	ResetDbSchemas()
	ResetJwtCacheMaxLifetime()
	ResetJwtRoleClaimKey()
	ResetOpenapiMode()
	ResetServerCorsAllowedOrigins()
	ResetServerTimingEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PostgresDataApiSpecOutputReference
type jsiiProxy_PostgresDataApiSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) DbAggregatesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dbAggregatesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) DbAggregatesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dbAggregatesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) DbExtraSearchPath() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbExtraSearchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) DbExtraSearchPathInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbExtraSearchPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) DbMaxRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbMaxRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) DbMaxRowsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbMaxRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) DbSchemas() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbSchemas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) DbSchemasInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbSchemasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) JwtCacheMaxLifetime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtCacheMaxLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) JwtCacheMaxLifetimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtCacheMaxLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) JwtRoleClaimKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtRoleClaimKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) JwtRoleClaimKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtRoleClaimKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) OpenapiMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openapiMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) OpenapiModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openapiModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) ServerCorsAllowedOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverCorsAllowedOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) ServerCorsAllowedOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverCorsAllowedOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) ServerTimingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverTimingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) ServerTimingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverTimingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPostgresDataApiSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PostgresDataApiSpecOutputReference {
	_init_.Initialize()

	if err := validateNewPostgresDataApiSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PostgresDataApiSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresDataApi.PostgresDataApiSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPostgresDataApiSpecOutputReference_Override(p PostgresDataApiSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.postgresDataApi.PostgresDataApiSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference)SetDbAggregatesEnabled(val interface{}) {
	if err := j.validateSetDbAggregatesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbAggregatesEnabled",
		val,
	)
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference)SetDbExtraSearchPath(val *[]*string) {
	if err := j.validateSetDbExtraSearchPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbExtraSearchPath",
		val,
	)
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference)SetDbMaxRows(val *float64) {
	if err := j.validateSetDbMaxRowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbMaxRows",
		val,
	)
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference)SetDbSchemas(val *[]*string) {
	if err := j.validateSetDbSchemasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSchemas",
		val,
	)
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference)SetJwtCacheMaxLifetime(val *string) {
	if err := j.validateSetJwtCacheMaxLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtCacheMaxLifetime",
		val,
	)
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference)SetJwtRoleClaimKey(val *string) {
	if err := j.validateSetJwtRoleClaimKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtRoleClaimKey",
		val,
	)
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference)SetOpenapiMode(val *string) {
	if err := j.validateSetOpenapiModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openapiMode",
		val,
	)
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference)SetServerCorsAllowedOrigins(val *[]*string) {
	if err := j.validateSetServerCorsAllowedOriginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCorsAllowedOrigins",
		val,
	)
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference)SetServerTimingEnabled(val interface{}) {
	if err := j.validateSetServerTimingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverTimingEnabled",
		val,
	)
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PostgresDataApiSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) ResetDbAggregatesEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetDbAggregatesEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) ResetDbExtraSearchPath() {
	_jsii_.InvokeVoid(
		p,
		"resetDbExtraSearchPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) ResetDbMaxRows() {
	_jsii_.InvokeVoid(
		p,
		"resetDbMaxRows",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) ResetDbSchemas() {
	_jsii_.InvokeVoid(
		p,
		"resetDbSchemas",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) ResetJwtCacheMaxLifetime() {
	_jsii_.InvokeVoid(
		p,
		"resetJwtCacheMaxLifetime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) ResetJwtRoleClaimKey() {
	_jsii_.InvokeVoid(
		p,
		"resetJwtRoleClaimKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) ResetOpenapiMode() {
	_jsii_.InvokeVoid(
		p,
		"resetOpenapiMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) ResetServerCorsAllowedOrigins() {
	_jsii_.InvokeVoid(
		p,
		"resetServerCorsAllowedOrigins",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) ResetServerTimingEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetServerTimingEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PostgresDataApiSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

