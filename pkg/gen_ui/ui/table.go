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