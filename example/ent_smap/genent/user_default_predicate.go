package genent

import (
	"ent_samp/ent"
	"ent_samp/ent/predicate"
	"ent_samp/ent/user"
)

type UserDefaultQuery struct {
	UserAge1EQ

	UserPaging
}

func (u *UserDefaultQuery) PredicatesExec() ([]predicate.User, error) {
	return UserPredicatesExec(

		u.BindUserAge1EQ,
	)
}

func (u *UserDefaultQuery) Exec(queryer *ent.UserQuery) error {
	ps, err := u.PredicatesExec()
	if err != nil {
		return err
	}

	queryer.Where(user.And(ps...))

	u.BindPagingUser(queryer)

	return nil
}
