package genrest

import (
	"ent_samp/ent"
	"ent_samp/ent/predicate"
	"ent_samp/ent/user"
)

func UserPredicatesExec(fs ...func() (predicate.User, error)) ([]predicate.User, error) {
	ps := make([]predicate.User, 0, len(fs))
	for _, f := range fs {
		p, err := f()
		if err != nil {
			return ps, err
		}
		if p != nil {
			ps = append(ps, p)
		}
	}
	return ps, nil
}

type UserPaging struct {
	Limit int `form:"limit" binding:"lte=100"`

	Page int `form:"page"`
}

func (m *UserPaging) BindPagingUser(queryer *ent.UserQuery) error {
	if m.Page == 0 {
		return nil
	}
	queryer.Limit(m.Limit).Offset((m.Page - 1) * m.Limit)
	return nil
}

type UserNameEQ struct {
	NameEQ *string `json:"eq_name" form:"eq_name"`
}

func (m *UserNameEQ) BindUserNameEQ() (predicate.User, error) {
	if m.NameEQ == nil {
		return nil, nil
	}
	return user.NameEQ(*m.NameEQ), nil
}

type UserNameOr struct {
	NameOr []string `form:"or_name"`
}

func (m *UserNameOr) BindUserNameOr() (predicate.User, error) {
	if len(m.NameOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.User, 0, len(m.NameOr))
	for i, _ := range m.NameOr {
		predicate = append(predicate, user.NameEQ(m.NameOr[i]))
	}
	return user.Or(predicate...), nil
}

type UserNameNEQ struct {
	NameNEQ *string `form:"neq_name"`
}

func (m *UserNameNEQ) BindUserNameNEQ() (predicate.User, error) {
	if m.NameNEQ == nil {
		return nil, nil
	}
	return user.NameNEQ(*m.NameNEQ), nil
}

type UserNameIn struct {
	NameIn []string `form:"in_name"`
}

func (m *UserNameIn) BindUserNameIn() (predicate.User, error) {
	if len(m.NameIn) == 0 {
		return nil, nil
	}
	return user.NameIn(m.NameIn...), nil
}

type UserNameNotIn struct {
	NameNotIn []string `form:"not_in_name"`
}

func (m *UserNameNotIn) BindUserNameNotIn() (predicate.User, error) {
	if len(m.NameNotIn) == 0 {
		return nil, nil
	}
	return user.NameNotIn(m.NameNotIn...), nil
}

type UserNameGT struct {
	NameGT *string `form:"gt_name"`
}

func (m *UserNameGT) BindUserNameGT() (predicate.User, error) {
	if m.NameGT == nil {
		return nil, nil
	}
	return user.NameGT(*m.NameGT), nil
}

type UserNameGTE struct {
	NameGTE *string `form:"gte_name"`
}

func (m *UserNameGTE) BindUserNameGTE() (predicate.User, error) {
	if m.NameGTE == nil {
		return nil, nil
	}
	return user.NameGTE(*m.NameGTE), nil
}

type UserNameLT struct {
	NameLT *string `form:"lt_name"`
}

func (m *UserNameLT) BindUserNameLT() (predicate.User, error) {
	if m.NameLT == nil {
		return nil, nil
	}
	return user.NameLT(*m.NameLT), nil
}

type UserNameLTE struct {
	NameLTE *string `form:"lte_name"`
}

func (m *UserNameLTE) BindUserNameLTE() (predicate.User, error) {
	if m.NameLTE == nil {
		return nil, nil
	}
	return user.NameLTE(*m.NameLTE), nil
}

type UserNameContains struct {
	NameContains *string `form:"contains_name"`
}

func (m *UserNameContains) BindUserNameContains() (predicate.User, error) {
	if m.NameContains == nil {
		return nil, nil
	}
	return user.NameContains(*m.NameContains), nil
}

type UserNameHasPrefix struct {
	NameHasPrefix *string `form:"has_prefix_name"`
}

func (m *UserNameHasPrefix) BindUserNameHasPrefix() (predicate.User, error) {
	if m.NameHasPrefix == nil {
		return nil, nil
	}
	return user.NameHasPrefix(*m.NameHasPrefix), nil

}

type UserNameHasSuffix struct {
	NameHasSuffix *string `form:"has_suffix_name"`
}

func (m *UserNameHasSuffix) BindUserNameHasSuffix() (predicate.User, error) {
	if m.NameHasSuffix == nil {
		return nil, nil
	}
	return user.NameHasSuffix(*m.NameHasSuffix), nil
}

type UserAge1EQ struct {
	Age1EQ *int `json:"eq_age1" form:"eq_age1"`
}

func (m *UserAge1EQ) BindUserAge1EQ() (predicate.User, error) {
	if m.Age1EQ == nil {
		return nil, nil
	}
	return user.Age1EQ(*m.Age1EQ), nil
}

type UserAge1Or struct {
	Age1Or []int `form:"or_age1"`
}

func (m *UserAge1Or) BindUserAge1Or() (predicate.User, error) {
	if len(m.Age1Or) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.User, 0, len(m.Age1Or))
	for i, _ := range m.Age1Or {
		predicate = append(predicate, user.Age1EQ(m.Age1Or[i]))
	}
	return user.Or(predicate...), nil
}

type UserAge1NEQ struct {
	Age1NEQ *int `form:"neq_age1"`
}

func (m *UserAge1NEQ) BindUserAge1NEQ() (predicate.User, error) {
	if m.Age1NEQ == nil {
		return nil, nil
	}
	return user.Age1NEQ(*m.Age1NEQ), nil
}

type UserAge1In struct {
	Age1In []int `form:"in_age1"`
}

func (m *UserAge1In) BindUserAge1In() (predicate.User, error) {
	if len(m.Age1In) == 0 {
		return nil, nil
	}
	return user.Age1In(m.Age1In...), nil
}

type UserAge1NotIn struct {
	Age1NotIn []int `form:"not_in_age1"`
}

func (m *UserAge1NotIn) BindUserAge1NotIn() (predicate.User, error) {
	if len(m.Age1NotIn) == 0 {
		return nil, nil
	}
	return user.Age1NotIn(m.Age1NotIn...), nil
}

type UserAge1GT struct {
	Age1GT *int `form:"gt_age1"`
}

func (m *UserAge1GT) BindUserAge1GT() (predicate.User, error) {
	if m.Age1GT == nil {
		return nil, nil
	}
	return user.Age1GT(*m.Age1GT), nil
}

type UserAge1GTE struct {
	Age1GTE *int `form:"gte_age1"`
}

func (m *UserAge1GTE) BindUserAge1GTE() (predicate.User, error) {
	if m.Age1GTE == nil {
		return nil, nil
	}
	return user.Age1GTE(*m.Age1GTE), nil
}

type UserAge1LT struct {
	Age1LT *int `form:"lt_age1"`
}

func (m *UserAge1LT) BindUserAge1LT() (predicate.User, error) {
	if m.Age1LT == nil {
		return nil, nil
	}
	return user.Age1LT(*m.Age1LT), nil
}

type UserAge1LTE struct {
	Age1LTE *int `form:"lte_age1"`
}

func (m *UserAge1LTE) BindUserAge1LTE() (predicate.User, error) {
	if m.Age1LTE == nil {
		return nil, nil
	}
	return user.Age1LTE(*m.Age1LTE), nil
}

type UserEnEQ struct {
	EnEQ *user.En `json:"eq_en" form:"eq_en"`
}

func (m *UserEnEQ) BindUserEnEQ() (predicate.User, error) {
	if m.EnEQ == nil {
		return nil, nil
	}
	return user.EnEQ(*m.EnEQ), nil
}

type UserEnOr struct {
	EnOr []user.En `form:"or_en"`
}

func (m *UserEnOr) BindUserEnOr() (predicate.User, error) {
	if len(m.EnOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.User, 0, len(m.EnOr))
	for i, _ := range m.EnOr {
		predicate = append(predicate, user.EnEQ(m.EnOr[i]))
	}
	return user.Or(predicate...), nil
}

type UserEnNEQ struct {
	EnNEQ *user.En `form:"neq_en"`
}

func (m *UserEnNEQ) BindUserEnNEQ() (predicate.User, error) {
	if m.EnNEQ == nil {
		return nil, nil
	}
	return user.EnNEQ(*m.EnNEQ), nil
}

type UserEnIn struct {
	EnIn []user.En `form:"in_en"`
}

func (m *UserEnIn) BindUserEnIn() (predicate.User, error) {
	if len(m.EnIn) == 0 {
		return nil, nil
	}
	return user.EnIn(m.EnIn...), nil
}

type UserEnNotIn struct {
	EnNotIn []user.En `form:"not_in_en"`
}

func (m *UserEnNotIn) BindUserEnNotIn() (predicate.User, error) {
	if len(m.EnNotIn) == 0 {
		return nil, nil
	}
	return user.EnNotIn(m.EnNotIn...), nil
}
