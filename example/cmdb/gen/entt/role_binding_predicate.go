package entt

import (
	"cmdb/ent"
	"cmdb/ent/predicate"
	"cmdb/ent/rolebinding"

	"time"
)

func RoleBindingPredicatesExec(fs ...func() (predicate.RoleBinding, error)) ([]predicate.RoleBinding, error) {
	ps := make([]predicate.RoleBinding, 0, len(fs))
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

type RoleBindingPaging struct {
	Limit int `form:"limit"`

	Page int `form:"page"`
}

func (m *RoleBindingPaging) BindPagingRoleBinding(queryer *ent.RoleBindingQuery) error {
	if m.Page == 0 {
		return nil
	}
	queryer.Limit(m.Limit).Offset((m.Page - 1) * m.Limit)
	return nil
}

type RoleBindingCreateTimeEQ struct {
	CreateTimeEQ *time.Time `json:"eq_create_time" form:"eq_create_time"`
}

func (m *RoleBindingCreateTimeEQ) BindRoleBindingCreateTimeEQ() (predicate.RoleBinding, error) {
	if m.CreateTimeEQ == nil {
		return nil, nil
	}
	return role_binding.CreateTimeEQ(*m.CreateTimeEQ), nil
}

type RoleBindingCreateTimeOr struct {
	CreateTimeOr []time.Time `form:"or_create_time"`
}

func (m *RoleBindingCreateTimeOr) BindRoleBindingCreateTimeOr() (predicate.RoleBinding, error) {
	if len(m.CreateTimeOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.RoleBinding, 0, len(m.CreateTimeOr))
	for i, _ := range m.CreateTimeOr {
		predicate = append(predicate, role_binding.CreateTimeEQ(m.CreateTimeOr[i]))
	}
	return role_binding.Or(predicate...), nil
}

type RoleBindingCreateTimeNEQ struct {
	CreateTimeNEQ *time.Time `form:"neq_create_time"`
}

func (m *RoleBindingCreateTimeNEQ) BindRoleBindingCreateTimeNEQ() (predicate.RoleBinding, error) {
	if m.CreateTimeNEQ == nil {
		return nil, nil
	}
	return role_binding.CreateTimeNEQ(*m.CreateTimeNEQ), nil
}

type RoleBindingCreateTimeIn struct {
	CreateTimeIn []time.Time `form:"in_create_time"`
}

func (m *RoleBindingCreateTimeIn) BindRoleBindingCreateTimeIn() (predicate.RoleBinding, error) {
	if len(m.CreateTimeIn) == 0 {
		return nil, nil
	}
	return role_binding.CreateTimeIn(m.CreateTimeIn...), nil
}

type RoleBindingCreateTimeNotIn struct {
	CreateTimeNotIn []time.Time `form:"not_in_create_time"`
}

func (m *RoleBindingCreateTimeNotIn) BindRoleBindingCreateTimeNotIn() (predicate.RoleBinding, error) {
	if len(m.CreateTimeNotIn) == 0 {
		return nil, nil
	}
	return role_binding.CreateTimeNotIn(m.CreateTimeNotIn...), nil
}

type RoleBindingCreateTimeGT struct {
	CreateTimeGT *time.Time `form:"gt_create_time"`
}

func (m *RoleBindingCreateTimeGT) BindRoleBindingCreateTimeGT() (predicate.RoleBinding, error) {
	if m.CreateTimeGT == nil {
		return nil, nil
	}
	return role_binding.CreateTimeGT(*m.CreateTimeGT), nil
}

type RoleBindingCreateTimeGTE struct {
	CreateTimeGTE *time.Time `form:"gte_create_time"`
}

func (m *RoleBindingCreateTimeGTE) BindRoleBindingCreateTimeGTE() (predicate.RoleBinding, error) {
	if m.CreateTimeGTE == nil {
		return nil, nil
	}
	return role_binding.CreateTimeGTE(*m.CreateTimeGTE), nil
}

type RoleBindingCreateTimeLT struct {
	CreateTimeLT *time.Time `form:"lt_create_time"`
}

func (m *RoleBindingCreateTimeLT) BindRoleBindingCreateTimeLT() (predicate.RoleBinding, error) {
	if m.CreateTimeLT == nil {
		return nil, nil
	}
	return role_binding.CreateTimeLT(*m.CreateTimeLT), nil
}

type RoleBindingCreateTimeLTE struct {
	CreateTimeLTE *time.Time `form:"lte_create_time"`
}

func (m *RoleBindingCreateTimeLTE) BindRoleBindingCreateTimeLTE() (predicate.RoleBinding, error) {
	if m.CreateTimeLTE == nil {
		return nil, nil
	}
	return role_binding.CreateTimeLTE(*m.CreateTimeLTE), nil
}

type RoleBindingUpdateTimeEQ struct {
	UpdateTimeEQ *time.Time `json:"eq_update_time" form:"eq_update_time"`
}

func (m *RoleBindingUpdateTimeEQ) BindRoleBindingUpdateTimeEQ() (predicate.RoleBinding, error) {
	if m.UpdateTimeEQ == nil {
		return nil, nil
	}
	return role_binding.UpdateTimeEQ(*m.UpdateTimeEQ), nil
}

type RoleBindingUpdateTimeOr struct {
	UpdateTimeOr []time.Time `form:"or_update_time"`
}

func (m *RoleBindingUpdateTimeOr) BindRoleBindingUpdateTimeOr() (predicate.RoleBinding, error) {
	if len(m.UpdateTimeOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.RoleBinding, 0, len(m.UpdateTimeOr))
	for i, _ := range m.UpdateTimeOr {
		predicate = append(predicate, role_binding.UpdateTimeEQ(m.UpdateTimeOr[i]))
	}
	return role_binding.Or(predicate...), nil
}

type RoleBindingUpdateTimeNEQ struct {
	UpdateTimeNEQ *time.Time `form:"neq_update_time"`
}

func (m *RoleBindingUpdateTimeNEQ) BindRoleBindingUpdateTimeNEQ() (predicate.RoleBinding, error) {
	if m.UpdateTimeNEQ == nil {
		return nil, nil
	}
	return role_binding.UpdateTimeNEQ(*m.UpdateTimeNEQ), nil
}

type RoleBindingUpdateTimeIn struct {
	UpdateTimeIn []time.Time `form:"in_update_time"`
}

func (m *RoleBindingUpdateTimeIn) BindRoleBindingUpdateTimeIn() (predicate.RoleBinding, error) {
	if len(m.UpdateTimeIn) == 0 {
		return nil, nil
	}
	return role_binding.UpdateTimeIn(m.UpdateTimeIn...), nil
}

type RoleBindingUpdateTimeNotIn struct {
	UpdateTimeNotIn []time.Time `form:"not_in_update_time"`
}

func (m *RoleBindingUpdateTimeNotIn) BindRoleBindingUpdateTimeNotIn() (predicate.RoleBinding, error) {
	if len(m.UpdateTimeNotIn) == 0 {
		return nil, nil
	}
	return role_binding.UpdateTimeNotIn(m.UpdateTimeNotIn...), nil
}

type RoleBindingUpdateTimeGT struct {
	UpdateTimeGT *time.Time `form:"gt_update_time"`
}

func (m *RoleBindingUpdateTimeGT) BindRoleBindingUpdateTimeGT() (predicate.RoleBinding, error) {
	if m.UpdateTimeGT == nil {
		return nil, nil
	}
	return role_binding.UpdateTimeGT(*m.UpdateTimeGT), nil
}

type RoleBindingUpdateTimeGTE struct {
	UpdateTimeGTE *time.Time `form:"gte_update_time"`
}

func (m *RoleBindingUpdateTimeGTE) BindRoleBindingUpdateTimeGTE() (predicate.RoleBinding, error) {
	if m.UpdateTimeGTE == nil {
		return nil, nil
	}
	return role_binding.UpdateTimeGTE(*m.UpdateTimeGTE), nil
}

type RoleBindingUpdateTimeLT struct {
	UpdateTimeLT *time.Time `form:"lt_update_time"`
}

func (m *RoleBindingUpdateTimeLT) BindRoleBindingUpdateTimeLT() (predicate.RoleBinding, error) {
	if m.UpdateTimeLT == nil {
		return nil, nil
	}
	return role_binding.UpdateTimeLT(*m.UpdateTimeLT), nil
}

type RoleBindingUpdateTimeLTE struct {
	UpdateTimeLTE *time.Time `form:"lte_update_time"`
}

func (m *RoleBindingUpdateTimeLTE) BindRoleBindingUpdateTimeLTE() (predicate.RoleBinding, error) {
	if m.UpdateTimeLTE == nil {
		return nil, nil
	}
	return role_binding.UpdateTimeLTE(*m.UpdateTimeLTE), nil
}

type RoleBindingRoleEQ struct {
	RoleEQ *rolebinding.Role `json:"eq_role" form:"eq_role"`
}

func (m *RoleBindingRoleEQ) BindRoleBindingRoleEQ() (predicate.RoleBinding, error) {
	if m.RoleEQ == nil {
		return nil, nil
	}
	return role_binding.RoleEQ(*m.RoleEQ), nil
}

type RoleBindingRoleOr struct {
	RoleOr []rolebinding.Role `form:"or_role"`
}

func (m *RoleBindingRoleOr) BindRoleBindingRoleOr() (predicate.RoleBinding, error) {
	if len(m.RoleOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.RoleBinding, 0, len(m.RoleOr))
	for i, _ := range m.RoleOr {
		predicate = append(predicate, role_binding.RoleEQ(m.RoleOr[i]))
	}
	return role_binding.Or(predicate...), nil
}

type RoleBindingRoleNEQ struct {
	RoleNEQ *rolebinding.Role `form:"neq_role"`
}

func (m *RoleBindingRoleNEQ) BindRoleBindingRoleNEQ() (predicate.RoleBinding, error) {
	if m.RoleNEQ == nil {
		return nil, nil
	}
	return role_binding.RoleNEQ(*m.RoleNEQ), nil
}

type RoleBindingRoleIn struct {
	RoleIn []rolebinding.Role `form:"in_role"`
}

func (m *RoleBindingRoleIn) BindRoleBindingRoleIn() (predicate.RoleBinding, error) {
	if len(m.RoleIn) == 0 {
		return nil, nil
	}
	return role_binding.RoleIn(m.RoleIn...), nil
}

type RoleBindingRoleNotIn struct {
	RoleNotIn []rolebinding.Role `form:"not_in_role"`
}

func (m *RoleBindingRoleNotIn) BindRoleBindingRoleNotIn() (predicate.RoleBinding, error) {
	if len(m.RoleNotIn) == 0 {
		return nil, nil
	}
	return role_binding.RoleNotIn(m.RoleNotIn...), nil
}
