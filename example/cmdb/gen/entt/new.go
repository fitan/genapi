package entt

import (
	"cmdb/ent"
)

type CURDALL struct {
	Project     *ProjectCURD
	RoleBinding *RoleBindingCURD
	Server      *ServerCURD
	Service     *ServiceCURD
	User        *UserCURD
}

func NewCURDALL(db *ent.Client) *CURDALL {
	return &CURDALL{
		Project:     NewProjectCURD(db),
		RoleBinding: NewRoleBindingCURD(db),
		Server:      NewServerCURD(db),
		Service:     NewServiceCURD(db),
		User:        NewUserCURD(db),
	}
}

func (c *CURDALL) RegisterRouterALL(r interface{}) {
	c.Project.RegisterRouter(r)
	c.RoleBinding.RegisterRouter(r)
	c.Server.RegisterRouter(r)
	c.Service.RegisterRouter(r)
	c.User.RegisterRouter(r)
}

func NewProjectCURD(db *ent.Client) *ProjectCURD {
	return &ProjectCURD{
		Db: db,

		RoleBindingObj: &RoleBindingCURD{
			db,
		},

		ServiceObj: &ServiceCURD{
			db,
		},
	}
}

func NewRoleBindingCURD(db *ent.Client) *RoleBindingCURD {
	return &RoleBindingCURD{
		Db: db,

		ProjectObj: &ProjectCURD{
			db,
		},

		ServiceObj: &ServiceCURD{
			db,
		},

		UserObj: &UserCURD{
			db,
		},
	}
}

func NewServerCURD(db *ent.Client) *ServerCURD {
	return &ServerCURD{
		Db: db,

		ServiceObj: &ServiceCURD{
			db,
		},
	}
}

func NewServiceCURD(db *ent.Client) *ServiceCURD {
	return &ServiceCURD{
		Db: db,

		RoleBindingObj: &RoleBindingCURD{
			db,
		},

		ServerObj: &ServerCURD{
			db,
		},

		ProjectObj: &ProjectCURD{
			db,
		},
	}
}

func NewUserCURD(db *ent.Client) *UserCURD {
	return &UserCURD{
		Db: db,

		RoleBindingObj: &RoleBindingCURD{
			db,
		},
	}
}
