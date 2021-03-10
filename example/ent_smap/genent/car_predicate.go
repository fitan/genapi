package genent

import (
	"ent_samp/ent"
	"ent_samp/ent/car"
	"ent_samp/ent/predicate"

	"time"
)

func CarPredicatesExec(fs ...func() (predicate.Car, error)) ([]predicate.Car, error) {
	ps := make([]predicate.Car, 0, len(fs))
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

type CarPaging struct {
	Limit int `form:"limit"`

	Page int `form:"page"`
}

func (m *CarPaging) BindPagingCar(queryer *ent.CarQuery) error {
	if m.Page == 0 {
		return nil
	}
	queryer.Limit(m.Limit).Offset((m.Page - 1) * m.Limit)
	return nil
}

type CarModelEQ struct {
	ModelEQ *string `json:"eq_model" form:"eq_model"`
}

func (m *CarModelEQ) BindCarModelEQ() (predicate.Car, error) {
	if m.ModelEQ == nil {
		return nil, nil
	}
	return car.ModelEQ(*m.ModelEQ), nil
}

type CarModelOr struct {
	ModelOr []string `form:"or_model"`
}

func (m *CarModelOr) BindCarModelOr() (predicate.Car, error) {
	if len(m.ModelOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.Car, 0, len(m.ModelOr))
	for i, _ := range m.ModelOr {
		predicate = append(predicate, car.ModelEQ(m.ModelOr[i]))
	}
	return car.Or(predicate...), nil
}

type CarModelNEQ struct {
	ModelNEQ *string `form:"neq_model"`
}

func (m *CarModelNEQ) BindCarModelNEQ() (predicate.Car, error) {
	if m.ModelNEQ == nil {
		return nil, nil
	}
	return car.ModelNEQ(*m.ModelNEQ), nil
}

type CarModelIn struct {
	ModelIn []string `form:"in_model"`
}

func (m *CarModelIn) BindCarModelIn() (predicate.Car, error) {
	if len(m.ModelIn) == 0 {
		return nil, nil
	}
	return car.ModelIn(m.ModelIn...), nil
}

type CarModelNotIn struct {
	ModelNotIn []string `form:"not_in_model"`
}

func (m *CarModelNotIn) BindCarModelNotIn() (predicate.Car, error) {
	if len(m.ModelNotIn) == 0 {
		return nil, nil
	}
	return car.ModelNotIn(m.ModelNotIn...), nil
}

type CarModelGT struct {
	ModelGT *string `form:"gt_model"`
}

func (m *CarModelGT) BindCarModelGT() (predicate.Car, error) {
	if m.ModelGT == nil {
		return nil, nil
	}
	return car.ModelGT(*m.ModelGT), nil
}

type CarModelGTE struct {
	ModelGTE *string `form:"gte_model"`
}

func (m *CarModelGTE) BindCarModelGTE() (predicate.Car, error) {
	if m.ModelGTE == nil {
		return nil, nil
	}
	return car.ModelGTE(*m.ModelGTE), nil
}

type CarModelLT struct {
	ModelLT *string `form:"lt_model"`
}

func (m *CarModelLT) BindCarModelLT() (predicate.Car, error) {
	if m.ModelLT == nil {
		return nil, nil
	}
	return car.ModelLT(*m.ModelLT), nil
}

type CarModelLTE struct {
	ModelLTE *string `form:"lte_model"`
}

func (m *CarModelLTE) BindCarModelLTE() (predicate.Car, error) {
	if m.ModelLTE == nil {
		return nil, nil
	}
	return car.ModelLTE(*m.ModelLTE), nil
}

type CarModelContains struct {
	ModelContains *string `form:"contains_model"`
}

func (m *CarModelContains) BindCarModelContains() (predicate.Car, error) {
	if m.ModelContains == nil {
		return nil, nil
	}
	return car.ModelContains(*m.ModelContains), nil
}

type CarModelHasPrefix struct {
	ModelHasPrefix *string `form:"has_prefix_model"`
}

func (m *CarModelHasPrefix) BindCarModelHasPrefix() (predicate.Car, error) {
	if m.ModelHasPrefix == nil {
		return nil, nil
	}
	return car.ModelHasPrefix(*m.ModelHasPrefix), nil

}

type CarModelHasSuffix struct {
	ModelHasSuffix *string `form:"has_suffix_model"`
}

func (m *CarModelHasSuffix) BindCarModelHasSuffix() (predicate.Car, error) {
	if m.ModelHasSuffix == nil {
		return nil, nil
	}
	return car.ModelHasSuffix(*m.ModelHasSuffix), nil
}

type CarRegisteredAtEQ struct {
	RegisteredAtEQ *time.Time `json:"eq_registered_at" form:"eq_registered_at"`
}

func (m *CarRegisteredAtEQ) BindCarRegisteredAtEQ() (predicate.Car, error) {
	if m.RegisteredAtEQ == nil {
		return nil, nil
	}
	return car.RegisteredAtEQ(*m.RegisteredAtEQ), nil
}

type CarRegisteredAtOr struct {
	RegisteredAtOr []time.Time `form:"or_registered_at"`
}

func (m *CarRegisteredAtOr) BindCarRegisteredAtOr() (predicate.Car, error) {
	if len(m.RegisteredAtOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.Car, 0, len(m.RegisteredAtOr))
	for i, _ := range m.RegisteredAtOr {
		predicate = append(predicate, car.RegisteredAtEQ(m.RegisteredAtOr[i]))
	}
	return car.Or(predicate...), nil
}

type CarRegisteredAtNEQ struct {
	RegisteredAtNEQ *time.Time `form:"neq_registered_at"`
}

func (m *CarRegisteredAtNEQ) BindCarRegisteredAtNEQ() (predicate.Car, error) {
	if m.RegisteredAtNEQ == nil {
		return nil, nil
	}
	return car.RegisteredAtNEQ(*m.RegisteredAtNEQ), nil
}

type CarRegisteredAtIn struct {
	RegisteredAtIn []time.Time `form:"in_registered_at"`
}

func (m *CarRegisteredAtIn) BindCarRegisteredAtIn() (predicate.Car, error) {
	if len(m.RegisteredAtIn) == 0 {
		return nil, nil
	}
	return car.RegisteredAtIn(m.RegisteredAtIn...), nil
}

type CarRegisteredAtNotIn struct {
	RegisteredAtNotIn []time.Time `form:"not_in_registered_at"`
}

func (m *CarRegisteredAtNotIn) BindCarRegisteredAtNotIn() (predicate.Car, error) {
	if len(m.RegisteredAtNotIn) == 0 {
		return nil, nil
	}
	return car.RegisteredAtNotIn(m.RegisteredAtNotIn...), nil
}

type CarRegisteredAtGT struct {
	RegisteredAtGT *time.Time `form:"gt_registered_at"`
}

func (m *CarRegisteredAtGT) BindCarRegisteredAtGT() (predicate.Car, error) {
	if m.RegisteredAtGT == nil {
		return nil, nil
	}
	return car.RegisteredAtGT(*m.RegisteredAtGT), nil
}

type CarRegisteredAtGTE struct {
	RegisteredAtGTE *time.Time `form:"gte_registered_at"`
}

func (m *CarRegisteredAtGTE) BindCarRegisteredAtGTE() (predicate.Car, error) {
	if m.RegisteredAtGTE == nil {
		return nil, nil
	}
	return car.RegisteredAtGTE(*m.RegisteredAtGTE), nil
}

type CarRegisteredAtLT struct {
	RegisteredAtLT *time.Time `form:"lt_registered_at"`
}

func (m *CarRegisteredAtLT) BindCarRegisteredAtLT() (predicate.Car, error) {
	if m.RegisteredAtLT == nil {
		return nil, nil
	}
	return car.RegisteredAtLT(*m.RegisteredAtLT), nil
}

type CarRegisteredAtLTE struct {
	RegisteredAtLTE *time.Time `form:"lte_registered_at"`
}

func (m *CarRegisteredAtLTE) BindCarRegisteredAtLTE() (predicate.Car, error) {
	if m.RegisteredAtLTE == nil {
		return nil, nil
	}
	return car.RegisteredAtLTE(*m.RegisteredAtLTE), nil
}
