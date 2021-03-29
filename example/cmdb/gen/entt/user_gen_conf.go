package entt

import (
	"cmdb/ent"
	"cmdb/ent/predicate"
	"cmdb/ent/user"
)

type UserIncludes struct {
	Includes []string `form:"includes" json:"includes" binding:"dive,oneof=role_binding.project role_binding.project.service role_binding.project.service.server role_binding.service role_binding.service.project role_binding alert"`
}

type GetUserListData struct {
	Count  int
	Result []*ent.User
}

func UserSelete(queryer *ent.UserQuery) {
	queryer.Select(

		user.FieldCreateTime,

		user.FieldUpdateTime,

		user.FieldName,

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

func UserGetIDs(users ent.Users) []int {
	IDs := make([]int, 0, len(users))
	for i, _ := range users {
		IDs[i] = users[i].ID
	}
	return IDs
}

type UserDefaultQuery struct {
	UserIncludes

	UserNameEQ

	UserNameNEQ

	UserNameGT

	UserNameGTE

	UserPaging
}

func (u *UserDefaultQuery) PredicatesExec() ([]predicate.User, error) {
	return UserPredicatesExec(

		u.BindUserNameEQ,

		u.BindUserNameNEQ,

		u.BindUserNameGT,

		u.BindUserNameGTE,
	)
}

func (u *UserDefaultQuery) Exec(queryer *ent.UserQuery) error {
	ps, err := u.PredicatesExec()
	if err != nil {
		return err
	}
	QueryerIncludes(queryer, u.Includes)

	queryer.Where(user.And(ps...))

	u.BindPagingUser(queryer)

	return nil
}
