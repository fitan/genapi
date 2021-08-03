package ui

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

func CreateFormField(point interface{}) *PointField {
	return &PointField{}
}


type PointField struct {
	Field string
	Component string
	Lable string
	SubLabel string
	ColProps string
	Suffix string
	ChangeEvent string
	HelpMessage string
}


func (f *PointField)SetField(value string) *PointField {
	return f
}

func (f *PointField) SetComponent(value string) *PointField {
	return f
}

func (f *PointField) SetLable(value string) *PointField {
	return f
}

func (f *PointField) SetSubLabel(value string) *PointField {
	return f
}

func (f *PointField) SetColProps(value string) *PointField {
	return f
}

func (f *PointField) SetSuffix(value string) *PointField {
	return f
}

func (f *PointField) SetChangeEvent(value string) *PointField {
	return f
}

func (f *PointField) SetHelpMessge(value string) *PointField {
	return f
}

//func (f *PointField) SetHandleSubmit(fc func(c *gin.Context) (data interface{}, err error)) *PointField {
//	return f
//}
//
//func (f *PointField) SetHandleReset() {
//
//}






