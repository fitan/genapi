package entt

import (
	"cmdb/ent"
)

type CURDALL struct {
	Alert       *AlertCURD
	Project     *ProjectCURD
	RoleBinding *RoleBindingCURD
	Server      *ServerCURD
	Service     *ServiceCURD
	User        *UserCURD
}

func NewCURDALL(db *ent.Client) *CURDALL {
	return &CURDALL{
		Alert:       NewAlertCURD(db),
		Project:     NewProjectCURD(db),
		RoleBinding: NewRoleBindingCURD(db),
		Server:      NewServerCURD(db),
		Service:     NewServiceCURD(db),
		User:        NewUserCURD(db),
	}
}

func (c *CURDALL) RegisterRouterALL(r interface{}) {
	c.Alert.RegisterRouter(r)
	c.Project.RegisterRouter(r)
	c.RoleBinding.RegisterRouter(r)
	c.Server.RegisterRouter(r)
	c.Service.RegisterRouter(r)
	c.User.RegisterRouter(r)
}

func NewAlertCURD(db *ent.Client) *AlertCURD {
	return &AlertCURD{
		Db: db,
	}
}

func NewProjectCURD(db *ent.Client) *ProjectCURD {
	return &ProjectCURD{
		Db: db,

		RoleBindingObj: &RoleBindingCURD{
			Db: db,
		},

		ServiceObj: &ServiceCURD{
			Db: db,
		},
	}
}

func NewRoleBindingCURD(db *ent.Client) *RoleBindingCURD {
	return &RoleBindingCURD{
		Db: db,

		ProjectObj: &ProjectCURD{
			Db: db,
		},

		ServiceObj: &ServiceCURD{
			Db: db,
		},

		UserObj: &UserCURD{
			Db: db,
		},
	}
}

func NewServerCURD(db *ent.Client) *ServerCURD {
	return &ServerCURD{
		Db: db,

		ServiceObj: &ServiceCURD{
			Db: db,
		},
	}
}

func NewServiceCURD(db *ent.Client) *ServiceCURD {
	return &ServiceCURD{
		Db: db,

		RoleBindingObj: &RoleBindingCURD{
			Db: db,
		},

		ServerObj: &ServerCURD{
			Db: db,
		},

		ProjectObj: &ProjectCURD{
			Db: db,
		},
	}
}

func NewUserCURD(db *ent.Client) *UserCURD {
	return &UserCURD{
		Db: db,

		RoleBindingObj: &RoleBindingCURD{
			Db: db,
		},

		AlertObj: &AlertCURD{
			Db: db,
		},
	}
}
