// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspostgresdataapi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/datadatabrickspostgresdataapi/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatabricksPostgresDataApiSpecOutputReference interface {
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
	InternalValue() *DataDatabricksPostgresDataApiSpec
	SetInternalValue(val *DataDatabricksPostgresDataApiSpec)
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

// The jsii proxy struct for DataDatabricksPostgresDataApiSpecOutputReference
type jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) DbAggregatesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dbAggregatesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) DbAggregatesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dbAggregatesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) DbExtraSearchPath() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbExtraSearchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) DbExtraSearchPathInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbExtraSearchPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) DbMaxRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbMaxRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) DbMaxRowsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbMaxRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) DbSchemas() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbSchemas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) DbSchemasInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbSchemasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) InternalValue() *DataDatabricksPostgresDataApiSpec {
	var returns *DataDatabricksPostgresDataApiSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) JwtCacheMaxLifetime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtCacheMaxLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) JwtCacheMaxLifetimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtCacheMaxLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) JwtRoleClaimKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtRoleClaimKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) JwtRoleClaimKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtRoleClaimKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) OpenapiMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openapiMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) OpenapiModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openapiModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ServerCorsAllowedOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverCorsAllowedOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ServerCorsAllowedOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverCorsAllowedOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ServerTimingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverTimingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ServerTimingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverTimingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksPostgresDataApiSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataDatabricksPostgresDataApiSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksPostgresDataApiSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresDataApi.DataDatabricksPostgresDataApiSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksPostgresDataApiSpecOutputReference_Override(d DataDatabricksPostgresDataApiSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.dataDatabricksPostgresDataApi.DataDatabricksPostgresDataApiSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference)SetDbAggregatesEnabled(val interface{}) {
	if err := j.validateSetDbAggregatesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbAggregatesEnabled",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference)SetDbExtraSearchPath(val *[]*string) {
	if err := j.validateSetDbExtraSearchPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbExtraSearchPath",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference)SetDbMaxRows(val *float64) {
	if err := j.validateSetDbMaxRowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbMaxRows",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference)SetDbSchemas(val *[]*string) {
	if err := j.validateSetDbSchemasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSchemas",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference)SetInternalValue(val *DataDatabricksPostgresDataApiSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference)SetJwtCacheMaxLifetime(val *string) {
	if err := j.validateSetJwtCacheMaxLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtCacheMaxLifetime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference)SetJwtRoleClaimKey(val *string) {
	if err := j.validateSetJwtRoleClaimKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtRoleClaimKey",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference)SetOpenapiMode(val *string) {
	if err := j.validateSetOpenapiModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openapiMode",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference)SetServerCorsAllowedOrigins(val *[]*string) {
	if err := j.validateSetServerCorsAllowedOriginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCorsAllowedOrigins",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference)SetServerTimingEnabled(val interface{}) {
	if err := j.validateSetServerTimingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverTimingEnabled",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ResetDbAggregatesEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDbAggregatesEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ResetDbExtraSearchPath() {
	_jsii_.InvokeVoid(
		d,
		"resetDbExtraSearchPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ResetDbMaxRows() {
	_jsii_.InvokeVoid(
		d,
		"resetDbMaxRows",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ResetDbSchemas() {
	_jsii_.InvokeVoid(
		d,
		"resetDbSchemas",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ResetJwtCacheMaxLifetime() {
	_jsii_.InvokeVoid(
		d,
		"resetJwtCacheMaxLifetime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ResetJwtRoleClaimKey() {
	_jsii_.InvokeVoid(
		d,
		"resetJwtRoleClaimKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ResetOpenapiMode() {
	_jsii_.InvokeVoid(
		d,
		"resetOpenapiMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ResetServerCorsAllowedOrigins() {
	_jsii_.InvokeVoid(
		d,
		"resetServerCorsAllowedOrigins",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ResetServerTimingEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetServerTimingEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksPostgresDataApiSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

