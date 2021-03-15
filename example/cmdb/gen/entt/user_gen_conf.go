package entt

import (
	"cmdb/ent"
	"cmdb/ent/predicate"
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

func UserGetIDs(users ent.Users) []int {
	IDs := make([]int, 0, len(users))
	for _, user := range users {
		IDs = append(IDs, user.ID)
	}
	return IDs
}

type UserDefaultQuery struct {
}

func (u *UserDefaultQuery) PredicatesExec() ([]predicate.User, error) {
	return UserPredicatesExec()
}

func (u *UserDefaultQuery) Exec(queryer *ent.UserQuery) error {
	ps, err := u.PredicatesExec()
	if err != nil {
		return err
	}

	queryer.Where(user.And(ps...))

	return nil
}
