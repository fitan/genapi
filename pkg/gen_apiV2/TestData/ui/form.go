package ui

import (
	"github.com/fitan/genapi/pkg/gen_apiV2/TestData"
	"github.com/fitan/genapi/pkg/gen_apiV2/ui"
)

type UserForm struct {

}

func (u *UserForm)Fields() []*ui.PointField {
	user := TestData.User{}
	return []*ui.PointField{
		ui.CreateFormField(user.ID).SetField("fsddfs").SetColProps("8m").SetComponent("fsdf"),
	}
}

func (u *UserForm) Submit()  {

}