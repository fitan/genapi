package genent

import (
	"ent_samp/ent"
)

type CURDALL struct {
	Car  *CarCURD
	User *UserCURD
}

func NewCURDALL(db *ent.Client) *CURDALL {
	return &CURDALL{
		Car:  NewCarCURD(db),
		User: NewUserCURD(db),
	}
}

func (c *CURDALL) RegisterRouterALL(r interface{}) {
	c.Car.RegisterRouter(r)
	c.User.RegisterRouter(r)
}

func NewCarCURD(db *ent.Client) *CarCURD {
	return &CarCURD{
		Db: db,
	}
}

func NewUserCURD(db *ent.Client) *UserCURD {
	return &UserCURD{
		Db: db,

		CarObj: &CarCURD{
			db,
		},
	}
}
