package user

import (
	"cmdb/ent"
	"context"
	"github.com/gin-gonic/gin"
	"log"
)

type UserCallIn struct {
	Query struct {
		Id string `json:"id"`
	} `json:"query"`
}


func (i UserCallIn) GetCasbinKeys() [][]interface{} {
	k := make([][]interface{},0,0)
	k = append(k, []interface{}{i.Query.Id})
	return k
}

func (i UserCallIn) GetRedisKey() string {
	return i.Query.Id
}




// @GenApi /api/usercall [get]
// @CallBack redis get
// @Casbin UserCall 呼叫usercall
func UserCall(c *gin.Context, in *UserCallIn) ([]*ent.User, error) {
	return nil, nil
}

func UserCallHandler(c *gin.Context)  {
	b := &BaseContext{GinContext: c}
	request := UserCallIn{}
	c.ShouldBindQuery(&request.Query)
	b.Request = request
	b.ResponseData, b.ResponseErr = NewUserCall(b)
	b.SetResultJson()
}

func NewUserCall(c *BaseContext) (UserCallIn, error) {
	in := c.Request.(UserCallIn)
	return in, nil
}

type BaseContext struct {
	GinContext *gin.Context
	Trace context.Context
	XLog log.Logger
	Request interface{}
	ResponseData interface{}
	ResponseErr error
}

func (b *BaseContext)WarpResult() struct {
	Code int
	Data interface{}
	Msg  string
} {
	res :=  struct {
		Code int
		Data interface{}
		Msg string
	}{
		Code: 0,
		Data: b.ResponseData,
		Msg:  "",
	}

	if b.ResponseErr != nil {
		res.Code = 503
		res.Msg = b.ResponseErr.Error()
		return res
	}

	res.Code = 200
	return res
}

func (b *BaseContext)SetResultJson()  {
	b.GinContext.JSON(200, b.WarpResult())
}

