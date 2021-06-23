package ui

type TableField struct {
	Name string
	Type string
	CanBeModifi struct{
		Has bool
	}
}

type Name struct {
	
}

func (Name)Field() []TableField  {
	return []TableField{}
}

func (Name)FieldButton() {

}

func (Name)HeadButtion()  {

}
