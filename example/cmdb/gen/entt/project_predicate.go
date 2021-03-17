package entt

import (
	"cmdb/ent/predicate"
	"cmdb/ent/project"
	"github.com/gin-gonic/gin"

	"time"
)

func ProjectPredicatesExec(fs ...func() (predicate.Project, error)) ([]predicate.Project, error) {
	ps := make([]predicate.Project, 0, len(fs))
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

type ProjectPaging struct {
	Limit int `form:"limit" json:"limit"`

	Page int `form:"page" json:"page"`
}

func (m *ProjectPaging) BindPagingProject(queryer *ent.ProjectQuery) error {
	if m.Page == 0 {
		return nil
	}
	queryer.Limit(m.Limit).Offset((m.Page - 1) * m.Limit)
	return nil
}

type ProjectCreateTimeEQ struct {
	CreateTimeEQ *time.Time `json:"eq_create_time" form:"eq_create_time"`
}

func (m *ProjectCreateTimeEQ) BindProjectCreateTimeEQ() (predicate.Project, error) {
	if m.CreateTimeEQ == nil {
		return nil, nil
	}
	return project.CreateTimeEQ(*m.CreateTimeEQ), nil
}

type ProjectCreateTimeOr struct {
	CreateTimeOr []time.Time `json:"or_create_time" form:"or_create_time"`
}

func (m *ProjectCreateTimeOr) BindProjectCreateTimeOr() (predicate.Project, error) {
	if len(m.CreateTimeOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.Project, 0, len(m.CreateTimeOr))
	for i, _ := range m.CreateTimeOr {
		predicate = append(predicate, project.CreateTimeEQ(m.CreateTimeOr[i]))
	}
	return project.Or(predicate...), nil
}

type ProjectCreateTimeNEQ struct {
	CreateTimeNEQ *time.Time `json:"neq_create_time" form:"neq_create_time"`
}

func (m *ProjectCreateTimeNEQ) BindProjectCreateTimeNEQ() (predicate.Project, error) {
	if m.CreateTimeNEQ == nil {
		return nil, nil
	}
	return project.CreateTimeNEQ(*m.CreateTimeNEQ), nil
}

type ProjectCreateTimeIn struct {
	CreateTimeIn []time.Time `json:"in_create_time" form:"in_create_time"`
}

func (m *ProjectCreateTimeIn) BindProjectCreateTimeIn() (predicate.Project, error) {
	if len(m.CreateTimeIn) == 0 {
		return nil, nil
	}
	return project.CreateTimeIn(m.CreateTimeIn...), nil
}

type ProjectCreateTimeNotIn struct {
	CreateTimeNotIn []time.Time `json:"not_in_create_time" form:"not_in_create_time"`
}

func (m *ProjectCreateTimeNotIn) BindProjectCreateTimeNotIn() (predicate.Project, error) {
	if len(m.CreateTimeNotIn) == 0 {
		return nil, nil
	}
	return project.CreateTimeNotIn(m.CreateTimeNotIn...), nil
}

type ProjectCreateTimeGT struct {
	CreateTimeGT *time.Time `json:"gt_create_time" form:"gt_create_time"`
}

func (m *ProjectCreateTimeGT) BindProjectCreateTimeGT() (predicate.Project, error) {
	if m.CreateTimeGT == nil {
		return nil, nil
	}
	return project.CreateTimeGT(*m.CreateTimeGT), nil
}

type ProjectCreateTimeGTE struct {
	CreateTimeGTE *time.Time `json:"gte_create_time" form:"gte_create_time"`
}

func (m *ProjectCreateTimeGTE) BindProjectCreateTimeGTE() (predicate.Project, error) {
	if m.CreateTimeGTE == nil {
		return nil, nil
	}
	return project.CreateTimeGTE(*m.CreateTimeGTE), nil
}

type ProjectCreateTimeLT struct {
	CreateTimeLT *time.Time `json:"lt_create_time" form:"lt_create_time"`
}

func (m *ProjectCreateTimeLT) BindProjectCreateTimeLT() (predicate.Project, error) {
	if m.CreateTimeLT == nil {
		return nil, nil
	}
	return project.CreateTimeLT(*m.CreateTimeLT), nil
}

type ProjectCreateTimeLTE struct {
	CreateTimeLTE *time.Time `json:"lte_create_time" form:"lte_create_time"`
}

func (m *ProjectCreateTimeLTE) BindProjectCreateTimeLTE() (predicate.Project, error) {
	if m.CreateTimeLTE == nil {
		return nil, nil
	}
	return project.CreateTimeLTE(*m.CreateTimeLTE), nil
}

type ProjectUpdateTimeEQ struct {
	UpdateTimeEQ *time.Time `json:"eq_update_time" form:"eq_update_time"`
}

func (m *ProjectUpdateTimeEQ) BindProjectUpdateTimeEQ() (predicate.Project, error) {
	if m.UpdateTimeEQ == nil {
		return nil, nil
	}
	return project.UpdateTimeEQ(*m.UpdateTimeEQ), nil
}

type ProjectUpdateTimeOr struct {
	UpdateTimeOr []time.Time `json:"or_update_time" form:"or_update_time"`
}

func (m *ProjectUpdateTimeOr) BindProjectUpdateTimeOr() (predicate.Project, error) {
	if len(m.UpdateTimeOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.Project, 0, len(m.UpdateTimeOr))
	for i, _ := range m.UpdateTimeOr {
		predicate = append(predicate, project.UpdateTimeEQ(m.UpdateTimeOr[i]))
	}
	return project.Or(predicate...), nil
}

type ProjectUpdateTimeNEQ struct {
	UpdateTimeNEQ *time.Time `json:"neq_update_time" form:"neq_update_time"`
}

func (m *ProjectUpdateTimeNEQ) BindProjectUpdateTimeNEQ() (predicate.Project, error) {
	if m.UpdateTimeNEQ == nil {
		return nil, nil
	}
	return project.UpdateTimeNEQ(*m.UpdateTimeNEQ), nil
}

type ProjectUpdateTimeIn struct {
	UpdateTimeIn []time.Time `json:"in_update_time" form:"in_update_time"`
}

func (m *ProjectUpdateTimeIn) BindProjectUpdateTimeIn() (predicate.Project, error) {
	if len(m.UpdateTimeIn) == 0 {
		return nil, nil
	}
	return project.UpdateTimeIn(m.UpdateTimeIn...), nil
}

type ProjectUpdateTimeNotIn struct {
	UpdateTimeNotIn []time.Time `json:"not_in_update_time" form:"not_in_update_time"`
}

func (m *ProjectUpdateTimeNotIn) BindProjectUpdateTimeNotIn() (predicate.Project, error) {
	if len(m.UpdateTimeNotIn) == 0 {
		return nil, nil
	}
	return project.UpdateTimeNotIn(m.UpdateTimeNotIn...), nil
}

type ProjectUpdateTimeGT struct {
	UpdateTimeGT *time.Time `json:"gt_update_time" form:"gt_update_time"`
}

func (m *ProjectUpdateTimeGT) BindProjectUpdateTimeGT() (predicate.Project, error) {
	if m.UpdateTimeGT == nil {
		return nil, nil
	}
	return project.UpdateTimeGT(*m.UpdateTimeGT), nil
}

type ProjectUpdateTimeGTE struct {
	UpdateTimeGTE *time.Time `json:"gte_update_time" form:"gte_update_time"`
}

func (m *ProjectUpdateTimeGTE) BindProjectUpdateTimeGTE() (predicate.Project, error) {
	if m.UpdateTimeGTE == nil {
		return nil, nil
	}
	return project.UpdateTimeGTE(*m.UpdateTimeGTE), nil
}

type ProjectUpdateTimeLT struct {
	UpdateTimeLT *time.Time `json:"lt_update_time" form:"lt_update_time"`
}

func (m *ProjectUpdateTimeLT) BindProjectUpdateTimeLT() (predicate.Project, error) {
	if m.UpdateTimeLT == nil {
		return nil, nil
	}
	return project.UpdateTimeLT(*m.UpdateTimeLT), nil
}

type ProjectUpdateTimeLTE struct {
	UpdateTimeLTE *time.Time `json:"lte_update_time" form:"lte_update_time"`
}

func (m *ProjectUpdateTimeLTE) BindProjectUpdateTimeLTE() (predicate.Project, error) {
	if m.UpdateTimeLTE == nil {
		return nil, nil
	}
	return project.UpdateTimeLTE(*m.UpdateTimeLTE), nil
}

type ProjectNameEQ struct {
	NameEQ *string `json:"eq_name" form:"eq_name"`
}

func (m *ProjectNameEQ) BindProjectNameEQ() (predicate.Project, error) {
	if m.NameEQ == nil {
		return nil, nil
	}
	return project.NameEQ(*m.NameEQ), nil
}

type ProjectNameOr struct {
	NameOr []string `json:"or_name" form:"or_name"`
}

func (m *ProjectNameOr) BindProjectNameOr() (predicate.Project, error) {
	if len(m.NameOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.Project, 0, len(m.NameOr))
	for i, _ := range m.NameOr {
		predicate = append(predicate, project.NameEQ(m.NameOr[i]))
	}
	return project.Or(predicate...), nil
}

type ProjectNameNEQ struct {
	NameNEQ *string `json:"neq_name" form:"neq_name"`
}

func (m *ProjectNameNEQ) BindProjectNameNEQ() (predicate.Project, error) {
	if m.NameNEQ == nil {
		return nil, nil
	}
	return project.NameNEQ(*m.NameNEQ), nil
}

type ProjectNameIn struct {
	NameIn []string `json:"in_name" form:"in_name"`
}

func (m *ProjectNameIn) BindProjectNameIn() (predicate.Project, error) {
	if len(m.NameIn) == 0 {
		return nil, nil
	}
	return project.NameIn(m.NameIn...), nil
}

type ProjectNameNotIn struct {
	NameNotIn []string `json:"not_in_name" form:"not_in_name"`
}

func (m *ProjectNameNotIn) BindProjectNameNotIn() (predicate.Project, error) {
	if len(m.NameNotIn) == 0 {
		return nil, nil
	}
	return project.NameNotIn(m.NameNotIn...), nil
}

type ProjectNameGT struct {
	NameGT *string `json:"gt_name" form:"gt_name"`
}

func (m *ProjectNameGT) BindProjectNameGT() (predicate.Project, error) {
	if m.NameGT == nil {
		return nil, nil
	}
	return project.NameGT(*m.NameGT), nil
}

type ProjectNameGTE struct {
	NameGTE *string `json:"gte_name" form:"gte_name"`
}

func (m *ProjectNameGTE) BindProjectNameGTE() (predicate.Project, error) {
	if m.NameGTE == nil {
		return nil, nil
	}
	return project.NameGTE(*m.NameGTE), nil
}

type ProjectNameLT struct {
	NameLT *string `json:"lt_name" form:"lt_name"`
}

func (m *ProjectNameLT) BindProjectNameLT() (predicate.Project, error) {
	if m.NameLT == nil {
		return nil, nil
	}
	return project.NameLT(*m.NameLT), nil
}

type ProjectNameLTE struct {
	NameLTE *string `json:"lte_name" form:"lte_name"`
}

func (m *ProjectNameLTE) BindProjectNameLTE() (predicate.Project, error) {
	if m.NameLTE == nil {
		return nil, nil
	}
	return project.NameLTE(*m.NameLTE), nil
}

type ProjectNameContains struct {
	NameContains *string `json:"contains_name" form:"contains_name"`
}

func (m *ProjectNameContains) BindProjectNameContains() (predicate.Project, error) {
	if m.NameContains == nil {
		return nil, nil
	}
	return project.NameContains(*m.NameContains), nil
}

type ProjectNameHasPrefix struct {
	NameHasPrefix *string `json:"has_prefix_name" form:"has_prefix_name"`
}

func (m *ProjectNameHasPrefix) BindProjectNameHasPrefix() (predicate.Project, error) {
	if m.NameHasPrefix == nil {
		return nil, nil
	}
	return project.NameHasPrefix(*m.NameHasPrefix), nil

}

type ProjectNameHasSuffix struct {
	NameHasSuffix *string `json:"has_suffix_name" form:"has_suffix_name"`
}

func (m *ProjectNameHasSuffix) BindProjectNameHasSuffix() (predicate.Project, error) {
	if m.NameHasSuffix == nil {
		return nil, nil
	}
	return project.NameHasSuffix(*m.NameHasSuffix), nil
}
