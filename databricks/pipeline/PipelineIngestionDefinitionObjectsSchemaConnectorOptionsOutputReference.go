// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v18/pipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference interface {
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
	ConfluenceOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsConfluenceOptionsOutputReference
	ConfluenceOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsConfluenceOptions
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GdriveOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsOutputReference
	GdriveOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptions
	GoogleAdsOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference
	GoogleAdsOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptions
	InternalValue() *PipelineIngestionDefinitionObjectsSchemaConnectorOptions
	SetInternalValue(val *PipelineIngestionDefinitionObjectsSchemaConnectorOptions)
	JiraOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsJiraOptionsOutputReference
	JiraOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsJiraOptions
	KafkaOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsOutputReference
	KafkaOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptions
	MetaAdsOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsMetaAdsOptionsOutputReference
	MetaAdsOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsMetaAdsOptions
	OutlookOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference
	OutlookOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptions
	RedditAdsOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsRedditAdsOptionsOutputReference
	RedditAdsOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsRedditAdsOptions
	SharepointOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSharepointOptionsOutputReference
	SharepointOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSharepointOptions
	SmartsheetOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSmartsheetOptionsOutputReference
	SmartsheetOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSmartsheetOptions
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TiktokAdsOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsTiktokAdsOptionsOutputReference
	TiktokAdsOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsTiktokAdsOptions
	ZendeskSupportOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsZendeskSupportOptionsOutputReference
	ZendeskSupportOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsZendeskSupportOptions
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
	PutConfluenceOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsConfluenceOptions)
	PutGdriveOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptions)
	PutGoogleAdsOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptions)
	PutJiraOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsJiraOptions)
	PutKafkaOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptions)
	PutMetaAdsOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsMetaAdsOptions)
	PutOutlookOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptions)
	PutRedditAdsOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsRedditAdsOptions)
	PutSharepointOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSharepointOptions)
	PutSmartsheetOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSmartsheetOptions)
	PutTiktokAdsOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsTiktokAdsOptions)
	PutZendeskSupportOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsZendeskSupportOptions)
	ResetConfluenceOptions()
	ResetGdriveOptions()
	ResetGoogleAdsOptions()
	ResetJiraOptions()
	ResetKafkaOptions()
	ResetMetaAdsOptions()
	ResetOutlookOptions()
	ResetRedditAdsOptions()
	ResetSharepointOptions()
	ResetSmartsheetOptions()
	ResetTiktokAdsOptions()
	ResetZendeskSupportOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference
type jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ConfluenceOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsConfluenceOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsSchemaConnectorOptionsConfluenceOptionsOutputReference
	_jsii_.Get(
		j,
		"confluenceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ConfluenceOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsConfluenceOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsConfluenceOptions
	_jsii_.Get(
		j,
		"confluenceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) GdriveOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptionsOutputReference
	_jsii_.Get(
		j,
		"gdriveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) GdriveOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptions
	_jsii_.Get(
		j,
		"gdriveOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) GoogleAdsOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptionsOutputReference
	_jsii_.Get(
		j,
		"googleAdsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) GoogleAdsOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptions
	_jsii_.Get(
		j,
		"googleAdsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) InternalValue() *PipelineIngestionDefinitionObjectsSchemaConnectorOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) JiraOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsJiraOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsSchemaConnectorOptionsJiraOptionsOutputReference
	_jsii_.Get(
		j,
		"jiraOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) JiraOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsJiraOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsJiraOptions
	_jsii_.Get(
		j,
		"jiraOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) KafkaOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptionsOutputReference
	_jsii_.Get(
		j,
		"kafkaOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) KafkaOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptions
	_jsii_.Get(
		j,
		"kafkaOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) MetaAdsOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsMetaAdsOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsSchemaConnectorOptionsMetaAdsOptionsOutputReference
	_jsii_.Get(
		j,
		"metaAdsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) MetaAdsOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsMetaAdsOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsMetaAdsOptions
	_jsii_.Get(
		j,
		"metaAdsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) OutlookOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptionsOutputReference
	_jsii_.Get(
		j,
		"outlookOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) OutlookOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptions
	_jsii_.Get(
		j,
		"outlookOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) RedditAdsOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsRedditAdsOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsSchemaConnectorOptionsRedditAdsOptionsOutputReference
	_jsii_.Get(
		j,
		"redditAdsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) RedditAdsOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsRedditAdsOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsRedditAdsOptions
	_jsii_.Get(
		j,
		"redditAdsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) SharepointOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSharepointOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSharepointOptionsOutputReference
	_jsii_.Get(
		j,
		"sharepointOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) SharepointOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSharepointOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSharepointOptions
	_jsii_.Get(
		j,
		"sharepointOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) SmartsheetOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSmartsheetOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSmartsheetOptionsOutputReference
	_jsii_.Get(
		j,
		"smartsheetOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) SmartsheetOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSmartsheetOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSmartsheetOptions
	_jsii_.Get(
		j,
		"smartsheetOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) TiktokAdsOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsTiktokAdsOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsSchemaConnectorOptionsTiktokAdsOptionsOutputReference
	_jsii_.Get(
		j,
		"tiktokAdsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) TiktokAdsOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsTiktokAdsOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsTiktokAdsOptions
	_jsii_.Get(
		j,
		"tiktokAdsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ZendeskSupportOptions() PipelineIngestionDefinitionObjectsSchemaConnectorOptionsZendeskSupportOptionsOutputReference {
	var returns PipelineIngestionDefinitionObjectsSchemaConnectorOptionsZendeskSupportOptionsOutputReference
	_jsii_.Get(
		j,
		"zendeskSupportOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ZendeskSupportOptionsInput() *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsZendeskSupportOptions {
	var returns *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsZendeskSupportOptions
	_jsii_.Get(
		j,
		"zendeskSupportOptionsInput",
		&returns,
	)
	return returns
}


func NewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference_Override(p PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.pipeline.PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference)SetInternalValue(val *PipelineIngestionDefinitionObjectsSchemaConnectorOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) PutConfluenceOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsConfluenceOptions) {
	if err := p.validatePutConfluenceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putConfluenceOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) PutGdriveOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGdriveOptions) {
	if err := p.validatePutGdriveOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGdriveOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) PutGoogleAdsOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsGoogleAdsOptions) {
	if err := p.validatePutGoogleAdsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGoogleAdsOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) PutJiraOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsJiraOptions) {
	if err := p.validatePutJiraOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putJiraOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) PutKafkaOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsKafkaOptions) {
	if err := p.validatePutKafkaOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putKafkaOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) PutMetaAdsOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsMetaAdsOptions) {
	if err := p.validatePutMetaAdsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMetaAdsOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) PutOutlookOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutlookOptions) {
	if err := p.validatePutOutlookOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putOutlookOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) PutRedditAdsOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsRedditAdsOptions) {
	if err := p.validatePutRedditAdsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRedditAdsOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) PutSharepointOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSharepointOptions) {
	if err := p.validatePutSharepointOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSharepointOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) PutSmartsheetOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsSmartsheetOptions) {
	if err := p.validatePutSmartsheetOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSmartsheetOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) PutTiktokAdsOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsTiktokAdsOptions) {
	if err := p.validatePutTiktokAdsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTiktokAdsOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) PutZendeskSupportOptions(value *PipelineIngestionDefinitionObjectsSchemaConnectorOptionsZendeskSupportOptions) {
	if err := p.validatePutZendeskSupportOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putZendeskSupportOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ResetConfluenceOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetConfluenceOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ResetGdriveOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetGdriveOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ResetGoogleAdsOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetGoogleAdsOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ResetJiraOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetJiraOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ResetKafkaOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetKafkaOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ResetMetaAdsOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetMetaAdsOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ResetOutlookOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetOutlookOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ResetRedditAdsOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetRedditAdsOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ResetSharepointOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetSharepointOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ResetSmartsheetOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetSmartsheetOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ResetTiktokAdsOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetTiktokAdsOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ResetZendeskSupportOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetZendeskSupportOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsSchemaConnectorOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

