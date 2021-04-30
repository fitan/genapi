package models

import (
	"cmdb/public"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var enforcer *casbin.Enforcer
var lock sync.Mutex

func newCasbin() (*casbin.Enforcer, error) {
	a, err := gormadapter.NewAdapter("mysql", public.GetConf().Mysql.Addr, true)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer("conf/rbac_model.conf", a)
	if err != nil {
		return nil, err
	}
	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return e, nil
}

func GetCasbin() *casbin.Enforcer {
	if enforcer == nil {
		lock.Lock()
		defer lock.Unlock()
		e, err := newCasbin()
		if err != nil {
			public.GetXLog().Error().Err(err).Msg("")
		}
		enforcer = e
	}
	return enforcer
}
