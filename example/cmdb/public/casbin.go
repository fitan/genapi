package public

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var enforcer *casbin.Enforcer
var casbinLock sync.Mutex

type Header struct {
	// 这是Header 的姓名
	Name string `header:"name"`
	// 这是age 的姓名
	Age int `header:"age"`
	HeaderSub HeaderSub

}

type HeaderSub struct {
	// 这是age sub 的姓名
	AgeSub int `header:"ageSub"`
}

func newCasbin() (*casbin.Enforcer, error) {
	a, err := gormadapter.NewAdapter("mysql", GetConf().Mysql.Addr, true)
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
		casbinLock.Lock()
		defer casbinLock.Unlock()
		e, err := newCasbin()
		if err != nil {
			GetXLog().Error().Err(err).Msg("")
		}
		enforcer = e
	}
	return enforcer
}
