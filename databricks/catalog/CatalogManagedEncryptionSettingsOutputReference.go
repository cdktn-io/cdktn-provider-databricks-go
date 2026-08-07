// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/catalog/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CatalogManagedEncryptionSettingsOutputReference interface {
	cdktn.ComplexObject
	AzureEncryptionSettings() CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference
	AzureEncryptionSettingsInput() *CatalogManagedEncryptionSettingsAzureEncryptionSettings
	AzureKeyVaultKeyId() *string
	SetAzureKeyVaultKeyId(val *string)
	AzureKeyVaultKeyIdInput() *string
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
	CustomerManagedKeyId() *string
	SetCustomerManagedKeyId(val *string)
	CustomerManagedKeyIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CatalogManagedEncryptionSettings
	SetInternalValue(val *CatalogManagedEncryptionSettings)
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
	PutAzureEncryptionSettings(value *CatalogManagedEncryptionSettingsAzureEncryptionSettings)
	ResetAzureEncryptionSettings()
	ResetAzureKeyVaultKeyId()
	ResetCustomerManagedKeyId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CatalogManagedEncryptionSettingsOutputReference
type jsiiProxy_CatalogManagedEncryptionSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) AzureEncryptionSettings() CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference {
	var returns CatalogManagedEncryptionSettingsAzureEncryptionSettingsOutputReference
	_jsii_.Get(
		j,
		"azureEncryptionSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) AzureEncryptionSettingsInput() *CatalogManagedEncryptionSettingsAzureEncryptionSettings {
	var returns *CatalogManagedEncryptionSettingsAzureEncryptionSettings
	_jsii_.Get(
		j,
		"azureEncryptionSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) AzureKeyVaultKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureKeyVaultKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) AzureKeyVaultKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureKeyVaultKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) CustomerManagedKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerManagedKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) CustomerManagedKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerManagedKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) InternalValue() *CatalogManagedEncryptionSettings {
	var returns *CatalogManagedEncryptionSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCatalogManagedEncryptionSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CatalogManagedEncryptionSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewCatalogManagedEncryptionSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CatalogManagedEncryptionSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.catalog.CatalogManagedEncryptionSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCatalogManagedEncryptionSettingsOutputReference_Override(c CatalogManagedEncryptionSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.catalog.CatalogManagedEncryptionSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference)SetAzureKeyVaultKeyId(val *string) {
	if err := j.validateSetAzureKeyVaultKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureKeyVaultKeyId",
		val,
	)
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference)SetCustomerManagedKeyId(val *string) {
	if err := j.validateSetCustomerManagedKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerManagedKeyId",
		val,
	)
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference)SetInternalValue(val *CatalogManagedEncryptionSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) PutAzureEncryptionSettings(value *CatalogManagedEncryptionSettingsAzureEncryptionSettings) {
	if err := c.validatePutAzureEncryptionSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAzureEncryptionSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) ResetAzureEncryptionSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureEncryptionSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) ResetAzureKeyVaultKeyId() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureKeyVaultKeyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) ResetCustomerManagedKeyId() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomerManagedKeyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CatalogManagedEncryptionSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

