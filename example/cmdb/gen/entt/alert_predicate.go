package entt

import (
	"cmdb/ent"
	"cmdb/ent/alert"
	"cmdb/ent/predicate"
)

func AlertPredicatesExec(fs ...func() (predicate.Alert, error)) ([]predicate.Alert, error) {
	ps := make([]predicate.Alert, 0, len(fs))
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

type AlertPaging struct {
	Limit int `form:"limit" json:"limit"`

	Page int `form:"page" json:"page"`
}

func (m *AlertPaging) BindPagingAlert(queryer *ent.AlertQuery) error {
	if m.Page == 0 {
		return nil
	}
	queryer.Limit(m.Limit).Offset((m.Page - 1) * m.Limit)
	return nil
}

type AlertNameEQ struct {
	NameEQ *string `json:"eq_name" form:"eq_name"`
}

func (m *AlertNameEQ) BindAlertNameEQ() (predicate.Alert, error) {
	if m.NameEQ == nil {
		return nil, nil
	}
	return alert.NameEQ(*m.NameEQ), nil
}

type AlertNameOr struct {
	NameOr []string `json:"or_name" form:"or_name"`
}

func (m *AlertNameOr) BindAlertNameOr() (predicate.Alert, error) {
	if len(m.NameOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.Alert, 0, len(m.NameOr))
	for i, _ := range m.NameOr {
		predicate = append(predicate, alert.NameEQ(m.NameOr[i]))
	}
	return alert.Or(predicate...), nil
}

type AlertNameNEQ struct {
	NameNEQ *string `json:"neq_name" form:"neq_name"`
}

func (m *AlertNameNEQ) BindAlertNameNEQ() (predicate.Alert, error) {
	if m.NameNEQ == nil {
		return nil, nil
	}
	return alert.NameNEQ(*m.NameNEQ), nil
}

type AlertNameIn struct {
	NameIn []string `json:"in_name" form:"in_name"`
}

func (m *AlertNameIn) BindAlertNameIn() (predicate.Alert, error) {
	if len(m.NameIn) == 0 {
		return nil, nil
	}
	return alert.NameIn(m.NameIn...), nil
}

type AlertNameNotIn struct {
	NameNotIn []string `json:"not_in_name" form:"not_in_name"`
}

func (m *AlertNameNotIn) BindAlertNameNotIn() (predicate.Alert, error) {
	if len(m.NameNotIn) == 0 {
		return nil, nil
	}
	return alert.NameNotIn(m.NameNotIn...), nil
}

type AlertNameGT struct {
	NameGT *string `json:"gt_name" form:"gt_name"`
}

func (m *AlertNameGT) BindAlertNameGT() (predicate.Alert, error) {
	if m.NameGT == nil {
		return nil, nil
	}
	return alert.NameGT(*m.NameGT), nil
}

type AlertNameGTE struct {
	NameGTE *string `json:"gte_name" form:"gte_name"`
}

func (m *AlertNameGTE) BindAlertNameGTE() (predicate.Alert, error) {
	if m.NameGTE == nil {
		return nil, nil
	}
	return alert.NameGTE(*m.NameGTE), nil
}

type AlertNameLT struct {
	NameLT *string `json:"lt_name" form:"lt_name"`
}

func (m *AlertNameLT) BindAlertNameLT() (predicate.Alert, error) {
	if m.NameLT == nil {
		return nil, nil
	}
	return alert.NameLT(*m.NameLT), nil
}

type AlertNameLTE struct {
	NameLTE *string `json:"lte_name" form:"lte_name"`
}

func (m *AlertNameLTE) BindAlertNameLTE() (predicate.Alert, error) {
	if m.NameLTE == nil {
		return nil, nil
	}
	return alert.NameLTE(*m.NameLTE), nil
}

type AlertNameContains struct {
	NameContains *string `json:"contains_name" form:"contains_name"`
}

func (m *AlertNameContains) BindAlertNameContains() (predicate.Alert, error) {
	if m.NameContains == nil {
		return nil, nil
	}
	return alert.NameContains(*m.NameContains), nil
}

type AlertNameHasPrefix struct {
	NameHasPrefix *string `json:"has_prefix_name" form:"has_prefix_name"`
}

func (m *AlertNameHasPrefix) BindAlertNameHasPrefix() (predicate.Alert, error) {
	if m.NameHasPrefix == nil {
		return nil, nil
	}
	return alert.NameHasPrefix(*m.NameHasPrefix), nil

}

type AlertNameHasSuffix struct {
	NameHasSuffix *string `json:"has_suffix_name" form:"has_suffix_name"`
}

func (m *AlertNameHasSuffix) BindAlertNameHasSuffix() (predicate.Alert, error) {
	if m.NameHasSuffix == nil {
		return nil, nil
	}
	return alert.NameHasSuffix(*m.NameHasSuffix), nil
}
