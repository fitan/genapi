package public

import (
	"strings"
	"text/template"
)

type TempString string


func (s TempString) Format(data interface{}) (out string, err error) {
	t := template.Must(template.New("").Parse(string(s)))
	builder := &strings.Builder{}
	if err = t.Execute(builder, data); err != nil {
		return
	}
	out = builder.String()
	return
}

func GenForMat(s string, data interface{}) string {
	res,_ := TempString(s).Format(data)
	return res
}