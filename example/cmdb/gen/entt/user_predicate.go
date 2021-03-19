package entt

import (
	"cmdb/ent"
	"cmdb/ent/predicate"
	"cmdb/ent/user"

	"time"
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
	Limit int `form:"limit" binding:"lte=10" json:"limit"`

	Page int `form:"page" json:"page"`
}

func (m *UserPaging) BindPagingUser(queryer *ent.UserQuery) error {
	if m.Page == 0 {
		return nil
	}
	queryer.Limit(m.Limit).Offset((m.Page - 1) * m.Limit)
	return nil
}

type UserCreateTimeEQ struct {
	CreateTimeEQ *time.Time `json:"eq_create_time" form:"eq_create_time"`
}

func (m *UserCreateTimeEQ) BindUserCreateTimeEQ() (predicate.User, error) {
	if m.CreateTimeEQ == nil {
		return nil, nil
	}
	return user.CreateTimeEQ(*m.CreateTimeEQ), nil
}

type UserCreateTimeOr struct {
	CreateTimeOr []time.Time `json:"or_create_time" form:"or_create_time"`
}

func (m *UserCreateTimeOr) BindUserCreateTimeOr() (predicate.User, error) {
	if len(m.CreateTimeOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.User, 0, len(m.CreateTimeOr))
	for i, _ := range m.CreateTimeOr {
		predicate = append(predicate, user.CreateTimeEQ(m.CreateTimeOr[i]))
	}
	return user.Or(predicate...), nil
}

type UserCreateTimeNEQ struct {
	CreateTimeNEQ *time.Time `json:"neq_create_time" form:"neq_create_time"`
}

func (m *UserCreateTimeNEQ) BindUserCreateTimeNEQ() (predicate.User, error) {
	if m.CreateTimeNEQ == nil {
		return nil, nil
	}
	return user.CreateTimeNEQ(*m.CreateTimeNEQ), nil
}

type UserCreateTimeIn struct {
	CreateTimeIn []time.Time `json:"in_create_time" form:"in_create_time"`
}

func (m *UserCreateTimeIn) BindUserCreateTimeIn() (predicate.User, error) {
	if len(m.CreateTimeIn) == 0 {
		return nil, nil
	}
	return user.CreateTimeIn(m.CreateTimeIn...), nil
}

type UserCreateTimeNotIn struct {
	CreateTimeNotIn []time.Time `json:"not_in_create_time" form:"not_in_create_time"`
}

func (m *UserCreateTimeNotIn) BindUserCreateTimeNotIn() (predicate.User, error) {
	if len(m.CreateTimeNotIn) == 0 {
		return nil, nil
	}
	return user.CreateTimeNotIn(m.CreateTimeNotIn...), nil
}

type UserCreateTimeGT struct {
	CreateTimeGT *time.Time `json:"gt_create_time" form:"gt_create_time"`
}

func (m *UserCreateTimeGT) BindUserCreateTimeGT() (predicate.User, error) {
	if m.CreateTimeGT == nil {
		return nil, nil
	}
	return user.CreateTimeGT(*m.CreateTimeGT), nil
}

type UserCreateTimeGTE struct {
	CreateTimeGTE *time.Time `json:"gte_create_time" form:"gte_create_time"`
}

func (m *UserCreateTimeGTE) BindUserCreateTimeGTE() (predicate.User, error) {
	if m.CreateTimeGTE == nil {
		return nil, nil
	}
	return user.CreateTimeGTE(*m.CreateTimeGTE), nil
}

type UserCreateTimeLT struct {
	CreateTimeLT *time.Time `json:"lt_create_time" form:"lt_create_time"`
}

func (m *UserCreateTimeLT) BindUserCreateTimeLT() (predicate.User, error) {
	if m.CreateTimeLT == nil {
		return nil, nil
	}
	return user.CreateTimeLT(*m.CreateTimeLT), nil
}

type UserCreateTimeLTE struct {
	CreateTimeLTE *time.Time `json:"lte_create_time" form:"lte_create_time"`
}

func (m *UserCreateTimeLTE) BindUserCreateTimeLTE() (predicate.User, error) {
	if m.CreateTimeLTE == nil {
		return nil, nil
	}
	return user.CreateTimeLTE(*m.CreateTimeLTE), nil
}

type UserUpdateTimeEQ struct {
	UpdateTimeEQ *time.Time `json:"eq_update_time" form:"eq_update_time"`
}

func (m *UserUpdateTimeEQ) BindUserUpdateTimeEQ() (predicate.User, error) {
	if m.UpdateTimeEQ == nil {
		return nil, nil
	}
	return user.UpdateTimeEQ(*m.UpdateTimeEQ), nil
}

type UserUpdateTimeOr struct {
	UpdateTimeOr []time.Time `json:"or_update_time" form:"or_update_time"`
}

func (m *UserUpdateTimeOr) BindUserUpdateTimeOr() (predicate.User, error) {
	if len(m.UpdateTimeOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.User, 0, len(m.UpdateTimeOr))
	for i, _ := range m.UpdateTimeOr {
		predicate = append(predicate, user.UpdateTimeEQ(m.UpdateTimeOr[i]))
	}
	return user.Or(predicate...), nil
}

type UserUpdateTimeNEQ struct {
	UpdateTimeNEQ *time.Time `json:"neq_update_time" form:"neq_update_time"`
}

func (m *UserUpdateTimeNEQ) BindUserUpdateTimeNEQ() (predicate.User, error) {
	if m.UpdateTimeNEQ == nil {
		return nil, nil
	}
	return user.UpdateTimeNEQ(*m.UpdateTimeNEQ), nil
}

type UserUpdateTimeIn struct {
	UpdateTimeIn []time.Time `json:"in_update_time" form:"in_update_time"`
}

func (m *UserUpdateTimeIn) BindUserUpdateTimeIn() (predicate.User, error) {
	if len(m.UpdateTimeIn) == 0 {
		return nil, nil
	}
	return user.UpdateTimeIn(m.UpdateTimeIn...), nil
}

type UserUpdateTimeNotIn struct {
	UpdateTimeNotIn []time.Time `json:"not_in_update_time" form:"not_in_update_time"`
}

func (m *UserUpdateTimeNotIn) BindUserUpdateTimeNotIn() (predicate.User, error) {
	if len(m.UpdateTimeNotIn) == 0 {
		return nil, nil
	}
	return user.UpdateTimeNotIn(m.UpdateTimeNotIn...), nil
}

type UserUpdateTimeGT struct {
	UpdateTimeGT *time.Time `json:"gt_update_time" form:"gt_update_time"`
}

func (m *UserUpdateTimeGT) BindUserUpdateTimeGT() (predicate.User, error) {
	if m.UpdateTimeGT == nil {
		return nil, nil
	}
	return user.UpdateTimeGT(*m.UpdateTimeGT), nil
}

type UserUpdateTimeGTE struct {
	UpdateTimeGTE *time.Time `json:"gte_update_time" form:"gte_update_time"`
}

func (m *UserUpdateTimeGTE) BindUserUpdateTimeGTE() (predicate.User, error) {
	if m.UpdateTimeGTE == nil {
		return nil, nil
	}
	return user.UpdateTimeGTE(*m.UpdateTimeGTE), nil
}

type UserUpdateTimeLT struct {
	UpdateTimeLT *time.Time `json:"lt_update_time" form:"lt_update_time"`
}

func (m *UserUpdateTimeLT) BindUserUpdateTimeLT() (predicate.User, error) {
	if m.UpdateTimeLT == nil {
		return nil, nil
	}
	return user.UpdateTimeLT(*m.UpdateTimeLT), nil
}

type UserUpdateTimeLTE struct {
	UpdateTimeLTE *time.Time `json:"lte_update_time" form:"lte_update_time"`
}

func (m *UserUpdateTimeLTE) BindUserUpdateTimeLTE() (predicate.User, error) {
	if m.UpdateTimeLTE == nil {
		return nil, nil
	}
	return user.UpdateTimeLTE(*m.UpdateTimeLTE), nil
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
	NameOr []string `json:"or_name" form:"or_name"`
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
	NameNEQ *string `json:"neq_name" form:"neq_name"`
}

func (m *UserNameNEQ) BindUserNameNEQ() (predicate.User, error) {
	if m.NameNEQ == nil {
		return nil, nil
	}
	return user.NameNEQ(*m.NameNEQ), nil
}

type UserNameIn struct {
	NameIn []string `json:"in_name" form:"in_name"`
}

func (m *UserNameIn) BindUserNameIn() (predicate.User, error) {
	if len(m.NameIn) == 0 {
		return nil, nil
	}
	return user.NameIn(m.NameIn...), nil
}

type UserNameNotIn struct {
	NameNotIn []string `json:"not_in_name" form:"not_in_name"`
}

func (m *UserNameNotIn) BindUserNameNotIn() (predicate.User, error) {
	if len(m.NameNotIn) == 0 {
		return nil, nil
	}
	return user.NameNotIn(m.NameNotIn...), nil
}

type UserNameGT struct {
	NameGT *string `json:"gt_name" form:"gt_name"`
}

func (m *UserNameGT) BindUserNameGT() (predicate.User, error) {
	if m.NameGT == nil {
		return nil, nil
	}
	return user.NameGT(*m.NameGT), nil
}

type UserNameGTE struct {
	NameGTE *string `json:"gte_name" form:"gte_name"`
}

func (m *UserNameGTE) BindUserNameGTE() (predicate.User, error) {
	if m.NameGTE == nil {
		return nil, nil
	}
	return user.NameGTE(*m.NameGTE), nil
}

type UserNameLT struct {
	NameLT *string `json:"lt_name" form:"lt_name"`
}

func (m *UserNameLT) BindUserNameLT() (predicate.User, error) {
	if m.NameLT == nil {
		return nil, nil
	}
	return user.NameLT(*m.NameLT), nil
}

type UserNameLTE struct {
	NameLTE *string `json:"lte_name" form:"lte_name"`
}

func (m *UserNameLTE) BindUserNameLTE() (predicate.User, error) {
	if m.NameLTE == nil {
		return nil, nil
	}
	return user.NameLTE(*m.NameLTE), nil
}

type UserNameContains struct {
	NameContains *string `json:"contains_name" form:"contains_name"`
}

func (m *UserNameContains) BindUserNameContains() (predicate.User, error) {
	if m.NameContains == nil {
		return nil, nil
	}
	return user.NameContains(*m.NameContains), nil
}

type UserNameHasPrefix struct {
	NameHasPrefix *string `json:"has_prefix_name" form:"has_prefix_name"`
}

func (m *UserNameHasPrefix) BindUserNameHasPrefix() (predicate.User, error) {
	if m.NameHasPrefix == nil {
		return nil, nil
	}
	return user.NameHasPrefix(*m.NameHasPrefix), nil

}

type UserNameHasSuffix struct {
	NameHasSuffix *string `json:"has_suffix_name" form:"has_suffix_name"`
}

func (m *UserNameHasSuffix) BindUserNameHasSuffix() (predicate.User, error) {
	if m.NameHasSuffix == nil {
		return nil, nil
	}
	return user.NameHasSuffix(*m.NameHasSuffix), nil
}

type UserPasswordEQ struct {
	PasswordEQ *string `json:"eq_password" form:"eq_password"`
}

func (m *UserPasswordEQ) BindUserPasswordEQ() (predicate.User, error) {
	if m.PasswordEQ == nil {
		return nil, nil
	}
	return user.PasswordEQ(*m.PasswordEQ), nil
}

type UserPasswordOr struct {
	PasswordOr []string `json:"or_password" form:"or_password"`
}

func (m *UserPasswordOr) BindUserPasswordOr() (predicate.User, error) {
	if len(m.PasswordOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.User, 0, len(m.PasswordOr))
	for i, _ := range m.PasswordOr {
		predicate = append(predicate, user.PasswordEQ(m.PasswordOr[i]))
	}
	return user.Or(predicate...), nil
}

type UserPasswordNEQ struct {
	PasswordNEQ *string `json:"neq_password" form:"neq_password"`
}

func (m *UserPasswordNEQ) BindUserPasswordNEQ() (predicate.User, error) {
	if m.PasswordNEQ == nil {
		return nil, nil
	}
	return user.PasswordNEQ(*m.PasswordNEQ), nil
}

type UserPasswordIn struct {
	PasswordIn []string `json:"in_password" form:"in_password"`
}

func (m *UserPasswordIn) BindUserPasswordIn() (predicate.User, error) {
	if len(m.PasswordIn) == 0 {
		return nil, nil
	}
	return user.PasswordIn(m.PasswordIn...), nil
}

type UserPasswordNotIn struct {
	PasswordNotIn []string `json:"not_in_password" form:"not_in_password"`
}

func (m *UserPasswordNotIn) BindUserPasswordNotIn() (predicate.User, error) {
	if len(m.PasswordNotIn) == 0 {
		return nil, nil
	}
	return user.PasswordNotIn(m.PasswordNotIn...), nil
}

type UserPasswordGT struct {
	PasswordGT *string `json:"gt_password" form:"gt_password"`
}

func (m *UserPasswordGT) BindUserPasswordGT() (predicate.User, error) {
	if m.PasswordGT == nil {
		return nil, nil
	}
	return user.PasswordGT(*m.PasswordGT), nil
}

type UserPasswordGTE struct {
	PasswordGTE *string `json:"gte_password" form:"gte_password"`
}

func (m *UserPasswordGTE) BindUserPasswordGTE() (predicate.User, error) {
	if m.PasswordGTE == nil {
		return nil, nil
	}
	return user.PasswordGTE(*m.PasswordGTE), nil
}

type UserPasswordLT struct {
	PasswordLT *string `json:"lt_password" form:"lt_password"`
}

func (m *UserPasswordLT) BindUserPasswordLT() (predicate.User, error) {
	if m.PasswordLT == nil {
		return nil, nil
	}
	return user.PasswordLT(*m.PasswordLT), nil
}

type UserPasswordLTE struct {
	PasswordLTE *string `json:"lte_password" form:"lte_password"`
}

func (m *UserPasswordLTE) BindUserPasswordLTE() (predicate.User, error) {
	if m.PasswordLTE == nil {
		return nil, nil
	}
	return user.PasswordLTE(*m.PasswordLTE), nil
}

type UserPasswordContains struct {
	PasswordContains *string `json:"contains_password" form:"contains_password"`
}

func (m *UserPasswordContains) BindUserPasswordContains() (predicate.User, error) {
	if m.PasswordContains == nil {
		return nil, nil
	}
	return user.PasswordContains(*m.PasswordContains), nil
}

type UserPasswordHasPrefix struct {
	PasswordHasPrefix *string `json:"has_prefix_password" form:"has_prefix_password"`
}

func (m *UserPasswordHasPrefix) BindUserPasswordHasPrefix() (predicate.User, error) {
	if m.PasswordHasPrefix == nil {
		return nil, nil
	}
	return user.PasswordHasPrefix(*m.PasswordHasPrefix), nil

}

type UserPasswordHasSuffix struct {
	PasswordHasSuffix *string `json:"has_suffix_password" form:"has_suffix_password"`
}

func (m *UserPasswordHasSuffix) BindUserPasswordHasSuffix() (predicate.User, error) {
	if m.PasswordHasSuffix == nil {
		return nil, nil
	}
	return user.PasswordHasSuffix(*m.PasswordHasSuffix), nil
}

type UserEmailEQ struct {
	EmailEQ *string `json:"eq_email" form:"eq_email"`
}

func (m *UserEmailEQ) BindUserEmailEQ() (predicate.User, error) {
	if m.EmailEQ == nil {
		return nil, nil
	}
	return user.EmailEQ(*m.EmailEQ), nil
}

type UserEmailOr struct {
	EmailOr []string `json:"or_email" form:"or_email"`
}

func (m *UserEmailOr) BindUserEmailOr() (predicate.User, error) {
	if len(m.EmailOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.User, 0, len(m.EmailOr))
	for i, _ := range m.EmailOr {
		predicate = append(predicate, user.EmailEQ(m.EmailOr[i]))
	}
	return user.Or(predicate...), nil
}

type UserEmailNEQ struct {
	EmailNEQ *string `json:"neq_email" form:"neq_email"`
}

func (m *UserEmailNEQ) BindUserEmailNEQ() (predicate.User, error) {
	if m.EmailNEQ == nil {
		return nil, nil
	}
	return user.EmailNEQ(*m.EmailNEQ), nil
}

type UserEmailIn struct {
	EmailIn []string `json:"in_email" form:"in_email"`
}

func (m *UserEmailIn) BindUserEmailIn() (predicate.User, error) {
	if len(m.EmailIn) == 0 {
		return nil, nil
	}
	return user.EmailIn(m.EmailIn...), nil
}

type UserEmailNotIn struct {
	EmailNotIn []string `json:"not_in_email" form:"not_in_email"`
}

func (m *UserEmailNotIn) BindUserEmailNotIn() (predicate.User, error) {
	if len(m.EmailNotIn) == 0 {
		return nil, nil
	}
	return user.EmailNotIn(m.EmailNotIn...), nil
}

type UserEmailGT struct {
	EmailGT *string `json:"gt_email" form:"gt_email"`
}

func (m *UserEmailGT) BindUserEmailGT() (predicate.User, error) {
	if m.EmailGT == nil {
		return nil, nil
	}
	return user.EmailGT(*m.EmailGT), nil
}

type UserEmailGTE struct {
	EmailGTE *string `json:"gte_email" form:"gte_email"`
}

func (m *UserEmailGTE) BindUserEmailGTE() (predicate.User, error) {
	if m.EmailGTE == nil {
		return nil, nil
	}
	return user.EmailGTE(*m.EmailGTE), nil
}

type UserEmailLT struct {
	EmailLT *string `json:"lt_email" form:"lt_email"`
}

func (m *UserEmailLT) BindUserEmailLT() (predicate.User, error) {
	if m.EmailLT == nil {
		return nil, nil
	}
	return user.EmailLT(*m.EmailLT), nil
}

type UserEmailLTE struct {
	EmailLTE *string `json:"lte_email" form:"lte_email"`
}

func (m *UserEmailLTE) BindUserEmailLTE() (predicate.User, error) {
	if m.EmailLTE == nil {
		return nil, nil
	}
	return user.EmailLTE(*m.EmailLTE), nil
}

type UserEmailContains struct {
	EmailContains *string `json:"contains_email" form:"contains_email"`
}

func (m *UserEmailContains) BindUserEmailContains() (predicate.User, error) {
	if m.EmailContains == nil {
		return nil, nil
	}
	return user.EmailContains(*m.EmailContains), nil
}

type UserEmailHasPrefix struct {
	EmailHasPrefix *string `json:"has_prefix_email" form:"has_prefix_email"`
}

func (m *UserEmailHasPrefix) BindUserEmailHasPrefix() (predicate.User, error) {
	if m.EmailHasPrefix == nil {
		return nil, nil
	}
	return user.EmailHasPrefix(*m.EmailHasPrefix), nil

}

type UserEmailHasSuffix struct {
	EmailHasSuffix *string `json:"has_suffix_email" form:"has_suffix_email"`
}

func (m *UserEmailHasSuffix) BindUserEmailHasSuffix() (predicate.User, error) {
	if m.EmailHasSuffix == nil {
		return nil, nil
	}
	return user.EmailHasSuffix(*m.EmailHasSuffix), nil
}

type UserPhoneEQ struct {
	PhoneEQ *string `json:"eq_phone" form:"eq_phone"`
}

func (m *UserPhoneEQ) BindUserPhoneEQ() (predicate.User, error) {
	if m.PhoneEQ == nil {
		return nil, nil
	}
	return user.PhoneEQ(*m.PhoneEQ), nil
}

type UserPhoneOr struct {
	PhoneOr []string `json:"or_phone" form:"or_phone"`
}

func (m *UserPhoneOr) BindUserPhoneOr() (predicate.User, error) {
	if len(m.PhoneOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.User, 0, len(m.PhoneOr))
	for i, _ := range m.PhoneOr {
		predicate = append(predicate, user.PhoneEQ(m.PhoneOr[i]))
	}
	return user.Or(predicate...), nil
}

type UserPhoneNEQ struct {
	PhoneNEQ *string `json:"neq_phone" form:"neq_phone"`
}

func (m *UserPhoneNEQ) BindUserPhoneNEQ() (predicate.User, error) {
	if m.PhoneNEQ == nil {
		return nil, nil
	}
	return user.PhoneNEQ(*m.PhoneNEQ), nil
}

type UserPhoneIn struct {
	PhoneIn []string `json:"in_phone" form:"in_phone"`
}

func (m *UserPhoneIn) BindUserPhoneIn() (predicate.User, error) {
	if len(m.PhoneIn) == 0 {
		return nil, nil
	}
	return user.PhoneIn(m.PhoneIn...), nil
}

type UserPhoneNotIn struct {
	PhoneNotIn []string `json:"not_in_phone" form:"not_in_phone"`
}

func (m *UserPhoneNotIn) BindUserPhoneNotIn() (predicate.User, error) {
	if len(m.PhoneNotIn) == 0 {
		return nil, nil
	}
	return user.PhoneNotIn(m.PhoneNotIn...), nil
}

type UserPhoneGT struct {
	PhoneGT *string `json:"gt_phone" form:"gt_phone"`
}

func (m *UserPhoneGT) BindUserPhoneGT() (predicate.User, error) {
	if m.PhoneGT == nil {
		return nil, nil
	}
	return user.PhoneGT(*m.PhoneGT), nil
}

type UserPhoneGTE struct {
	PhoneGTE *string `json:"gte_phone" form:"gte_phone"`
}

func (m *UserPhoneGTE) BindUserPhoneGTE() (predicate.User, error) {
	if m.PhoneGTE == nil {
		return nil, nil
	}
	return user.PhoneGTE(*m.PhoneGTE), nil
}

type UserPhoneLT struct {
	PhoneLT *string `json:"lt_phone" form:"lt_phone"`
}

func (m *UserPhoneLT) BindUserPhoneLT() (predicate.User, error) {
	if m.PhoneLT == nil {
		return nil, nil
	}
	return user.PhoneLT(*m.PhoneLT), nil
}

type UserPhoneLTE struct {
	PhoneLTE *string `json:"lte_phone" form:"lte_phone"`
}

func (m *UserPhoneLTE) BindUserPhoneLTE() (predicate.User, error) {
	if m.PhoneLTE == nil {
		return nil, nil
	}
	return user.PhoneLTE(*m.PhoneLTE), nil
}

type UserPhoneContains struct {
	PhoneContains *string `json:"contains_phone" form:"contains_phone"`
}

func (m *UserPhoneContains) BindUserPhoneContains() (predicate.User, error) {
	if m.PhoneContains == nil {
		return nil, nil
	}
	return user.PhoneContains(*m.PhoneContains), nil
}

type UserPhoneHasPrefix struct {
	PhoneHasPrefix *string `json:"has_prefix_phone" form:"has_prefix_phone"`
}

func (m *UserPhoneHasPrefix) BindUserPhoneHasPrefix() (predicate.User, error) {
	if m.PhoneHasPrefix == nil {
		return nil, nil
	}
	return user.PhoneHasPrefix(*m.PhoneHasPrefix), nil

}

type UserPhoneHasSuffix struct {
	PhoneHasSuffix *string `json:"has_suffix_phone" form:"has_suffix_phone"`
}

func (m *UserPhoneHasSuffix) BindUserPhoneHasSuffix() (predicate.User, error) {
	if m.PhoneHasSuffix == nil {
		return nil, nil
	}
	return user.PhoneHasSuffix(*m.PhoneHasSuffix), nil
}

type UserRoleEQ struct {
	RoleEQ *user.Role `json:"eq_role" form:"eq_role"`
}

func (m *UserRoleEQ) BindUserRoleEQ() (predicate.User, error) {
	if m.RoleEQ == nil {
		return nil, nil
	}
	return user.RoleEQ(*m.RoleEQ), nil
}

type UserRoleOr struct {
	RoleOr []user.Role `json:"or_role" form:"or_role"`
}

func (m *UserRoleOr) BindUserRoleOr() (predicate.User, error) {
	if len(m.RoleOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.User, 0, len(m.RoleOr))
	for i, _ := range m.RoleOr {
		predicate = append(predicate, user.RoleEQ(m.RoleOr[i]))
	}
	return user.Or(predicate...), nil
}

type UserRoleNEQ struct {
	RoleNEQ *user.Role `json:"neq_role" form:"neq_role"`
}

func (m *UserRoleNEQ) BindUserRoleNEQ() (predicate.User, error) {
	if m.RoleNEQ == nil {
		return nil, nil
	}
	return user.RoleNEQ(*m.RoleNEQ), nil
}

type UserRoleIn struct {
	RoleIn []user.Role `json:"in_role" form:"in_role"`
}

func (m *UserRoleIn) BindUserRoleIn() (predicate.User, error) {
	if len(m.RoleIn) == 0 {
		return nil, nil
	}
	return user.RoleIn(m.RoleIn...), nil
}

type UserRoleNotIn struct {
	RoleNotIn []user.Role `json:"not_in_role" form:"not_in_role"`
}

func (m *UserRoleNotIn) BindUserRoleNotIn() (predicate.User, error) {
	if len(m.RoleNotIn) == 0 {
		return nil, nil
	}
	return user.RoleNotIn(m.RoleNotIn...), nil
}
