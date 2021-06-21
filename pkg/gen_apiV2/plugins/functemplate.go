package plugins

var FuncTemplates = map[string]FuncTemplate{CasbinKeyserName: {
	ImportPath: "cmdb/public",
	Template:   `
	data,err = public.CheckKeysCasbin(c,%s,in.GetCasbinKeys())
	if err != nil {
		return data, err
	}`,
},CasbinListKeyserName: {
	ImportPath: "cmdb/public",
	Template:   `
	data,err = public.CheckListKeysCasbin(c,%s,in.GetCasbinKeys())
	if err != nil {
		return data, err
	}`,
}}

type FuncTemplate struct {
	ImportPath string
	Template string
}
