package entt

import (
	"cmdb/ent/predicate"
	"cmdb/ent/service"
	"github.com/gin-gonic/gin"

	"time"
)

func ServicePredicatesExec(fs ...func() (predicate.Service, error)) ([]predicate.Service, error) {
	ps := make([]predicate.Service, 0, len(fs))
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

type ServicePaging struct {
	Limit int `form:"limit" json:"limit"`

	Page int `form:"page" json:"page"`
}

func (m *ServicePaging) BindPagingService(queryer *ent.ServiceQuery) error {
	if m.Page == 0 {
		return nil
	}
	queryer.Limit(m.Limit).Offset((m.Page - 1) * m.Limit)
	return nil
}

type ServiceCreateTimeEQ struct {
	CreateTimeEQ *time.Time `json:"eq_create_time" form:"eq_create_time"`
}

func (m *ServiceCreateTimeEQ) BindServiceCreateTimeEQ() (predicate.Service, error) {
	if m.CreateTimeEQ == nil {
		return nil, nil
	}
	return service.CreateTimeEQ(*m.CreateTimeEQ), nil
}

type ServiceCreateTimeOr struct {
	CreateTimeOr []time.Time `json:"or_create_time" form:"or_create_time"`
}

func (m *ServiceCreateTimeOr) BindServiceCreateTimeOr() (predicate.Service, error) {
	if len(m.CreateTimeOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.Service, 0, len(m.CreateTimeOr))
	for i, _ := range m.CreateTimeOr {
		predicate = append(predicate, service.CreateTimeEQ(m.CreateTimeOr[i]))
	}
	return service.Or(predicate...), nil
}

type ServiceCreateTimeNEQ struct {
	CreateTimeNEQ *time.Time `json:"neq_create_time" form:"neq_create_time"`
}

func (m *ServiceCreateTimeNEQ) BindServiceCreateTimeNEQ() (predicate.Service, error) {
	if m.CreateTimeNEQ == nil {
		return nil, nil
	}
	return service.CreateTimeNEQ(*m.CreateTimeNEQ), nil
}

type ServiceCreateTimeIn struct {
	CreateTimeIn []time.Time `json:"in_create_time" form:"in_create_time"`
}

func (m *ServiceCreateTimeIn) BindServiceCreateTimeIn() (predicate.Service, error) {
	if len(m.CreateTimeIn) == 0 {
		return nil, nil
	}
	return service.CreateTimeIn(m.CreateTimeIn...), nil
}

type ServiceCreateTimeNotIn struct {
	CreateTimeNotIn []time.Time `json:"not_in_create_time" form:"not_in_create_time"`
}

func (m *ServiceCreateTimeNotIn) BindServiceCreateTimeNotIn() (predicate.Service, error) {
	if len(m.CreateTimeNotIn) == 0 {
		return nil, nil
	}
	return service.CreateTimeNotIn(m.CreateTimeNotIn...), nil
}

type ServiceCreateTimeGT struct {
	CreateTimeGT *time.Time `json:"gt_create_time" form:"gt_create_time"`
}

func (m *ServiceCreateTimeGT) BindServiceCreateTimeGT() (predicate.Service, error) {
	if m.CreateTimeGT == nil {
		return nil, nil
	}
	return service.CreateTimeGT(*m.CreateTimeGT), nil
}

type ServiceCreateTimeGTE struct {
	CreateTimeGTE *time.Time `json:"gte_create_time" form:"gte_create_time"`
}

func (m *ServiceCreateTimeGTE) BindServiceCreateTimeGTE() (predicate.Service, error) {
	if m.CreateTimeGTE == nil {
		return nil, nil
	}
	return service.CreateTimeGTE(*m.CreateTimeGTE), nil
}

type ServiceCreateTimeLT struct {
	CreateTimeLT *time.Time `json:"lt_create_time" form:"lt_create_time"`
}

func (m *ServiceCreateTimeLT) BindServiceCreateTimeLT() (predicate.Service, error) {
	if m.CreateTimeLT == nil {
		return nil, nil
	}
	return service.CreateTimeLT(*m.CreateTimeLT), nil
}

type ServiceCreateTimeLTE struct {
	CreateTimeLTE *time.Time `json:"lte_create_time" form:"lte_create_time"`
}

func (m *ServiceCreateTimeLTE) BindServiceCreateTimeLTE() (predicate.Service, error) {
	if m.CreateTimeLTE == nil {
		return nil, nil
	}
	return service.CreateTimeLTE(*m.CreateTimeLTE), nil
}

type ServiceUpdateTimeEQ struct {
	UpdateTimeEQ *time.Time `json:"eq_update_time" form:"eq_update_time"`
}

func (m *ServiceUpdateTimeEQ) BindServiceUpdateTimeEQ() (predicate.Service, error) {
	if m.UpdateTimeEQ == nil {
		return nil, nil
	}
	return service.UpdateTimeEQ(*m.UpdateTimeEQ), nil
}

type ServiceUpdateTimeOr struct {
	UpdateTimeOr []time.Time `json:"or_update_time" form:"or_update_time"`
}

func (m *ServiceUpdateTimeOr) BindServiceUpdateTimeOr() (predicate.Service, error) {
	if len(m.UpdateTimeOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.Service, 0, len(m.UpdateTimeOr))
	for i, _ := range m.UpdateTimeOr {
		predicate = append(predicate, service.UpdateTimeEQ(m.UpdateTimeOr[i]))
	}
	return service.Or(predicate...), nil
}

type ServiceUpdateTimeNEQ struct {
	UpdateTimeNEQ *time.Time `json:"neq_update_time" form:"neq_update_time"`
}

func (m *ServiceUpdateTimeNEQ) BindServiceUpdateTimeNEQ() (predicate.Service, error) {
	if m.UpdateTimeNEQ == nil {
		return nil, nil
	}
	return service.UpdateTimeNEQ(*m.UpdateTimeNEQ), nil
}

type ServiceUpdateTimeIn struct {
	UpdateTimeIn []time.Time `json:"in_update_time" form:"in_update_time"`
}

func (m *ServiceUpdateTimeIn) BindServiceUpdateTimeIn() (predicate.Service, error) {
	if len(m.UpdateTimeIn) == 0 {
		return nil, nil
	}
	return service.UpdateTimeIn(m.UpdateTimeIn...), nil
}

type ServiceUpdateTimeNotIn struct {
	UpdateTimeNotIn []time.Time `json:"not_in_update_time" form:"not_in_update_time"`
}

func (m *ServiceUpdateTimeNotIn) BindServiceUpdateTimeNotIn() (predicate.Service, error) {
	if len(m.UpdateTimeNotIn) == 0 {
		return nil, nil
	}
	return service.UpdateTimeNotIn(m.UpdateTimeNotIn...), nil
}

type ServiceUpdateTimeGT struct {
	UpdateTimeGT *time.Time `json:"gt_update_time" form:"gt_update_time"`
}

func (m *ServiceUpdateTimeGT) BindServiceUpdateTimeGT() (predicate.Service, error) {
	if m.UpdateTimeGT == nil {
		return nil, nil
	}
	return service.UpdateTimeGT(*m.UpdateTimeGT), nil
}

type ServiceUpdateTimeGTE struct {
	UpdateTimeGTE *time.Time `json:"gte_update_time" form:"gte_update_time"`
}

func (m *ServiceUpdateTimeGTE) BindServiceUpdateTimeGTE() (predicate.Service, error) {
	if m.UpdateTimeGTE == nil {
		return nil, nil
	}
	return service.UpdateTimeGTE(*m.UpdateTimeGTE), nil
}

type ServiceUpdateTimeLT struct {
	UpdateTimeLT *time.Time `json:"lt_update_time" form:"lt_update_time"`
}

func (m *ServiceUpdateTimeLT) BindServiceUpdateTimeLT() (predicate.Service, error) {
	if m.UpdateTimeLT == nil {
		return nil, nil
	}
	return service.UpdateTimeLT(*m.UpdateTimeLT), nil
}

type ServiceUpdateTimeLTE struct {
	UpdateTimeLTE *time.Time `json:"lte_update_time" form:"lte_update_time"`
}

func (m *ServiceUpdateTimeLTE) BindServiceUpdateTimeLTE() (predicate.Service, error) {
	if m.UpdateTimeLTE == nil {
		return nil, nil
	}
	return service.UpdateTimeLTE(*m.UpdateTimeLTE), nil
}

type ServiceNameEQ struct {
	NameEQ *string `json:"eq_name" form:"eq_name"`
}

func (m *ServiceNameEQ) BindServiceNameEQ() (predicate.Service, error) {
	if m.NameEQ == nil {
		return nil, nil
	}
	return service.NameEQ(*m.NameEQ), nil
}

type ServiceNameOr struct {
	NameOr []string `json:"or_name" form:"or_name"`
}

func (m *ServiceNameOr) BindServiceNameOr() (predicate.Service, error) {
	if len(m.NameOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.Service, 0, len(m.NameOr))
	for i, _ := range m.NameOr {
		predicate = append(predicate, service.NameEQ(m.NameOr[i]))
	}
	return service.Or(predicate...), nil
}

type ServiceNameNEQ struct {
	NameNEQ *string `json:"neq_name" form:"neq_name"`
}

func (m *ServiceNameNEQ) BindServiceNameNEQ() (predicate.Service, error) {
	if m.NameNEQ == nil {
		return nil, nil
	}
	return service.NameNEQ(*m.NameNEQ), nil
}

type ServiceNameIn struct {
	NameIn []string `json:"in_name" form:"in_name"`
}

func (m *ServiceNameIn) BindServiceNameIn() (predicate.Service, error) {
	if len(m.NameIn) == 0 {
		return nil, nil
	}
	return service.NameIn(m.NameIn...), nil
}

type ServiceNameNotIn struct {
	NameNotIn []string `json:"not_in_name" form:"not_in_name"`
}

func (m *ServiceNameNotIn) BindServiceNameNotIn() (predicate.Service, error) {
	if len(m.NameNotIn) == 0 {
		return nil, nil
	}
	return service.NameNotIn(m.NameNotIn...), nil
}

type ServiceNameGT struct {
	NameGT *string `json:"gt_name" form:"gt_name"`
}

func (m *ServiceNameGT) BindServiceNameGT() (predicate.Service, error) {
	if m.NameGT == nil {
		return nil, nil
	}
	return service.NameGT(*m.NameGT), nil
}

type ServiceNameGTE struct {
	NameGTE *string `json:"gte_name" form:"gte_name"`
}

func (m *ServiceNameGTE) BindServiceNameGTE() (predicate.Service, error) {
	if m.NameGTE == nil {
		return nil, nil
	}
	return service.NameGTE(*m.NameGTE), nil
}

type ServiceNameLT struct {
	NameLT *string `json:"lt_name" form:"lt_name"`
}

func (m *ServiceNameLT) BindServiceNameLT() (predicate.Service, error) {
	if m.NameLT == nil {
		return nil, nil
	}
	return service.NameLT(*m.NameLT), nil
}

type ServiceNameLTE struct {
	NameLTE *string `json:"lte_name" form:"lte_name"`
}

func (m *ServiceNameLTE) BindServiceNameLTE() (predicate.Service, error) {
	if m.NameLTE == nil {
		return nil, nil
	}
	return service.NameLTE(*m.NameLTE), nil
}

type ServiceNameContains struct {
	NameContains *string `json:"contains_name" form:"contains_name"`
}

func (m *ServiceNameContains) BindServiceNameContains() (predicate.Service, error) {
	if m.NameContains == nil {
		return nil, nil
	}
	return service.NameContains(*m.NameContains), nil
}

type ServiceNameHasPrefix struct {
	NameHasPrefix *string `json:"has_prefix_name" form:"has_prefix_name"`
}

func (m *ServiceNameHasPrefix) BindServiceNameHasPrefix() (predicate.Service, error) {
	if m.NameHasPrefix == nil {
		return nil, nil
	}
	return service.NameHasPrefix(*m.NameHasPrefix), nil

}

type ServiceNameHasSuffix struct {
	NameHasSuffix *string `json:"has_suffix_name" form:"has_suffix_name"`
}

func (m *ServiceNameHasSuffix) BindServiceNameHasSuffix() (predicate.Service, error) {
	if m.NameHasSuffix == nil {
		return nil, nil
	}
	return service.NameHasSuffix(*m.NameHasSuffix), nil
}
