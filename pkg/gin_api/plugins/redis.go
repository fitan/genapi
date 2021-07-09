package plugins

import (
	"github.com/fitan/genapi/public"
	"go/types"
	"log"
)



func GetRedisCallBackTemplate(docFields []string, inFieldType types.Type, outFieldType types.Type) CallBackTemplate {
	if len(docFields) < 4 {
		log.Fatalln("redis callback param must 4")
	}

	for _, cover := range public.GetConfKey().GetCallBack("redis").CallBack.Cover {
		if CheckMatch(cover.Match, docFields, inFieldType, outFieldType) {
			return CallBackTemplate{
				Has:      true,
				Keys:     nil,
				Template: HandlerTemplate{
					ImportPath: cover.ImportPath,
					Template:   cover.Template,
				},
			}
		}
	}
	return CallBackTemplate{}
}