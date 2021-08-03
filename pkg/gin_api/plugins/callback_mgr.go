package plugins

import (
	"go/types"
	"log"
)

func init() {
	callBackMap = make(map[string]func(DocFields []string, inFieldType types.Type, outFieldType types.Type) CallBackTemplate)
	RegisterCallBack("redis", GetRedisCallBackTemplate)
}

const CallBackMark = "@CallBack"

var callBackMap  map[string]func(DocFields []string, inFieldType types.Type, outFieldType types.Type) CallBackTemplate


func RegisterCallBack(name string,fc func(DocFields []string, inFieldType types.Type, outFieldType types.Type) CallBackTemplate)  {
	callBackMap[name] = fc
}

func GetCallBackTemplate(docFields []string, inFieldType types.Type, outFieldType types.Type) CallBackTemplate {
	if len(docFields)<3 {
		log.Fatalln("@CallBack param must 3")
	}
	callBackName := docFields[2]
	if fc, ok := callBackMap[callBackName]; ok {
		return fc(docFields, inFieldType, outFieldType)
	}
	log.Fatalln("not found callback " + callBackName)
	return CallBackTemplate{}
}