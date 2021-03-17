package entt

import (
	"cmdb/ent/predicate"
	"cmdb/ent/rolebinding"
	"github.com/gin-gonic/gin"

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
	Limit int `form:"limit" json:"limit"`

	Page int `form:"page" json:"page"`
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
	return rolebinding.CreateTimeEQ(*m.CreateTimeEQ), nil
}

type RoleBindingCreateTimeOr struct {
	CreateTimeOr []time.Time `json:"or_create_time" form:"or_create_time"`
}

func (m *RoleBindingCreateTimeOr) BindRoleBindingCreateTimeOr() (predicate.RoleBinding, error) {
	if len(m.CreateTimeOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.RoleBinding, 0, len(m.CreateTimeOr))
	for i, _ := range m.CreateTimeOr {
		predicate = append(predicate, rolebinding.CreateTimeEQ(m.CreateTimeOr[i]))
	}
	return rolebinding.Or(predicate...), nil
}

type RoleBindingCreateTimeNEQ struct {
	CreateTimeNEQ *time.Time `json:"neq_create_time" form:"neq_create_time"`
}

func (m *RoleBindingCreateTimeNEQ) BindRoleBindingCreateTimeNEQ() (predicate.RoleBinding, error) {
	if m.CreateTimeNEQ == nil {
		return nil, nil
	}
	return rolebinding.CreateTimeNEQ(*m.CreateTimeNEQ), nil
}

type RoleBindingCreateTimeIn struct {
	CreateTimeIn []time.Time `json:"in_create_time" form:"in_create_time"`
}

func (m *RoleBindingCreateTimeIn) BindRoleBindingCreateTimeIn() (predicate.RoleBinding, error) {
	if len(m.CreateTimeIn) == 0 {
		return nil, nil
	}
	return rolebinding.CreateTimeIn(m.CreateTimeIn...), nil
}

type RoleBindingCreateTimeNotIn struct {
	CreateTimeNotIn []time.Time `json:"not_in_create_time" form:"not_in_create_time"`
}

func (m *RoleBindingCreateTimeNotIn) BindRoleBindingCreateTimeNotIn() (predicate.RoleBinding, error) {
	if len(m.CreateTimeNotIn) == 0 {
		return nil, nil
	}
	return rolebinding.CreateTimeNotIn(m.CreateTimeNotIn...), nil
}

type RoleBindingCreateTimeGT struct {
	CreateTimeGT *time.Time `json:"gt_create_time" form:"gt_create_time"`
}

func (m *RoleBindingCreateTimeGT) BindRoleBindingCreateTimeGT() (predicate.RoleBinding, error) {
	if m.CreateTimeGT == nil {
		return nil, nil
	}
	return rolebinding.CreateTimeGT(*m.CreateTimeGT), nil
}

type RoleBindingCreateTimeGTE struct {
	CreateTimeGTE *time.Time `json:"gte_create_time" form:"gte_create_time"`
}

func (m *RoleBindingCreateTimeGTE) BindRoleBindingCreateTimeGTE() (predicate.RoleBinding, error) {
	if m.CreateTimeGTE == nil {
		return nil, nil
	}
	return rolebinding.CreateTimeGTE(*m.CreateTimeGTE), nil
}

type RoleBindingCreateTimeLT struct {
	CreateTimeLT *time.Time `json:"lt_create_time" form:"lt_create_time"`
}

func (m *RoleBindingCreateTimeLT) BindRoleBindingCreateTimeLT() (predicate.RoleBinding, error) {
	if m.CreateTimeLT == nil {
		return nil, nil
	}
	return rolebinding.CreateTimeLT(*m.CreateTimeLT), nil
}

type RoleBindingCreateTimeLTE struct {
	CreateTimeLTE *time.Time `json:"lte_create_time" form:"lte_create_time"`
}

func (m *RoleBindingCreateTimeLTE) BindRoleBindingCreateTimeLTE() (predicate.RoleBinding, error) {
	if m.CreateTimeLTE == nil {
		return nil, nil
	}
	return rolebinding.CreateTimeLTE(*m.CreateTimeLTE), nil
}

type RoleBindingUpdateTimeEQ struct {
	UpdateTimeEQ *time.Time `json:"eq_update_time" form:"eq_update_time"`
}

func (m *RoleBindingUpdateTimeEQ) BindRoleBindingUpdateTimeEQ() (predicate.RoleBinding, error) {
	if m.UpdateTimeEQ == nil {
		return nil, nil
	}
	return rolebinding.UpdateTimeEQ(*m.UpdateTimeEQ), nil
}

type RoleBindingUpdateTimeOr struct {
	UpdateTimeOr []time.Time `json:"or_update_time" form:"or_update_time"`
}

func (m *RoleBindingUpdateTimeOr) BindRoleBindingUpdateTimeOr() (predicate.RoleBinding, error) {
	if len(m.UpdateTimeOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.RoleBinding, 0, len(m.UpdateTimeOr))
	for i, _ := range m.UpdateTimeOr {
		predicate = append(predicate, rolebinding.UpdateTimeEQ(m.UpdateTimeOr[i]))
	}
	return rolebinding.Or(predicate...), nil
}

type RoleBindingUpdateTimeNEQ struct {
	UpdateTimeNEQ *time.Time `json:"neq_update_time" form:"neq_update_time"`
}

func (m *RoleBindingUpdateTimeNEQ) BindRoleBindingUpdateTimeNEQ() (predicate.RoleBinding, error) {
	if m.UpdateTimeNEQ == nil {
		return nil, nil
	}
	return rolebinding.UpdateTimeNEQ(*m.UpdateTimeNEQ), nil
}

type RoleBindingUpdateTimeIn struct {
	UpdateTimeIn []time.Time `json:"in_update_time" form:"in_update_time"`
}

func (m *RoleBindingUpdateTimeIn) BindRoleBindingUpdateTimeIn() (predicate.RoleBinding, error) {
	if len(m.UpdateTimeIn) == 0 {
		return nil, nil
	}
	return rolebinding.UpdateTimeIn(m.UpdateTimeIn...), nil
}

type RoleBindingUpdateTimeNotIn struct {
	UpdateTimeNotIn []time.Time `json:"not_in_update_time" form:"not_in_update_time"`
}

func (m *RoleBindingUpdateTimeNotIn) BindRoleBindingUpdateTimeNotIn() (predicate.RoleBinding, error) {
	if len(m.UpdateTimeNotIn) == 0 {
		return nil, nil
	}
	return rolebinding.UpdateTimeNotIn(m.UpdateTimeNotIn...), nil
}

type RoleBindingUpdateTimeGT struct {
	UpdateTimeGT *time.Time `json:"gt_update_time" form:"gt_update_time"`
}

func (m *RoleBindingUpdateTimeGT) BindRoleBindingUpdateTimeGT() (predicate.RoleBinding, error) {
	if m.UpdateTimeGT == nil {
		return nil, nil
	}
	return rolebinding.UpdateTimeGT(*m.UpdateTimeGT), nil
}

type RoleBindingUpdateTimeGTE struct {
	UpdateTimeGTE *time.Time `json:"gte_update_time" form:"gte_update_time"`
}

func (m *RoleBindingUpdateTimeGTE) BindRoleBindingUpdateTimeGTE() (predicate.RoleBinding, error) {
	if m.UpdateTimeGTE == nil {
		return nil, nil
	}
	return rolebinding.UpdateTimeGTE(*m.UpdateTimeGTE), nil
}

type RoleBindingUpdateTimeLT struct {
	UpdateTimeLT *time.Time `json:"lt_update_time" form:"lt_update_time"`
}

func (m *RoleBindingUpdateTimeLT) BindRoleBindingUpdateTimeLT() (predicate.RoleBinding, error) {
	if m.UpdateTimeLT == nil {
		return nil, nil
	}
	return rolebinding.UpdateTimeLT(*m.UpdateTimeLT), nil
}

type RoleBindingUpdateTimeLTE struct {
	UpdateTimeLTE *time.Time `json:"lte_update_time" form:"lte_update_time"`
}

func (m *RoleBindingUpdateTimeLTE) BindRoleBindingUpdateTimeLTE() (predicate.RoleBinding, error) {
	if m.UpdateTimeLTE == nil {
		return nil, nil
	}
	return rolebinding.UpdateTimeLTE(*m.UpdateTimeLTE), nil
}

type RoleBindingRoleEQ struct {
	RoleEQ *rolebinding.Role `json:"eq_role" form:"eq_role"`
}

func (m *RoleBindingRoleEQ) BindRoleBindingRoleEQ() (predicate.RoleBinding, error) {
	if m.RoleEQ == nil {
		return nil, nil
	}
	return rolebinding.RoleEQ(*m.RoleEQ), nil
}

type RoleBindingRoleOr struct {
	RoleOr []rolebinding.Role `json:"or_role" form:"or_role"`
}

func (m *RoleBindingRoleOr) BindRoleBindingRoleOr() (predicate.RoleBinding, error) {
	if len(m.RoleOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.RoleBinding, 0, len(m.RoleOr))
	for i, _ := range m.RoleOr {
		predicate = append(predicate, rolebinding.RoleEQ(m.RoleOr[i]))
	}
	return rolebinding.Or(predicate...), nil
}

type RoleBindingRoleNEQ struct {
	RoleNEQ *rolebinding.Role `json:"neq_role" form:"neq_role"`
}

func (m *RoleBindingRoleNEQ) BindRoleBindingRoleNEQ() (predicate.RoleBinding, error) {
	if m.RoleNEQ == nil {
		return nil, nil
	}
	return rolebinding.RoleNEQ(*m.RoleNEQ), nil
}

type RoleBindingRoleIn struct {
	RoleIn []rolebinding.Role `json:"in_role" form:"in_role"`
}

func (m *RoleBindingRoleIn) BindRoleBindingRoleIn() (predicate.RoleBinding, error) {
	if len(m.RoleIn) == 0 {
		return nil, nil
	}
	return rolebinding.RoleIn(m.RoleIn...), nil
}

type RoleBindingRoleNotIn struct {
	RoleNotIn []rolebinding.Role `json:"not_in_role" form:"not_in_role"`
}

func (m *RoleBindingRoleNotIn) BindRoleBindingRoleNotIn() (predicate.RoleBinding, error) {
	if len(m.RoleNotIn) == 0 {
		return nil, nil
	}
	return rolebinding.RoleNotIn(m.RoleNotIn...), nil
}
