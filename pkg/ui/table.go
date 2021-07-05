package ui

import (
	"github.com/gin-gonic/gin"
)

type FormType string

const (
	InputFormType            FormType = "Input"
	InputGroupFormType       FormType = "InputGroup"
	InputPasswordFormType    FormType = "InputPassword"
	InputSearchFormType      FormType = "InputSearch"
	InputTextAreaFormType    FormType = "InputTextArea"
	InputNumberFormType      FormType = "InputNumber"
	InputCountDownFormType   FormType = "InputCountDown"
	SelectFormType           FormType = "Select"
	ApiSelectFormType        FormType = "ApiSelect"
	TreeSelectFormType       FormType = "TreeSelect"
	RadioButtonGroupFormType FormType = "RadioButtonGroup"
	RadioGroupFormType       FormType = "RadioGroup"
	CheckboxFormType         FormType = "Checkbox"
	CheckboxGroupFormType    FormType = "CheckboxGroup"
	AutoCompleteFormType     FormType = "AutoComplete"
	CascaderFormType         FormType = "Cascader"
	DatePickerFormType       FormType = "DatePicker"
	MonthPickerFormType      FormType = "MonthPicker"
	RangePickerFormType      FormType = "RangePicker"
	WeekPickerFormType       FormType = "WeekPicker"
	TimePickerFormType       FormType = "TimePicker"
	SwitchFormType           FormType = "Switch"
	StrengthMeterFormType    FormType = "StrengthMeter"
	UploadFormType           FormType = "Upload"
	IconPickerFormType       FormType = "IconPicker"
	RenderFormType           FormType = "Render"
	SliderFormType           FormType = "Slider"
	RateFormType             FormType = "Rate"
)

type FormField struct {
	Field string
	Component string
	Lable string
	SubLabel string
	ColProps string
	Suffix string
	ChangeEvent string
	HelpMessage string
	Tests []TestData
	TestMap map[string]TestData
}

type TestData struct {
	Name string
	Hello string
}

func (f *FormField)SetField(value string) *FormField {
	return f
}

func (f *FormField) SetComponent(value string) *FormField {
	return f
}

func (f *FormField) SetLable(value string) *FormField {
	return f
}

func (f *FormField) SetSubLabel(value string) *FormField {
	return f
}

func (f *FormField) SetChangeEvent(value string) *FormField {
	return f
}

func (f *FormField) SetHelpMessge(value string) *FormField {
	return f
}

func (f *FormField) SetHandleSubmit(fc func(c *gin.Context) (data interface{}, err error)) *FormField {
	return f
}

func (f *FormField) SetHandleReset() {

}






