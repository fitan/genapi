package ui

type TableColumn struct {
	// 列头显示文字
	Title string `json:"title"`
	// 列宽度
	Width string `json:"width"`
	// 列数据在数据项中对应的 key，支持 a.b.c 的嵌套写法	
	DataIndex string `json:"data_index"`
	// 设置列内容的对齐方式 'left' | 'right' | 'center'
	Align string `json:"align"`
	// 超过宽度将自动省略，暂不支持和排序筛选一起使用。
	// 设置为 true 时，表格布局将变成 tableLayout="fixed"。
	Ellipsis bool `json:"ellipsis"`
	// 表头列合并,设置为 0 时，不渲染
	colSpan int `json:"col_span"`
	// 默认筛选值
	DefaultFilteredValue []string `json:"defaultFilteredValue"`
	// 标识数据是否经过过滤，筛选图标会高亮
	Filtered bool `json:"filtered"`
	// 筛选的受控属性，外界可用此控制列的筛选状态，值为已筛选的 value 数组	
	FilteredValue []string `json:"filteredValue"`
	// 是否多选
	FilterMultiple bool `json:"filterMultiple"`
	// 默认隐藏，可在列配置显示
	DefaultHiddena bool `json:"defaultHiddena"`
	// 列头右侧帮助文本 string｜string[]	
	HelpMessage interface{} `json:"helpMessage"`
	// 是否开启单元格编辑
	Edit bool `json:"edit"`
	// 是否开启行编辑
	EditRow bool `json:"editRow"`
	// 是否处于编辑状态
	Editable bool `json:"editable"`
}

func (t TableColumn) SetTitle(title string) TableColumn {
	return t
}

func (t TableColumn) SetWidth(width string) TableColumn {
	return t
}

func (t TableColumn) SetDataIndex(dataIndex string) TableColumn {
	return t
}

func (t TableColumn) SetAlign(align string) TableColumn {
	return t
}

func (t TableColumn) SetEllipsis(has bool) TableColumn {
	return t
}

func (t TableColumn) SetColSpan(span int) TableColumn {
	return t
}

func (t TableColumn) SetDefaultFilteredValue([]string) TableColumn {
	return t
}

func (t TableColumn) SetHelpMessage(msg interface{}) TableColumn {
	return t
}

func (t TableColumn) SetEdit(has bool) TableColumn {
	return t
}

func (t TableColumn) SetEditRow(has bool) TableColumn {
	return t
}

func (t TableColumn) SetEditable(has bool) TableColumn {
	return t
}

func (t TableColumn) SetApi(api interface{}) TableColumn {
	return t
}

type ActionItem struct {
	// 按钮文本
	Label string `json:"label"`
	// 是否禁用
	Disabled bool `json:"disabled"`
	// 按钮颜色 'success' | 'error' | 'warning'
	Color string `json:"color"`
	// 按钮类型
	ButionType string `json:"type"`
	// button组件props
	Props interface{} `json:"props"`
	// 按钮图标
	Icon string `json:"icon"`
	// 气泡确认框
	PopConfirm interface{} `json:"popConfirm"`
	// 是否显示分隔线，v2.0.0+
	Divider bool `json:"divider"`
	// 根据权限编码来控制当前列是否显示，v2.4.0+
	//auth?: RoleEnum | RoleEnum[] | string | string[];
	// 根据业务状态来控制当前列是否显示，v2.4.0+
	//ifShow?: boolean | ((action: ActionItem) => boolean);
	// 点击回调
	OnClick interface{} `json:"onClick"`
	// Tooltip配置，2.5.3以上版本支持，可以配置为string，或者完整的tooltip属性
	Tooltip interface{} `json:"tooltip"`

}




type ComponentType string

const  (
    Input ComponentType = "Input"
    InputNumber ComponentType = "InputNumber"
    Select ComponentType = "Select"
    ApiSelect ComponentType = "ApiSelect"
    Checkbox ComponentType = "Checkbox"
    Switch ComponentType = "Switch"
    DatePicker ComponentType = "DatePicker"  // v2.5.0 以上
    TimePicker ComponentType = "TimePicker"
)

type Tableer interface {
	Columns() []TableColumn
	Actions() []ActionItem
}

type Table struct {
	
}

func (t Table) Columns() []TableColumn {
	return []TableColumn{{
		Title:                "",
		Width:                "",
		DataIndex:            "",
		Align:                "",
		Ellipsis:             false,
		colSpan:              0,
		DefaultFilteredValue: nil,
		Filtered:             false,
		FilteredValue:        nil,
		FilterMultiple:       false,
		DefaultHiddena:       false,
		HelpMessage:          nil,
		Edit:                 false,
		EditRow:              false,
		Editable:             false,
	}}
}

func (t Table) Actions() []ActionItem {
	panic("implement me")
}


