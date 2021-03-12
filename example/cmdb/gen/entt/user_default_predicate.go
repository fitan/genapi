package entt

import (
	"cmdb/ent"
	"cmdb/ent/predicate"
	"cmdb/ent/user"
)

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
