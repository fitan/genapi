package ui

import (
	"github.com/fitan/genapi/pkg/gin_api/TestData"
	ui2 "github.com/fitan/genapi/pkg/gen_ui/ui"
)

type UserForm struct {

}

func (u *UserForm)Fields() []*ui2.PointField {
	user := TestData.User{}
	return []*ui2.PointField{
		ui2.CreateFormField(user.ID).SetField("fsddfs").SetColProps("8m").SetComponent("fsdf"),
	}
}

func (u *UserForm) Submit()  {

}