// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v17/catalog/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference interface {
	cdktn.ComplexObject
	AzureCmkAccessConnectorId() *string
	SetAzureCmkAccessConnectorId(val *string)
	AzureCmkAccessConnectorIdInput() *string
	AzureCmkManagedIdentityId() *string
	SetAzureCmkManagedIdentityId(val *string)
	AzureCmkManagedIdentityIdInput() *string
	AzureTenantId() *string
	SetAzureTenantId(val *string)
	AzureTenantIdInput() *string
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
	InternalValue() *CatalogManagedEncryptionSettingsAzureEncryptionSettings
	SetInternalValue(val *CatalogManagedEncryptionSettingsAzureEncryptionSettings)
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
	ResetAzureCmkAccessConnectorId()
	ResetAzureCmkManagedIdentityId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference
type jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) AzureCmkAccessConnectorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureCmkAccessConnectorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) AzureCmkAccessConnectorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureCmkAccessConnectorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) AzureCmkManagedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureCmkManagedIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) AzureCmkManagedIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureCmkManagedIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) AzureTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) AzureTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) InternalValue() *CatalogManagedEncryptionSettingsAzureEncryptionSettings {
	var returns *CatalogManagedEncryptionSettingsAzureEncryptionSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewCatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.catalog.CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference_Override(c CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.catalog.CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference)SetAzureCmkAccessConnectorId(val *string) {
	if err := j.validateSetAzureCmkAccessConnectorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureCmkAccessConnectorId",
		val,
	)
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference)SetAzureCmkManagedIdentityId(val *string) {
	if err := j.validateSetAzureCmkManagedIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureCmkManagedIdentityId",
		val,
	)
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference)SetAzureTenantId(val *string) {
	if err := j.validateSetAzureTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureTenantId",
		val,
	)
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference)SetInternalValue(val *CatalogManagedEncryptionSettingsAzureEncryptionSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) ResetAzureCmkAccessConnectorId() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureCmkAccessConnectorId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) ResetAzureCmkManagedIdentityId() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureCmkManagedIdentityId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

