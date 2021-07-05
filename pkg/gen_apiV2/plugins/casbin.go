package plugins

import (
	public2 "github.com/fitan/genapi/public"
	"go/types"
	"log"
)



func GetCasbinPluginTemplate(docFields []string, inFieldType types.Type, outFieldType types.Type) PointTemplate {
	if len(docFields) < 4 {
		log.Println("@CasbinMark need 3 parse")
		panic(nil)
	}

	pt := PointTemplate{Has: true, BindBefor: HandlerTemplate{}, BindAfter: HandlerTemplate{}}
	pt.Keys = map[string]string{"key": docFields[2], "annotation": docFields[3]}
	pointConf := public2.GetConfKey().GetPoint("Casbin")

	for _, mount := range pointConf.Point.Mount {
		if CheckMatch(mount.Match, docFields, inFieldType, outFieldType) {
		//if CheckHasInterface(outFieldType, mount.Match.OutInterfaceName) && CheckHasInterface(inFieldType, mount.Match.InInterfaceName) {
			pt.BindBefor.Template = mount.MountBindBefor.Template
			pt.BindBefor.ImportPath = mount.MountBindBefor.ImportPath
			pt.BindAfter.Template = mount.MountBindAfter.Template
			pt.BindAfter.ImportPath = mount.MountBindAfter.ImportPath
			return pt
		}
	}
	//i1 := CheckHasInterface(inFieldType,)
	//if i1{
	//	pointConf := public2.GetConfKey().GetPoint("casbin").Point.Mount[0].Match.InInterfaceName
	//	importPath := public2.GetConfKey().GetPoint("casbin").GetInterface(CasbinKeyserName).BindAfter.ImportPath
	//	template := public2.GetConfKey().GetPoint("casbin").GetInterface(CasbinKeyserName).BindAfter.Template
	//	pt.BindAfter = HandlerTemplate{importPath,fmt.Sprintf(template, pt.Keys["key"])}
	//}
	//
	//i2 := CheckHasInterface(inFieldType,CasbinListKeyserName)
	//if i2 {
	//	importPath := public2.GetConfKey().GetPlugin("casbin").GetInterface(CasbinListKeyserName).BindAfter.ImportPath
	//	template := public2.GetConfKey().GetPlugin("casbin").GetInterface(CasbinListKeyserName).BindAfter.Template
	//	pt.BindAfter = HandlerTemplate{importPath, fmt.Sprintf(template, pt.Keys["key"])}
	//}
	//
	//if i1 == false && i2 == false {
	//	log.Panic("casbin plugin not found ", CasbinKeyserName +  CasbinListKeyserName)
	//}
	return pt
}