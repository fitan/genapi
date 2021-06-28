package plugins

import (
	"fmt"
	public2 "github.com/fitan/genapi/public"
	"go/types"
	"log"
)

const (
	CasbinKeyserName     = "CasbinKeyser"
	CasbinListKeyserName = "CasbinListKeyser"
)

var HandlerTemplateMap = map[string]HandlerTemplate{CasbinKeyserName: {
	ImportPath: `"cmdb/public"`,
	Template: `
	data,err = public.CheckKeysCasbin(c,"%s",in.GetCasbinKeys())
	if err != nil {
		return data, err
	}`,
}, CasbinListKeyserName: {
	ImportPath: `"cmdb/public"`,
	Template: `
	data,err = public.CheckListKeysCasbin(c,"%s",in.GetCasbinKeys())
	if err != nil {
		return data, err
	}`,
}}

func GetCasbinPluginTemplate(DocFields []string, inFieldType types.Type) PluginTemplate {
	if len(DocFields) < 4 {
		log.Fatalln("@Casbin Mark need 3 parse")
	}

	pt := PluginTemplate{Has: true}
	pt.Keys = map[string]string{"key": DocFields[2], "annotation": DocFields[3]}
	i1 := CheckHasInterface(inFieldType, CasbinKeyserName)
	if i1 {
		importPath := public2.GetConfKey().GetPlugin("casbin").GetInInterface(CasbinKeyserName).After.ImportPath
		template := public2.GetConfKey().GetPlugin("casbin").GetInInterface(CasbinKeyserName).After.Template
		pt.InBindAfter = HandlerTemplate{importPath, fmt.Sprintf(template, pt.Keys["key"])}
	}

	i2 := CheckHasInterface(inFieldType, CasbinListKeyserName)
	if i2 {
		importPath := public2.GetConfKey().GetPlugin("casbin").GetInInterface(CasbinListKeyserName).After.ImportPath
		template := public2.GetConfKey().GetPlugin("casbin").GetInInterface(CasbinListKeyserName).After.Template
		pt.InBindAfter = HandlerTemplate{importPath, fmt.Sprintf(template, pt.Keys["key"])}
	}

	if i1 == false && i2 == false {
		log.Panic("casbin plugin not found ", CasbinKeyserName+CasbinListKeyserName)
	}
	return pt
}
