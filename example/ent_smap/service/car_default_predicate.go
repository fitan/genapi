package genrest

import (
	"ent_samp/ent"
	"ent_samp/ent/car"
	"ent_samp/ent/predicate"
)

type CarDefaultQuery struct {
}

func (c *CarDefaultQuery) PredicatesExec() ([]predicate.Car, error) {
	return CarPredicatesExec()
}

func (c *CarDefaultQuery) Exec(queryer *ent.CarQuery) error {
	ps, err := c.PredicatesExec()
	if err != nil {
		return err
	}

	queryer.Where(car.And(ps...))

	return nil
}
