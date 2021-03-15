package entt

import (
	"cmdb/ent"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type IdsQuery struct {
	Ids []int `form:"ids" binding:"dive,gt=0"`
}

func BindIds(c *gin.Context) (IdsQuery, error) {
	ids := IdsQuery{}
	err := c.ShouldBindQuery(&ids)
	return ids, err
}

type IdUri struct {
	ID int `uri:"id" binding:"gt=0"`
}

func BindId(c *gin.Context) (IdUri, error) {
	id := IdUri{}
	err := c.ShouldBindUri(&id)
	return id, err
}

type Order struct {
	Order []string `form:"order"`
}

func (o *Order) BindOrder(expect map[string]int) ([]ent.OrderFunc, error) {
	astList := make([]string, 0, 0)
	descList := make([]string, 0, 0)

	for _, o := range o.Order {
		if strings.HasPrefix(o, "-") {
			if _, ok := expect[o[1:len(o)]]; !ok {
				return nil, fmt.Errorf("%v order Not allowed", o[1:len(o)])
			}
			descList = append(descList, o[1:len(o)])
			continue
		}
		if _, ok := expect[o]; !ok {
			return nil, fmt.Errorf("%v order Not allowed", o)
		}
		astList = append(astList, o)
	}
	return []ent.OrderFunc{ent.Asc(astList...), ent.Desc(descList...)}, nil
}

func BindOrder(c *gin.Context, expect map[string]int) ([]ent.OrderFunc, error) {
	bindQuery := Order{}
	err := c.ShouldBindQuery(&bindQuery)
	if err != nil {
		return nil, err
	}
	astList := make([]string, 0, 0)
	descList := make([]string, 0, 0)

	for _, o := range bindQuery.Order {
		if strings.HasPrefix(o, "-") {
			if _, ok := expect[o[1:len(o)]]; !ok {
				return nil, fmt.Errorf("%v order Not allowed", o[1:len(o)])
			}
			descList = append(descList, o[1:len(o)])
			continue
		}
		if _, ok := expect[o]; !ok {
			return nil, fmt.Errorf("%v order Not allowed", o)
		}
		astList = append(astList, o)
	}
	return []ent.OrderFunc{ent.Asc(astList...), ent.Desc(descList...)}, nil
}

type RestReturn struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  interface{} `json:"msg"`
}

func GinRestReturnFunc(c *gin.Context, f func(c *gin.Context) (interface{}, error)) {
	data, err := f(c)
	if err != nil {
		c.JSON(200, RestReturn{
			Code: 503,
			Data: data,
			Msg:  err.Error(),
		})
		return
	}

	c.JSON(200, RestReturn{
		Code: 200,
		Data: data,
		Msg:  nil,
	})
}

func RestReturnFunc(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.JSON(200, RestReturn{
			Code: 503,
			Data: data,
			Msg:  err.Error(),
		})
		return
	}

	c.JSON(200, RestReturn{
		Code: 200,
		Data: data,
		Msg:  nil,
	})
}
