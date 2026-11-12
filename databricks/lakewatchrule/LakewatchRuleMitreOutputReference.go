// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakewatchrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/lakewatchrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LakewatchRuleMitreOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SubTechnique() *string
	SetSubTechnique(val *string)
	SubTechniqueId() *string
	SetSubTechniqueId(val *string)
	SubTechniqueIdInput() *string
	SubTechniqueInput() *string
	Tactic() *string
	SetTactic(val *string)
	TacticInput() *string
	Taxonomy() *string
	SetTaxonomy(val *string)
	TaxonomyInput() *string
	Technique() *string
	SetTechnique(val *string)
	TechniqueId() *string
	SetTechniqueId(val *string)
	TechniqueIdInput() *string
	TechniqueInput() *string
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
	ResetSubTechnique()
	ResetSubTechniqueId()
	ResetTactic()
	ResetTaxonomy()
	ResetTechnique()
	ResetTechniqueId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LakewatchRuleMitreOutputReference
type jsiiProxy_LakewatchRuleMitreOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) SubTechnique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTechnique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) SubTechniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTechniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) SubTechniqueIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTechniqueIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) SubTechniqueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subTechniqueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) Tactic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tactic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) TacticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tacticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) Taxonomy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taxonomy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) TaxonomyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taxonomyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) Technique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) TechniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"techniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) TechniqueIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"techniqueIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) TechniqueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"techniqueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLakewatchRuleMitreOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LakewatchRuleMitreOutputReference {
	_init_.Initialize()

	if err := validateNewLakewatchRuleMitreOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LakewatchRuleMitreOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.lakewatchRule.LakewatchRuleMitreOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLakewatchRuleMitreOutputReference_Override(l LakewatchRuleMitreOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.lakewatchRule.LakewatchRuleMitreOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference)SetSubTechnique(val *string) {
	if err := j.validateSetSubTechniqueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subTechnique",
		val,
	)
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference)SetSubTechniqueId(val *string) {
	if err := j.validateSetSubTechniqueIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subTechniqueId",
		val,
	)
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference)SetTactic(val *string) {
	if err := j.validateSetTacticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tactic",
		val,
	)
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference)SetTaxonomy(val *string) {
	if err := j.validateSetTaxonomyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taxonomy",
		val,
	)
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference)SetTechnique(val *string) {
	if err := j.validateSetTechniqueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technique",
		val,
	)
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference)SetTechniqueId(val *string) {
	if err := j.validateSetTechniqueIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"techniqueId",
		val,
	)
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LakewatchRuleMitreOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) ResetSubTechnique() {
	_jsii_.InvokeVoid(
		l,
		"resetSubTechnique",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) ResetSubTechniqueId() {
	_jsii_.InvokeVoid(
		l,
		"resetSubTechniqueId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) ResetTactic() {
	_jsii_.InvokeVoid(
		l,
		"resetTactic",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) ResetTaxonomy() {
	_jsii_.InvokeVoid(
		l,
		"resetTaxonomy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) ResetTechnique() {
	_jsii_.InvokeVoid(
		l,
		"resetTechnique",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) ResetTechniqueId() {
	_jsii_.InvokeVoid(
		l,
		"resetTechniqueId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (l *jsiiProxy_LakewatchRuleMitreOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

