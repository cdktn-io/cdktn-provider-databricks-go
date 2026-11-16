// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-databricks-go/databricks/v19/workspacesettingv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference interface {
	cdktn.ComplexObject
	AxisLineColor() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeAxisLineColorOutputReference
	AxisLineColorInput() interface{}
	CanvasBackgroundColor() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeCanvasBackgroundColorOutputReference
	CanvasBackgroundColorInput() interface{}
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
	ContinuousGradient() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference
	ContinuousGradientInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FontColor() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontColorOutputReference
	FontColorInput() interface{}
	FontFamily() *string
	SetFontFamily(val *string)
	FontFamilyInput() *string
	FontSettings() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference
	FontSettingsInput() interface{}
	// Experimental.
	Fqn() *string
	GridLineColor() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeGridLineColorOutputReference
	GridLineColorInput() interface{}
	InternalValue() *WorkspaceSettingV2EffectiveWorkspaceDashboardTheme
	SetInternalValue(val *WorkspaceSettingV2EffectiveWorkspaceDashboardTheme)
	MarkCornerRadius() *float64
	SetMarkCornerRadius(val *float64)
	MarkCornerRadiusInput() *float64
	SelectionColor() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeSelectionColorOutputReference
	SelectionColorInput() interface{}
	SentimentColors() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference
	SentimentColorsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TextBoxVerticalAlign() *string
	SetTextBoxVerticalAlign(val *string)
	TextBoxVerticalAlignInput() *string
	VisualizationColor01() *string
	SetVisualizationColor01(val *string)
	VisualizationColor01Input() *string
	VisualizationColor02() *string
	SetVisualizationColor02(val *string)
	VisualizationColor02Input() *string
	VisualizationColor03() *string
	SetVisualizationColor03(val *string)
	VisualizationColor03Input() *string
	VisualizationColor04() *string
	SetVisualizationColor04(val *string)
	VisualizationColor04Input() *string
	VisualizationColor05() *string
	SetVisualizationColor05(val *string)
	VisualizationColor05Input() *string
	VisualizationColor06() *string
	SetVisualizationColor06(val *string)
	VisualizationColor06Input() *string
	VisualizationColor07() *string
	SetVisualizationColor07(val *string)
	VisualizationColor07Input() *string
	VisualizationColor08() *string
	SetVisualizationColor08(val *string)
	VisualizationColor08Input() *string
	VisualizationColor09() *string
	SetVisualizationColor09(val *string)
	VisualizationColor09Input() *string
	VisualizationColor10() *string
	SetVisualizationColor10(val *string)
	VisualizationColor10Input() *string
	VisualizationColor11() *string
	SetVisualizationColor11(val *string)
	VisualizationColor11Input() *string
	VisualizationColor12() *string
	SetVisualizationColor12(val *string)
	VisualizationColor12Input() *string
	VisualizationColor13() *string
	SetVisualizationColor13(val *string)
	VisualizationColor13Input() *string
	VisualizationColor14() *string
	SetVisualizationColor14(val *string)
	VisualizationColor14Input() *string
	VisualizationColor15() *string
	SetVisualizationColor15(val *string)
	VisualizationColor15Input() *string
	VisualizationColor16() *string
	SetVisualizationColor16(val *string)
	VisualizationColor16Input() *string
	VisualizationColor17() *string
	SetVisualizationColor17(val *string)
	VisualizationColor17Input() *string
	VisualizationColor18() *string
	SetVisualizationColor18(val *string)
	VisualizationColor18Input() *string
	VisualizationColor19() *string
	SetVisualizationColor19(val *string)
	VisualizationColor19Input() *string
	VisualizationColor20() *string
	SetVisualizationColor20(val *string)
	VisualizationColor20Input() *string
	WidgetBackgroundColor() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBackgroundColorOutputReference
	WidgetBackgroundColorInput() interface{}
	WidgetBorderColor() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBorderColorOutputReference
	WidgetBorderColorInput() interface{}
	WidgetCornerRadius() *float64
	SetWidgetCornerRadius(val *float64)
	WidgetCornerRadiusInput() *float64
	WidgetHeaderAlignment() *string
	SetWidgetHeaderAlignment(val *string)
	WidgetHeaderAlignmentInput() *string
	WidgetMargin() *float64
	SetWidgetMargin(val *float64)
	WidgetMarginInput() *float64
	WidgetPadding() *float64
	SetWidgetPadding(val *float64)
	WidgetPaddingInput() *float64
	WidgetShadow() *float64
	SetWidgetShadow(val *float64)
	WidgetShadowInput() *float64
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
	PutAxisLineColor(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeAxisLineColor)
	PutCanvasBackgroundColor(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeCanvasBackgroundColor)
	PutContinuousGradient(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradient)
	PutFontColor(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontColor)
	PutFontSettings(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettings)
	PutGridLineColor(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeGridLineColor)
	PutSelectionColor(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeSelectionColor)
	PutSentimentColors(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeSentimentColors)
	PutWidgetBackgroundColor(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBackgroundColor)
	PutWidgetBorderColor(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBorderColor)
	ResetAxisLineColor()
	ResetCanvasBackgroundColor()
	ResetContinuousGradient()
	ResetFontColor()
	ResetFontFamily()
	ResetFontSettings()
	ResetGridLineColor()
	ResetMarkCornerRadius()
	ResetSelectionColor()
	ResetSentimentColors()
	ResetTextBoxVerticalAlign()
	ResetVisualizationColor01()
	ResetVisualizationColor02()
	ResetVisualizationColor03()
	ResetVisualizationColor04()
	ResetVisualizationColor05()
	ResetVisualizationColor06()
	ResetVisualizationColor07()
	ResetVisualizationColor08()
	ResetVisualizationColor09()
	ResetVisualizationColor10()
	ResetVisualizationColor11()
	ResetVisualizationColor12()
	ResetVisualizationColor13()
	ResetVisualizationColor14()
	ResetVisualizationColor15()
	ResetVisualizationColor16()
	ResetVisualizationColor17()
	ResetVisualizationColor18()
	ResetVisualizationColor19()
	ResetVisualizationColor20()
	ResetWidgetBackgroundColor()
	ResetWidgetBorderColor()
	ResetWidgetCornerRadius()
	ResetWidgetHeaderAlignment()
	ResetWidgetMargin()
	ResetWidgetPadding()
	ResetWidgetShadow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference
type jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) AxisLineColor() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeAxisLineColorOutputReference {
	var returns WorkspaceSettingV2EffectiveWorkspaceDashboardThemeAxisLineColorOutputReference
	_jsii_.Get(
		j,
		"axisLineColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) AxisLineColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"axisLineColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) CanvasBackgroundColor() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeCanvasBackgroundColorOutputReference {
	var returns WorkspaceSettingV2EffectiveWorkspaceDashboardThemeCanvasBackgroundColorOutputReference
	_jsii_.Get(
		j,
		"canvasBackgroundColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) CanvasBackgroundColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canvasBackgroundColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ContinuousGradient() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference {
	var returns WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradientOutputReference
	_jsii_.Get(
		j,
		"continuousGradient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ContinuousGradientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuousGradientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) FontColor() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontColorOutputReference {
	var returns WorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontColorOutputReference
	_jsii_.Get(
		j,
		"fontColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) FontColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fontColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) FontFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) FontFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) FontSettings() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference {
	var returns WorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettingsOutputReference
	_jsii_.Get(
		j,
		"fontSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) FontSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fontSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GridLineColor() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeGridLineColorOutputReference {
	var returns WorkspaceSettingV2EffectiveWorkspaceDashboardThemeGridLineColorOutputReference
	_jsii_.Get(
		j,
		"gridLineColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GridLineColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gridLineColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) InternalValue() *WorkspaceSettingV2EffectiveWorkspaceDashboardTheme {
	var returns *WorkspaceSettingV2EffectiveWorkspaceDashboardTheme
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) MarkCornerRadius() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"markCornerRadius",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) MarkCornerRadiusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"markCornerRadiusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) SelectionColor() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeSelectionColorOutputReference {
	var returns WorkspaceSettingV2EffectiveWorkspaceDashboardThemeSelectionColorOutputReference
	_jsii_.Get(
		j,
		"selectionColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) SelectionColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selectionColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) SentimentColors() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference {
	var returns WorkspaceSettingV2EffectiveWorkspaceDashboardThemeSentimentColorsOutputReference
	_jsii_.Get(
		j,
		"sentimentColors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) SentimentColorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sentimentColorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) TextBoxVerticalAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textBoxVerticalAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) TextBoxVerticalAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textBoxVerticalAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor01() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor01",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor01Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor01Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor02() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor02",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor02Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor02Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor03() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor03",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor03Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor03Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor04() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor04",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor04Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor04Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor05() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor05",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor05Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor05Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor06() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor06",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor06Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor06Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor07() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor07",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor07Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor07Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor08() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor08",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor08Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor08Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor09() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor09",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor09Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor09Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor10() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor10",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor10Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor10Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor11() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor11",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor11Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor11Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor12() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor12",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor12Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor12Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor13() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor13",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor13Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor13Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor14() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor14",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor14Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor14Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor15() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor15",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor15Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor15Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor16() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor16",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor16Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor16Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor17() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor17",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor17Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor17Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor18() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor18",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor18Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor18Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor19() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor19",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor19Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor19Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor20() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor20",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) VisualizationColor20Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationColor20Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetBackgroundColor() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBackgroundColorOutputReference {
	var returns WorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBackgroundColorOutputReference
	_jsii_.Get(
		j,
		"widgetBackgroundColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetBackgroundColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetBackgroundColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetBorderColor() WorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBorderColorOutputReference {
	var returns WorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBorderColorOutputReference
	_jsii_.Get(
		j,
		"widgetBorderColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetBorderColorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetBorderColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetCornerRadius() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetCornerRadius",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetCornerRadiusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetCornerRadiusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetHeaderAlignment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetHeaderAlignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetHeaderAlignmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetHeaderAlignmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetMargin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetMargin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetMarginInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetMarginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetPadding() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetPadding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetPaddingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetPaddingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetShadow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetShadow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) WidgetShadowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetShadowInput",
		&returns,
	)
	return returns
}


func NewWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference_Override(w WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-databricks.workspaceSettingV2.WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetFontFamily(val *string) {
	if err := j.validateSetFontFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontFamily",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetInternalValue(val *WorkspaceSettingV2EffectiveWorkspaceDashboardTheme) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetMarkCornerRadius(val *float64) {
	if err := j.validateSetMarkCornerRadiusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"markCornerRadius",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetTextBoxVerticalAlign(val *string) {
	if err := j.validateSetTextBoxVerticalAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textBoxVerticalAlign",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor01(val *string) {
	if err := j.validateSetVisualizationColor01Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor01",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor02(val *string) {
	if err := j.validateSetVisualizationColor02Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor02",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor03(val *string) {
	if err := j.validateSetVisualizationColor03Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor03",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor04(val *string) {
	if err := j.validateSetVisualizationColor04Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor04",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor05(val *string) {
	if err := j.validateSetVisualizationColor05Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor05",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor06(val *string) {
	if err := j.validateSetVisualizationColor06Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor06",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor07(val *string) {
	if err := j.validateSetVisualizationColor07Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor07",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor08(val *string) {
	if err := j.validateSetVisualizationColor08Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor08",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor09(val *string) {
	if err := j.validateSetVisualizationColor09Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor09",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor10(val *string) {
	if err := j.validateSetVisualizationColor10Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor10",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor11(val *string) {
	if err := j.validateSetVisualizationColor11Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor11",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor12(val *string) {
	if err := j.validateSetVisualizationColor12Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor12",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor13(val *string) {
	if err := j.validateSetVisualizationColor13Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor13",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor14(val *string) {
	if err := j.validateSetVisualizationColor14Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor14",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor15(val *string) {
	if err := j.validateSetVisualizationColor15Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor15",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor16(val *string) {
	if err := j.validateSetVisualizationColor16Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor16",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor17(val *string) {
	if err := j.validateSetVisualizationColor17Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor17",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor18(val *string) {
	if err := j.validateSetVisualizationColor18Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor18",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor19(val *string) {
	if err := j.validateSetVisualizationColor19Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor19",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetVisualizationColor20(val *string) {
	if err := j.validateSetVisualizationColor20Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualizationColor20",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetWidgetCornerRadius(val *float64) {
	if err := j.validateSetWidgetCornerRadiusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widgetCornerRadius",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetWidgetHeaderAlignment(val *string) {
	if err := j.validateSetWidgetHeaderAlignmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widgetHeaderAlignment",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetWidgetMargin(val *float64) {
	if err := j.validateSetWidgetMarginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widgetMargin",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetWidgetPadding(val *float64) {
	if err := j.validateSetWidgetPaddingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widgetPadding",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference)SetWidgetShadow(val *float64) {
	if err := j.validateSetWidgetShadowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widgetShadow",
		val,
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutAxisLineColor(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeAxisLineColor) {
	if err := w.validatePutAxisLineColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAxisLineColor",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutCanvasBackgroundColor(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeCanvasBackgroundColor) {
	if err := w.validatePutCanvasBackgroundColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCanvasBackgroundColor",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutContinuousGradient(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeContinuousGradient) {
	if err := w.validatePutContinuousGradientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putContinuousGradient",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutFontColor(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontColor) {
	if err := w.validatePutFontColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putFontColor",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutFontSettings(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeFontSettings) {
	if err := w.validatePutFontSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putFontSettings",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutGridLineColor(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeGridLineColor) {
	if err := w.validatePutGridLineColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putGridLineColor",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutSelectionColor(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeSelectionColor) {
	if err := w.validatePutSelectionColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSelectionColor",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutSentimentColors(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeSentimentColors) {
	if err := w.validatePutSentimentColorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSentimentColors",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutWidgetBackgroundColor(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBackgroundColor) {
	if err := w.validatePutWidgetBackgroundColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putWidgetBackgroundColor",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) PutWidgetBorderColor(value *WorkspaceSettingV2EffectiveWorkspaceDashboardThemeWidgetBorderColor) {
	if err := w.validatePutWidgetBorderColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putWidgetBorderColor",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetAxisLineColor() {
	_jsii_.InvokeVoid(
		w,
		"resetAxisLineColor",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetCanvasBackgroundColor() {
	_jsii_.InvokeVoid(
		w,
		"resetCanvasBackgroundColor",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetContinuousGradient() {
	_jsii_.InvokeVoid(
		w,
		"resetContinuousGradient",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetFontColor() {
	_jsii_.InvokeVoid(
		w,
		"resetFontColor",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetFontFamily() {
	_jsii_.InvokeVoid(
		w,
		"resetFontFamily",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetFontSettings() {
	_jsii_.InvokeVoid(
		w,
		"resetFontSettings",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetGridLineColor() {
	_jsii_.InvokeVoid(
		w,
		"resetGridLineColor",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetMarkCornerRadius() {
	_jsii_.InvokeVoid(
		w,
		"resetMarkCornerRadius",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetSelectionColor() {
	_jsii_.InvokeVoid(
		w,
		"resetSelectionColor",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetSentimentColors() {
	_jsii_.InvokeVoid(
		w,
		"resetSentimentColors",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetTextBoxVerticalAlign() {
	_jsii_.InvokeVoid(
		w,
		"resetTextBoxVerticalAlign",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor01() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor01",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor02() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor02",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor03() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor03",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor04() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor04",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor05() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor05",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor06() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor06",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor07() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor07",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor08() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor08",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor09() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor09",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor10() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor10",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor11() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor11",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor12() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor12",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor13() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor13",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor14() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor14",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor15() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor15",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor16() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor16",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor17() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor17",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor18() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor18",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor19() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor19",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetVisualizationColor20() {
	_jsii_.InvokeVoid(
		w,
		"resetVisualizationColor20",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetWidgetBackgroundColor() {
	_jsii_.InvokeVoid(
		w,
		"resetWidgetBackgroundColor",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetWidgetBorderColor() {
	_jsii_.InvokeVoid(
		w,
		"resetWidgetBorderColor",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetWidgetCornerRadius() {
	_jsii_.InvokeVoid(
		w,
		"resetWidgetCornerRadius",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetWidgetHeaderAlignment() {
	_jsii_.InvokeVoid(
		w,
		"resetWidgetHeaderAlignment",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetWidgetMargin() {
	_jsii_.InvokeVoid(
		w,
		"resetWidgetMargin",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetWidgetPadding() {
	_jsii_.InvokeVoid(
		w,
		"resetWidgetPadding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ResetWidgetShadow() {
	_jsii_.InvokeVoid(
		w,
		"resetWidgetShadow",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveWorkspaceDashboardThemeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

