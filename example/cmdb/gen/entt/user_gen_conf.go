package entt

import (
	"cmdb/ent"
	"cmdb/ent/user"
)

func UserSelete(queryer *ent.UserQuery) {
	queryer.Select(

		user.FieldCreateTime,

		user.FieldUpdateTime,

		user.FieldName,

		user.FieldPassword,

		user.FieldEmail,

		user.FieldPhone,

		user.FieldRole,
	)
}

func UserCreateMutation(m *ent.UserMutation, v *ent.User) {

	m.SetCreateTime(v.CreateTime)

	m.SetUpdateTime(v.UpdateTime)

	m.SetName(v.Name)

	m.SetPassword(v.Password)

	m.SetEmail(v.Email)

	m.SetPhone(v.Phone)

	m.SetRole(v.Role)

}

func UserUpdateMutation(m *ent.UserMutation, v *ent.User) {

	m.SetCreateTime(v.CreateTime)

	m.SetUpdateTime(v.UpdateTime)

	m.SetName(v.Name)

	m.SetPassword(v.Password)

	m.SetEmail(v.Email)

	m.SetPhone(v.Phone)

	m.SetRole(v.Role)

}
